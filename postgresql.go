package data_postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/infrago/data"
	_ "github.com/lib/pq"
)

type (
	postgresDriver struct{}

	postgresConnection struct {
		instance *data.Instance
		db       *sql.DB
		actives  int64
	}

	postgresDialect struct{}
)

func (d *postgresDriver) Connect(inst *data.Instance) (data.Connection, error) {
	return &postgresConnection{instance: inst}, nil
}

func (c *postgresConnection) Open() error {
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

func (c *postgresConnection) Close() error {
	if c.db == nil {
		return nil
	}
	err := c.db.Close()
	c.db = nil
	return err
}

func (c *postgresConnection) Health() data.Health {
	return data.Health{Workload: atomic.LoadInt64(&c.actives)}
}

func (c *postgresConnection) DB() *sql.DB {
	return c.db
}

func (c *postgresConnection) Dialect() data.Dialect {
	return postgresDialect{}
}

func (postgresDialect) Name() string { return "pgsql" }
func (postgresDialect) Quote(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, `"`, ``)
	return `"` + s + `"`
}
func (postgresDialect) Placeholder(n int) string { return fmt.Sprintf("$%d", n) }
func (postgresDialect) SupportsILike() bool      { return true }
func (postgresDialect) SupportsReturning() bool  { return true }
