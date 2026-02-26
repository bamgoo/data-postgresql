package data_pgsql

import (
	"database/sql"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/bamgoo/data"
	_ "github.com/lib/pq"
)

type (
	postgresqlDriver struct{}

	postgresqlConnection struct {
		instance *data.Instance
		db       *sql.DB
		actives  int64
	}

	postgresqlDialect struct{}
)

func (d *postgresqlDriver) Connect(inst *data.Instance) (data.Connection, error) {
	return &postgresqlConnection{instance: inst}, nil
}

func (c *postgresqlConnection) Open() error {
	dsn := strings.TrimSpace(c.instance.Config.Url)
	if dsn == "" {
		if v, ok := c.instance.Setting["dsn"].(string); ok {
			dsn = v
		}
	}
	if dsn == "" {
		return fmt.Errorf("missing pgsql dsn")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return err
	}
	c.db = db
	return nil
}

func (c *postgresqlConnection) Close() error {
	if c.db == nil {
		return nil
	}
	err := c.db.Close()
	c.db = nil
	return err
}

func (c *postgresqlConnection) Health() data.Health {
	return data.Health{Workload: atomic.LoadInt64(&c.actives)}
}

func (c *postgresqlConnection) DB() *sql.DB {
	return c.db
}

func (c *postgresqlConnection) Dialect() data.Dialect {
	return postgresqlDialect{}
}

func (postgresqlDialect) Name() string { return "pgsql" }
func (postgresqlDialect) Quote(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, `"`, ``)
	return `"` + s + `"`
}
func (postgresqlDialect) Placeholder(n int) string { return fmt.Sprintf("$%d", n) }
func (postgresqlDialect) SupportsILike() bool      { return true }
func (postgresqlDialect) SupportsReturning() bool  { return true }
