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
	pgsqlDriver struct{}

	pgsqlConnection struct {
		instance *data.Instance
		db       *sql.DB
		actives  int64
	}

	pgsqlDialect struct{}
)

func (d *pgsqlDriver) Connect(inst *data.Instance) (data.Connection, error) {
	return &pgsqlConnection{instance: inst}, nil
}

func (c *pgsqlConnection) Open() error {
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

func (c *pgsqlConnection) Close() error {
	if c.db == nil {
		return nil
	}
	err := c.db.Close()
	c.db = nil
	return err
}

func (c *pgsqlConnection) Health() data.Health {
	return data.Health{Workload: atomic.LoadInt64(&c.actives)}
}

func (c *pgsqlConnection) DB() *sql.DB {
	return c.db
}

func (c *pgsqlConnection) Dialect() data.Dialect {
	return pgsqlDialect{}
}

func (pgsqlDialect) Name() string { return "pgsql" }
func (pgsqlDialect) Quote(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, `"`, ``)
	return `"` + s + `"`
}
func (pgsqlDialect) Placeholder(n int) string { return fmt.Sprintf("$%d", n) }
func (pgsqlDialect) SupportsILike() bool      { return true }
func (pgsqlDialect) SupportsReturning() bool  { return true }
