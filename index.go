package data_postgres

import (
	"github.com/infrago/data"
	"github.com/infrago/infra"
)

func Driver() data.Driver {
	return &postgresDriver{}
}

func init() {
	drv := Driver()
	infra.Register("pgsql", drv)
	infra.Register("postgres", drv)
	infra.Register("postgresql", drv)
}
