package repository

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"restApi/internal/app/common"
)

type FileRepositoryImpl struct {
	// Base directory where user files are stored
	baseDir string
}

func NewFileRepository(baseDir string) *FileRepositoryImpl {
	return &FileRepositoryImpl{baseDir}
}

type FileRepository interface {
	GetFile(username, docID string) (string, error)
	CreateFile(username, docID string, content []byte) (int, error)
	UpdateFile(username, docID string, content []byte) (int, error)
	DeleteFile(username, docID string)
	GetAllUserDocs(username string) (map[string]string, error)
}

// UTILS

func (fr *FileRepositoryImpl) filePath(username, docID string) string {
	return filepath.Join(fr.baseDir, username, docID+".json")
}

func (fr *FileRepositoryImpl) directoryPath(username string) string {
	return filepath.Join(fr.baseDir, username)
}

func (fr *FileRepositoryImpl) ensureUserFolderExists(username string) error {
	userFolderPath := filepath.Join(fr.baseDir, username)
	return os.MkdirAll(userFolderPath, 0755)
}

//READ ONLY

func (fr *FileRepositoryImpl) GetFile(username, docID string) (string, error) {
	filePath := fr.filePath(username, docID)
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", common.NewAPIError(http.StatusNotFound, err, "file not found")
		} else {
			// Handle other types of errors
			return "", common.NewAPIError(http.StatusNotFound, err, "error reading the file")
		}
	}
	return string(content), nil
}

func (fr *FileRepositoryImpl) GetAllUserDocs(username string) (map[string]string, error) {
	userFolderPath := filepath.Join(fr.baseDir, username)
	files, err := os.ReadDir(userFolderPath)
	if err != nil {
		return nil, common.NewAPIError(http.StatusInternalServerError, err, "error reading directory")
	}

	userDocs := make(map[string]string)
	for _, file := range files {
		if !file.IsDir() {
			docID := file.Name()
			content, err := os.ReadFile(filepath.Join(userFolderPath, docID))
			if err != nil {
				return nil, common.NewAPIError(http.StatusInternalServerError, err, fmt.Sprintf("error reading file %s", docID))
			}
			userDocs[docID] = string(content)
		}
	}

	return userDocs, nil
}

//READ/WRITE

func (fr *FileRepositoryImpl) CreateFile(username, docID string, content []byte) (int, error) {
	if err := fr.ensureUserFolderExists(username); err != nil {
		return 0, common.NewAPIError(http.StatusInternalServerError, err, "error creating user folder")
	}

	filePath := fr.filePath(username, docID)
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return 0, common.NewAPIError(http.StatusInternalServerError, err, "error writing file")
	}
	return len(content), nil
}

func (fr *FileRepositoryImpl) UpdateFile(username, docID string, content []byte) (int, error) {
	size, err := fr.CreateFile(username, docID, content)
	if err != nil {
		return 0, common.NewAPIError(http.StatusInternalServerError, err, "error updating file")
	}
	return size, nil
}

func (fr *FileRepositoryImpl) DeleteFile(username, docID string) {
	filePath := fr.filePath(username, docID)
	err := os.Remove(filePath)
	if err != nil {
		_ = common.NewAPIError(http.StatusInternalServerError, err, "error deleting file")
		return
	}

	dirPath := fr.directoryPath(username) // Assuming you have a method to get the directory path
	files, err := os.ReadDir(dirPath)
	if err != nil {
		_ = common.NewAPIError(http.StatusInternalServerError, err, "error reading directory")
		return
	}

	if len(files) == 0 {
		err = os.Remove(dirPath)
		if err != nil {
			_ = common.NewAPIError(http.StatusInternalServerError, err, "error deleting directory")
		}
	}
}
