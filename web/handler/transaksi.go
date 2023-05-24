package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"umiEvient/transaction"

	"github.com/gin-gonic/gin"
)

type transaksi_handler struct {
	transaksiService transaction.Service
}

func Newtransaksi_handler(transaksiService transaction.Service) *transaksi_handler {
	return &transaksi_handler{transaksiService}
}

func (h *transaksi_handler) Index(c *gin.Context) {
	transaksi, err := h.transaksiService.GetAllTransactions()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	fmt.Println(transaksi, "coba")
	c.HTML(http.StatusOK, "transaksi_index.html", gin.H{"transaksi": transaksi})
}

func (h *transaksi_handler) Status(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	_, err := h.transaksiService.UpdateStatus(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/transaksi")
}
