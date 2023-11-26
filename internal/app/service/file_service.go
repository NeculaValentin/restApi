package service

import (
	"net/http"
	"restApi/internal/app/common"
	"restApi/internal/app/repository"
)

type FileServiceImpl struct {
	repo repository.FileRepository
}

func NewFileService(repo repository.FileRepository) *FileServiceImpl {
	return &FileServiceImpl{repo}
}

type FileService interface {
	GetFile(username, docID string) string
	CreateFile(username, docID, content string) int
	UpdateFile(username, docID, content string) int
	DeleteFile(username, docID string)
	GetAllUserDocs(username string) map[string]string
}

func (fs *FileServiceImpl) GetFile(username, docID string) string {
	if username == "" || docID == "" {
		_ = common.NewAPIError(http.StatusBadRequest, nil, "username or document ID cannot be empty")
	}
	content, _ := fs.repo.GetFile(username, docID)
	return content
}

func (fs *FileServiceImpl) CreateFile(username, docID, content string) int {
	if username == "" || docID == "" || content == "" {
		_ = common.NewAPIError(http.StatusBadRequest, nil, "invalid input parameters")
	}
	size, err := fs.repo.CreateFile(username, docID, content)
	if err != nil {
		// Handle specific errors (e.g., write errors, permission issues)
		return 0
	}
	return size
}

func (fs *FileServiceImpl) UpdateFile(username, docID, content string) int {
	size, _ := fs.repo.UpdateFile(username, docID, content)
	return size
}

func (fs *FileServiceImpl) DeleteFile(username, docID string) {
	fs.repo.DeleteFile(username, docID)

}

func (fs *FileServiceImpl) GetAllUserDocs(username string) map[string]string {
	if username == "" {
		_ = common.NewAPIError(http.StatusBadRequest, nil, "username cannot be empty")
	}
	docs, _ := fs.repo.GetAllUserDocs(username)
	return docs
}
