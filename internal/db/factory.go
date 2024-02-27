package db

import (
	"go-make-tests/internal/types"
)

func CredentialsDbFactory(systemFlag string) types.DBIface[types.Creds] {
	if systemFlag == "darwin" {
		return NewTempMacDB()
	}
	return NewCredsDB()
}

type TempMacDB struct{}

func (t *TempMacDB) GetAll() ([]types.Creds, error) {
	return []types.Creds{}, nil
}

func (t *TempMacDB) GetById(id int) (*types.Creds, error) {
	return nil, nil
}

func (t *TempMacDB) Upsert(c types.Creds) {}

func NewTempMacDB() types.DBIface[types.Creds] {
	return &TempMacDB{}
}
