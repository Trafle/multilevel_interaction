package db

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type Connection struct {
	DbName         string
	User, Password string
	Host           string
}

func (c *Connection) formatConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Password, c.Host, c.DbName)
}

func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("mysql", c.formatConnection())
}