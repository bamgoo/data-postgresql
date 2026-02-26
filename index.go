package data_pgsql

import (
	"github.com/bamgoo/bamgoo"
	"github.com/bamgoo/data"
)

func Driver() data.Driver {
	return &pgsqlDriver{}
}

func init() {
	bamgoo.Register("pgsql", Driver())
	bamgoo.Register("postgres", Driver())
}
