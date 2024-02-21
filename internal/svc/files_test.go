package svc_test

import (
	"encoding/json"
	"errors"
	"go-make-tests/internal/svc"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileHandler struct {
	mock.Mock
}

func (m *MockFileHandler) ReadFile(name string) ([]byte, error) {
	args := m.Called(name)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockFileHandler) WriteFile(name string, data []byte, perm fs.FileMode) error {
	args := m.Called(name, data, perm)
	return args.Error(0)
}

func TestFileService_ReadFile(t *testing.T) {
	fileHandler := new(MockFileHandler)
	fileService := svc.NewFileService(fileHandler.ReadFile, fileHandler.WriteFile)

	fileData := []byte("file content")
	fileName := "testfile.txt"

	fileHandler.On("ReadFile", fileName).Return(fileData, nil)

	content, err := fileService.ReadFile(fileName)

	assert.NoError(t, err)
	assert.Equal(t, string(fileData), content)

	fileHandler.AssertExpectations(t)
}

func TestFileService_SaveFile(t *testing.T) {
	fileHandler := new(MockFileHandler)
	fileService := svc.NewFileService(fileHandler.ReadFile, fileHandler.WriteFile)

	fileContent := "file content"
	fileName := "testfile.txt"

	fileHandler.On("WriteFile", fileName, []byte(fileContent), fs.FileMode(0777)).Return(nil)

	err := fileService.SaveFile(fileName, fileContent)

	assert.NoError(t, err)

	fileHandler.AssertExpectations(t)
}

func TestFileService_SaveJson(t *testing.T) {
	fileHandler := new(MockFileHandler)
	fileService := svc.NewFileService(fileHandler.ReadFile, fileHandler.WriteFile)

	jsonContent := map[string]interface{}{
		"key": "value",
	}
	marshaledContent, _ := json.Marshal(jsonContent)

	fileName := "test.json"

	fileHandler.On("WriteFile", fileName, marshaledContent, fs.FileMode(0777)).Return(nil)

	err := fileService.SaveJson(fileName, jsonContent)

	assert.NoError(t, err)

	fileHandler.AssertExpectations(t)
}

func TestFileService_SaveJson_MarshalError(t *testing.T) {
	fileHandler := new(MockFileHandler)
	fileService := svc.NewFileService(fileHandler.ReadFile, fileHandler.WriteFile)

	jsonContent := make(chan int) // Unsupported type which will cause Marshal error
	fileName := "test.json"

	err := fileService.SaveJson(fileName, jsonContent)

	assert.Error(t, err)
	assert.EqualError(t, err, "json: unsupported type: chan int")

	fileHandler.AssertExpectations(t)
}

func TestFileService_ReadFile_Error(t *testing.T) {
	fileHandler := new(MockFileHandler)
	fileService := svc.NewFileService(fileHandler.ReadFile, fileHandler.WriteFile)

	fileName := "invalidfile.txt"
	expectedError := errors.New("read file error")

	fileHandler.On("ReadFile", fileName).Return([]byte(nil), expectedError)

	_, err := fileService.ReadFile(fileName)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())

	fileHandler.AssertExpectations(t)
}

func TestFileService_SaveFile_Error(t *testing.T) {
	fileHandler := new(MockFileHandler)
	fileService := svc.NewFileService(fileHandler.ReadFile, fileHandler.WriteFile)

	fileContent := "file content"
	fileName := "testfile.txt"
	expectedError := errors.New("write file error")

	fileHandler.On("WriteFile", fileName, []byte(fileContent), fs.FileMode(0777)).Return(expectedError)

	err := fileService.SaveFile(fileName, fileContent)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())

	fileHandler.AssertExpectations(t)
}
