package controller

import (
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

type FileController interface {
	GetFile(c *gin.Context)
	GetUserById(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUserDocs(c *gin.Context)
}

// RegisterRoutes registers the authentication routes
func (fc *FileControllerImpl) RegisterRoutes(router *gin.RouterGroup) {

	router.GET("/<string:username>/<string:doc_id>", fc.GetFile)
	router.POST("/<string:username>/<string:doc_id>", fc.CreateFile)
	router.PUT("/<string:username>/<string:doc_id>", fc.UpdateFile)
	router.DELETE("/<string:username>/<string:doc_id>", fc.DeleteFile)
	router.GET("/<string:username>/_all_docs", fc.GetAllUserDocs)

}

// GetFile godoc
// @Summary Get a file
// @Description get file by username and document ID
// @Tags files
// @Accept  json
// @Produce  json
// @Param   username path string true "Username"
// @Param   doc_id path string true "Document ID"
// @Success 200 {string} string "successful operation"
// @Failure 400 {object} object "Bad request"
// @Failure 404 {object} object "Not Found"
// @Router /{username}/{doc_id} [get]
func (fc *FileControllerImpl) GetFile(c *gin.Context) {
	username := c.Param("username")
	docID := c.Param("docID")

	content := fc.svc.GetFile(username, docID)
	c.JSON(http.StatusOK, gin.H{"content": content})
}

func (fc *FileControllerImpl) CreateFile(c *gin.Context) {
	username := c.Param("username")
	docID := c.Param("doc_id")

	var requestBody struct {
		DocContent string `json:"doc_content"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	size := fc.svc.CreateFile(username, docID, requestBody.DocContent)
	c.JSON(http.StatusOK, gin.H{"size": size})
}

func (fc *FileControllerImpl) UpdateFile(c *gin.Context) {
	username := c.Param("username")
	docID := c.Param("doc_id")

	var requestBody struct {
		DocContent string `json:"doc_content"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	size := fc.svc.UpdateFile(username, docID, requestBody.DocContent)
	c.JSON(http.StatusOK, gin.H{"size": size})
}

func (fc *FileControllerImpl) DeleteFile(c *gin.Context) {
	username := c.Param("username")
	docID := c.Param("doc_id")
	fc.svc.DeleteFile(username, docID)
	c.JSON(http.StatusOK, gin.H{})
}

func (fc *FileControllerImpl) GetAllUserDocs(c *gin.Context) {
	username := c.Param("username")
	docs := fc.svc.GetAllUserDocs(username)
	c.JSON(http.StatusOK, docs)
}
