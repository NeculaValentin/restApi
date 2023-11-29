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

func (fc *FileControllerImpl) GetFile(c *gin.Context) {
	username, docID := checkParams(c)
	_, err := checkAuthorization(c, fc, username)
	if err != nil {
		return
	}
	content, err := fc.svc.GetFile(username, docID)
	if err != nil {
		common.NewAPIError(c, http.StatusNotFound, err, "file not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"content": content})
}

func (fc *FileControllerImpl) CreateFile(c *gin.Context) {
	username, docID := checkParams(c)
	_, err := checkAuthorization(c, fc, username)
	if err != nil {
		return
	}
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		common.NewAPIError(c, http.StatusBadRequest, err, "invalid request body")
		return
	}

	size := fc.svc.CreateFile(username, docID, requestBody)
	c.JSON(http.StatusOK, gin.H{"size": size})
}

func (fc *FileControllerImpl) UpdateFile(c *gin.Context) {
	username, docID := checkParams(c)
	_, err := checkAuthorization(c, fc, username)
	if err != nil {
		return
	}
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		common.NewAPIError(c, http.StatusBadRequest, err, "invalid request body")
		return
	}

	size := fc.svc.UpdateFile(username, docID, requestBody)
	c.JSON(http.StatusOK, gin.H{"size": size})
}

func (fc *FileControllerImpl) DeleteFile(c *gin.Context) {
	username, docID := checkParams(c)
	_, errAuth := checkAuthorization(c, fc, username)
	if errAuth != nil {
		return
	}
	err := fc.svc.DeleteFile(username, docID)
	if err != nil {
		common.NewAPIError(c, http.StatusInternalServerError, err, "error deleting file")
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (fc *FileControllerImpl) GetAllUserDocs(c *gin.Context) {
	username := c.Param("username")
	_, errAuth := checkAuthorization(c, fc, username)
	if errAuth != nil {
		return
	}
	if username == "" {
		common.NewAPIError(c, http.StatusBadRequest, nil, "username cannot be empty")
		return
	}
	docs := fc.svc.GetAllUserDocs(username)
	if docs == nil {
		docs = make(map[string]string)
	}
	c.JSON(http.StatusOK, docs)
}

// checkAuthorization checks if the request is authorized
func checkAuthorization(c *gin.Context, fc *FileControllerImpl, usernameParam string) (string, error) {
	token := c.GetHeader("Authorization")
	if token == "" {
		common.NewAPIError(c, http.StatusUnauthorized, fmt.Errorf("authorization header is required"), "authorization required")
		return "", fmt.Errorf("authorization header is required")
	}

	// Validate token
	username, err := fc.as.ValidateToken(token)
	if err != nil {
		common.NewAPIError(c, http.StatusUnauthorized, err, "invalid token")
		return "", err
	}
	if usernameParam != "" && usernameParam != username {
		common.NewAPIError(c, http.StatusUnauthorized, fmt.Errorf("username in token and path do not match"), "invalid username")
		return "", fmt.Errorf("username in token and path do not match")
	}
	return username, nil
}

// checkParams checks if the username and docID are valid
func checkParams(c *gin.Context) (string, string) {
	username := c.Param("username")
	docID := c.Param("doc_id")
	if username == "" || docID == "" {
		common.NewAPIError(c, http.StatusBadRequest, nil, "invalid input parameters")
	}
	return username, docID
}
