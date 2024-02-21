package svc

import (
	"encoding/json"
	"io/fs"
)

type FileHandlerIface interface {
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, data []byte, perm fs.FileMode) error
}

type FileServiceIface interface {
	ReadFile(string) (string, error)
	SaveFile(fileName string, content string) error
	SaveJson(fileName string, content any) error
}

type FileService struct {
	reader func(name string) ([]byte, error)
	writer func(name string, data []byte, perm fs.FileMode) error
}

func (fls FileService) ReadFile(fileName string) (string, error) {
	bts, err := fls.reader(fileName)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

func (fls FileService) SaveFile(fileName, content string) error {
	err := fls.writer(fileName, []byte(content), 0777)
	if err != nil {
		return err
	}
	return nil
}

func (fls FileService) SaveJson(fileName string, content any) error {

	marshaled, err := json.Marshal(content)
	if err != nil {
		return err
	}

	err = fls.writer(fileName, marshaled, 0777)
	if err != nil {
		return err
	}
	return nil
}

func NewFileService(
	reader func(name string) ([]byte, error),
	writer func(name string, data []byte, perm fs.FileMode) error,
) FileServiceIface {
	return FileService{reader: reader, writer: writer}
}
