package data_pgsql

import (
	"github.com/infrago/infra"
	"github.com/infrago/data"
)

func Driver() data.Driver {
	return &postgresqlDriver{}
}

func init() {
	drv := Driver()
	infra.Register("pgsql", drv)
	infra.Register("postgres", drv)
	infra.Register("postgresql", drv)
}
