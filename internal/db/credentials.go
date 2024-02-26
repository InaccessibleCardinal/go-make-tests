package db

import (
	"database/sql"
	"fmt"
	"go-make-tests/internal/types"
	"os"
)

type CredsDB struct {
	db *sql.DB
	dbName string
	tableName string
}

func (d *CredsDB) GetById(id int) (*types.Creds, error) {
	var creds types.Creds
	query := fmt.Sprintf("select * from %s.%s where ID=%d", d.dbName, d.tableName, id)
	err := d.db.QueryRow(query).Scan(&creds.ID, &creds.User, &creds.AccessKey, &creds.SecretKey)
	if err != nil {
		return nil, err
	}
	return &creds, nil
}

func (d *CredsDB) Upsert(creds types.Creds) {}

func (d *CredsDB) GetAll() ([]types.Creds, error) {
	query := fmt.Sprintf("select * from %s.%s", d.dbName, d.tableName)
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	var allCreds []types.Creds
	for rows.Next() {
		var creds types.Creds
		if err := rows.Scan(&creds.ID, &creds.User, &creds.AccessKey, &creds.SecretKey); err != nil {
			return nil, err
		}
		allCreds = append(allCreds, creds)
	}
	return allCreds, nil
}

func NewCredsDB() types.DBIface[types.Creds] {
	var (
		dbName = os.Getenv("DATABASE")
		tableName = os.Getenv("CREDS_TABLE")
	)
	return &CredsDB{db: getConnection(), dbName: dbName, tableName: tableName}
}