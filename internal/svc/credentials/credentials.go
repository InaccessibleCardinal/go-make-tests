package credentials

import (
	"go-make-tests/internal/types"
)

type CredentialsService struct {
	db types.DBIface[types.Creds]
}

func (cs *CredentialsService) GetCredentialsById(id int) (*types.Creds, error) {
	creds, err := cs.db.GetById(id)
	if err != nil {
		return nil, err
	}
	return creds, nil
}

func (cs *CredentialsService) GetAllCredentials() ([]types.Creds, error) {
	allCreds, err := cs.db.GetAll()
	if err != nil {
		return nil, err
	}
	return allCreds, err
}

func (cs *CredentialsService) AddCredentials(creds types.Creds) {
	
}

func New(db types.DBIface[types.Creds]) *CredentialsService {
	return &CredentialsService{db: db}
}