package controller

import (
	"io"
	"net/http"
	"restApi/internal/app/repository"
	"restApi/internal/app/service"

	"github.com/gin-gonic/gin"
)

type FileControllerImpl struct {
	svc service.FileService
}

func NewFileController() *FileControllerImpl {
	return &FileControllerImpl{svc: service.NewFileService(repository.NewFileRepository(""))}
}

// RegisterRoutes registers the authentication routes
func (fc *FileControllerImpl) RegisterRoutes(router *gin.RouterGroup) {

	router.GET("/:username/:doc_id", fc.GetFile)
	router.POST("/:username/:doc_id", fc.CreateFile)
	router.PUT("/:username/:doc_id", fc.UpdateFile)
	router.DELETE("/:username/:doc_id", fc.DeleteFile)
	router.GET("/:username/_all_docs", fc.GetAllUserDocs)

}

// GetFile godoc
// @Summary Get file content
// @Description Retrieves the content of the specified document
// @Tags files
// @Accept  json
// @Produce  json
// @Param   username path string true "Username"
// @Param   doc_id path string true "Document ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /{username}/{doc_id} [get]
func (fc *FileControllerImpl) GetFile(c *gin.Context) {
	username := c.Param("username")
	docID := c.Param("docID")

	content := fc.svc.GetFile(username, docID)
	c.JSON(http.StatusOK, gin.H{"content": content})
}

// CreateFile godoc
// @Summary Create a new file
// @Description Creates a new document with the specified content
// @Tags files
// @Accept  json
// @Produce  json
// @Param   username path string true "Username"
// @Param   doc_id path string true "Document ID"
// @Param   doc_content body string true "Document Content"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Router /{username}/{doc_id} [post]
func (fc *FileControllerImpl) CreateFile(c *gin.Context) {
	username := c.Param("username")
	docID := c.Param("doc_id")

	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	size := fc.svc.CreateFile(username, docID, requestBody)
	c.JSON(http.StatusOK, gin.H{"size": size})
}

// UpdateFile godoc
// @Summary Update file content
// @Description Updates the content of the specified document
// @Tags files
// @Accept  json
// @Produce  json
// @Param   username path string true "Username"
// @Param   doc_id path string true "Document ID"
// @Param   doc_content body string true "New Document Content"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Router /{username}/{doc_id} [put]
func (fc *FileControllerImpl) UpdateFile(c *gin.Context) {
	username := c.Param("username")
	docID := c.Param("doc_id")

	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	size := fc.svc.UpdateFile(username, docID, requestBody)
	c.JSON(http.StatusOK, gin.H{"size": size})
}

// DeleteFile godoc
// @Summary Delete a file
// @Description Deletes the specified document
// @Tags files
// @Accept	json
// @Produce	json
// @Param	username path string true "Username"
// @Param	doc_id path string true "Document ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /{username}/{doc_id} [delete]
func (fc *FileControllerImpl) DeleteFile(c *gin.Context) {
	username := c.Param("username")
	docID := c.Param("doc_id")
	fc.svc.DeleteFile(username, docID)
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllUserDocs godoc
// @Summary Get all user documents
// @Description Retrieves all documents for the specified user
// @Tags files
// @Accept	json
// @Produce	json
// @Param	username path string true "Username"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /{username}/_all_docs [get]
func (fc *FileControllerImpl) GetAllUserDocs(c *gin.Context) {
	username := c.Param("username")
	docs := fc.svc.GetAllUserDocs(username)
	c.JSON(http.StatusOK, docs)
}
