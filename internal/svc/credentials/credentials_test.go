package credentials_test

import (
	"errors"
	"fmt"
	c "go-make-tests/internal/svc/credentials"
	"go-make-tests/internal/types"
	"testing"
)

var (
	testAccessKey = "testAccessKey"
	testSecretKey = "testSecretKey"
	testUser = "testUser"
)

type MockCredsDB struct {
	creds *types.Creds
	allCreds []types.Creds
	err error
}

func (m *MockCredsDB) GetAll() ([]types.Creds, error) {
	return m.allCreds, m.err
}

func (m *MockCredsDB) GetById(id int) (*types.Creds, error) {
	return m.creds, m.err
}

func (m *MockCredsDB) Upsert(c types.Creds) {}

func Test_GetCredentialsById_Success(t *testing.T) {
	testService := c.New(&MockCredsDB{creds: &types.Creds{AccessKey: &testAccessKey, SecretKey: &testSecretKey}})

	actual, err := testService.GetCredentialsById(1)
	if err != nil {
		t.Fatalf("expected nil error but got %s\n", err.Error())
	}
	if *actual.AccessKey != testAccessKey {
		t.Fatalf("expected %s but got %s\n", testAccessKey, *actual.AccessKey)
	}
	if *actual.SecretKey != testSecretKey {
		t.Fatalf("expected %s but got %s\n", testSecretKey, *actual.SecretKey)
	}
}


func Test_GetCredentialsById_Error(t *testing.T) {
	testService := c.New(&MockCredsDB{creds: nil, err: errors.New("epic fail")})

	_, err := testService.GetCredentialsById(1); if err == nil {
		t.Fatal("error expected to be not nil")
	}
	
	if err.Error() != "epic fail" {
		t.Fatalf("expected nil error but got %s\n", err.Error())
	}
}

func TestGetAllCredentials_Success(t *testing.T) {
	access1 := "a1"
	secret1 := "s1"
	access2 := "a2"
	secret2 := "s2"
	access3 := "a3"
	secret3 := "s3"
	testCreds := []types.Creds{
		{AccessKey: &access1, SecretKey: &secret1},
		{AccessKey: &access2, SecretKey: &secret2},
		{AccessKey: &access3, SecretKey: &secret3},
	}
	mockDB := &MockCredsDB{allCreds: testCreds}
	testService := c.New(mockDB)

	allCreds, err := testService.GetAllCredentials()
	if err != nil {
		t.Fatalf("expected nil error but got %s\n", err.Error())
	}

	for i, creds := range allCreds {
		expectedAccessKey := fmt.Sprintf("a%d", i+1)
		expectedSecretKey := fmt.Sprintf("s%d", i+1)
		if (*creds.AccessKey != expectedAccessKey) {
			t.Fatalf("expected accessKey %s but got %s\n", expectedAccessKey, *creds.AccessKey)
		}
		if (*creds.SecretKey != expectedSecretKey) {
			t.Fatalf("expected accessKey %s but got %s\n", expectedSecretKey, *creds.SecretKey)
		}
	}
}