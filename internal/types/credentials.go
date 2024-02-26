package types

type Creds struct {
	ID        *int
	User      *string `mysql:"user"`
	AccessKey *string `mysql:"access_key"`
	SecretKey *string `mysql:"secret_key"`
}