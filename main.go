package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"umiEvient/acara"
	"umiEvient/auth"
	"umiEvient/handler"
	"umiEvient/transaction"
	"umiEvient/user"
	webHandler "umiEvient/web/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/umievent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// repository
	userRepository := user.NewRepository(db)
	acaraRepository := acara.NewRepository(db)
	transaksiRepository := transaction.NewRepository(db)

	// service
	authService := auth.NewService()
	userService := user.NewService(userRepository)
	acaraService := acara.NewService(acaraRepository)
	transaksiService := transaction.NewService(transaksiRepository)

	// handler api
	userHandler := handler.NewUserHandler(userService, authService)
	acaraHandler := handler.NewacaraHandler(acaraService)
	transaksiHandler := handler.NewTransactionHandler(transaksiService)

	// handler web
	userWebHandler := webHandler.NewUserHandler(userService)
	acaraWebHandler := webHandler.NewacaraHandler(acaraService)
	transaksiwebHandler := webHandler.Newtransaksi_handler(transaksiService)
	sessionWebHandler := webHandler.NewSessionHandler(userService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// cookie
	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("umiEvient", cookieStore))

	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")
	api := router.Group("api/v1")

	// api routing
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/acara", acaraHandler.GetAcaras)
	api.GET("/acara/:id", acaraHandler.GetAcara)
	api.POST("/transaksi", authMiddleware(authService, userService), transaksiHandler.CreateTransaction)

	// routing
	router.GET("/users", authAdminMiddleware(), userWebHandler.Index)
	router.GET("/acara", authAdminMiddleware(), acaraWebHandler.Index)
	router.POST("/acara/create", authAdminMiddleware(), acaraWebHandler.Create)
	router.GET("/acara/new", authAdminMiddleware(), acaraWebHandler.New)
	router.GET("/acara/edit/:id", authAdminMiddleware(), acaraWebHandler.Edit)
	router.POST("/acara/update/:id", authAdminMiddleware(), acaraWebHandler.Update)
	router.GET("/acara/delete/:id", authAdminMiddleware(), acaraWebHandler.Delete)
	router.GET("/transaksi", authAdminMiddleware(), transaksiwebHandler.Index)
	router.GET("/transaksi/status/:id", authAdminMiddleware(), transaksiwebHandler.Status)

	router.GET("/login", sessionWebHandler.New)
	router.POST("/session", sessionWebHandler.Create)
	router.GET("/logout", sessionWebHandler.Destroy)

	router.Run(":9090")

}

// middlewar
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := gin.H{"error": "Authorization header is missing or invalid"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := authService.ValidateToken(tokenString)
		// fmt.Println(token, "apap")
		if err != nil {
			log.Printf("Error validating token: %s", err.Error())
			response := gin.H{"error": "Invalid token Apa"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := gin.H{"error": "Invalid token"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID, ok := claim["id"].(float64)
		if !ok {
			response := gin.H{"error": "Invalid user ID in token"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user, err := userService.GetUserByID(int(userID))
		if err != nil {
			response := gin.H{"error": "User not found"}
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}

		c.Set("currentUser", user)
	}
}

func authAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("userID")

		if userIDSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
