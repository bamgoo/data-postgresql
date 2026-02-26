package data_pgsql

import (
	"github.com/bamgoo/bamgoo"
	"github.com/bamgoo/data"
)

func Driver() data.Driver {
	return &postgresqlDriver{}
}

func init() {
	drv := Driver()
	bamgoo.Register("pgsql", drv)
	bamgoo.Register("postgres", drv)
	bamgoo.Register("postgresql", drv)
}
