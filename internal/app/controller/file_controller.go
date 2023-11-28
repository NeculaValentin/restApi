package controller

import (
	"fmt"
	"io"
	"net/http"
	"restApi/internal/app/common"
	"restApi/internal/app/repository"
	"restApi/internal/app/service"

	"github.com/gin-gonic/gin"
)

type FileControllerImpl struct {
	svc service.FileService
	as  service.AuthService
}

func NewFileController() *FileControllerImpl {
	return &FileControllerImpl{
		as:  service.NewAuthService(repository.NewUserRepository(common.ConnectToDB())),
		svc: service.NewFileService(repository.NewFileRepository(""))}
}

type FileController interface {
	GetFile(c *gin.Context)
	CreateFile(c *gin.Context)
	UpdateFile(c *gin.Context)
	DeleteFile(c *gin.Context)
	GetAllUserDocs(c *gin.Context)
}

// RegisterRoutes registers the authentication routes
func (fc *FileControllerImpl) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/:username/:doc_id", fc.GetFile)
	router.POST("/:username/:doc_id", fc.CreateFile)
	router.PUT("/:username/:doc_id", fc.UpdateFile)
	router.DELETE("/:username/:doc_id", fc.DeleteFile)
	router.GET("/:username/_all_docs", fc.GetAllUserDocs)
}

func checkParams(c *gin.Context) (string, string) {
	username := c.Param("username")
	docID := c.Param("doc_id")
	if username == "" || docID == "" {
		common.NewAPIError(c, http.StatusBadRequest, nil, "invalid input parameters")
	}
	return username, docID
}

func (fc *FileControllerImpl) GetFile(c *gin.Context) {
	if checkAuthorization(c, fc) {
		return
	}
	username, docID := checkParams(c)
	content := fc.svc.GetFile(username, docID)
	c.JSON(http.StatusOK, gin.H{"content": content})
}

func checkAuthorization(c *gin.Context, fc *FileControllerImpl) bool {
	token := c.GetHeader("Authorization")
	if token == "" {
		common.NewAPIError(c, http.StatusUnauthorized, fmt.Errorf("authorization header is required"), "authorization required")
		return true
	}

	// Validate token
	_, err := fc.as.ValidateToken(token)
	if err != nil {
		common.NewAPIError(c, http.StatusUnauthorized, err, "invalid token")
		return true
	}
	return false
}

func (fc *FileControllerImpl) CreateFile(c *gin.Context) {
	if checkAuthorization(c, fc) {
		return
	}
	username, docID := checkParams(c)
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		common.NewAPIError(c, http.StatusBadRequest, err, "invalid request body")
		return
	}

	size := fc.svc.CreateFile(username, docID, requestBody)
	c.JSON(http.StatusOK, gin.H{"size": size})
}

func (fc *FileControllerImpl) UpdateFile(c *gin.Context) {
	if checkAuthorization(c, fc) {
		return
	}
	username, docID := checkParams(c)
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		common.NewAPIError(c, http.StatusBadRequest, err, "invalid request body")
		return
	}

	size := fc.svc.UpdateFile(username, docID, requestBody)
	c.JSON(http.StatusOK, gin.H{"size": size})
}

func (fc *FileControllerImpl) DeleteFile(c *gin.Context) {
	if checkAuthorization(c, fc) {
		return
	}
	username, docID := checkParams(c)
	err := fc.svc.DeleteFile(username, docID)
	if err != nil {
		common.NewAPIError(c, http.StatusInternalServerError, err, "error deleting file")
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (fc *FileControllerImpl) GetAllUserDocs(c *gin.Context) {
	if checkAuthorization(c, fc) {
		return
	}
	username := c.Param("username")
	if username == "" {
		common.NewAPIError(c, http.StatusBadRequest, nil, "username cannot be empty")
		return
	}
	docs := fc.svc.GetAllUserDocs(username)
	c.JSON(http.StatusOK, docs)
}
