package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"umiEvient/acara"

	"github.com/gin-gonic/gin"
)

type acaraHandler struct {
	acaraService acara.Service
}

func NewacaraHandler(acaraService acara.Service) *acaraHandler {
	return &acaraHandler{acaraService}
}

func (h *acaraHandler) Index(c *gin.Context) {
	campaigns, err := h.acaraService.GetAllAcaras()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "acara_index.html", gin.H{"campaigns": campaigns})
}

func (h *acaraHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new_acara.html", nil)
}

func (h *acaraHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingAcara, err := h.acaraService.GetAcaraByID(acara.GetAcaraDetailInput{ID: id})
	fmt.Println(existingAcara, "acaraa")
	if err != nil {

		c.HTML(http.StatusOK, "error.html", nil)
		return
	}

	input := acara.FormUpdateAcaraInput{}

	input.ID = existingAcara.ID
	input.NamaAcara = existingAcara.Nama_acara
	input.Harga = existingAcara.Harga
	input.TanggalAcara = existingAcara.Tanggal_acara
	input.Deskripsi = existingAcara.Deskripsi
	input.ImgUrl = existingAcara.ImgURL

	c.HTML(http.StatusOK, "acara_edit.html", input)
}

func (h *acaraHandler) Create(c *gin.Context) {
	var input acara.FormCreateAcaraInput
	file, err := c.FormFile("file")

	err = c.ShouldBind(&input)
	if err != nil {

		data := acara.FormCreateAcaraTemplateData{
			FormCreateAcaraInput: input,
			Error:                err.Error(),
		}
		c.HTML(http.StatusOK, "new_acara.html", data)
		return
	}

	path := fmt.Sprintf("images/%d-%s", 1, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	createAcaraInput := acara.CreateAcaraInput{}
	createAcaraInput.NamaAcara = input.NamaAcara
	createAcaraInput.Harga = input.Harga
	createAcaraInput.TanggalAcara = input.TanggalAcara
	createAcaraInput.Deskripsi = input.Deskripsi
	createAcaraInput.ImgUrl = path

	fmt.Println(createAcaraInput.NamaAcara, "dapatnya?")
	fmt.Println(createAcaraInput, "sini")

	_, err = h.acaraService.CreateAcara(createAcaraInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/acara")
}

func (h *acaraHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input acara.FormUpdateAcaraInput
	file, err := c.FormFile("file")

	err = c.ShouldBind(&input)
	fmt.Println(err, "kenpa?")
	if err != nil {
		input.Error = err
		input.ID = id
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	path := fmt.Sprintf("images/%d-%s", 1, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	updateInput := acara.CreateAcaraInput{}
	updateInput.NamaAcara = input.NamaAcara
	updateInput.Harga = input.Harga
	updateInput.TanggalAcara = input.TanggalAcara
	updateInput.Deskripsi = input.Deskripsi
	updateInput.ImgUrl = path

	_, err = h.acaraService.UpdateAcara(acara.GetAcaraDetailInput{ID: id}, updateInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/acara")
}

func (h *acaraHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	_, err := h.acaraService.DeleteAcara(acara.GetAcaraDetailInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/acara")
}
