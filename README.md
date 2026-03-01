# data-postgresql

`data-postgresql` 是 `data` 模块的 `postgresql` 驱动。

## 安装

```bash
go get github.com/infrago/data@latest
go get github.com/infrago/data-postgresql@latest
```

## 接入

```go
import (
    _ "github.com/infrago/data"
    _ "github.com/infrago/data-postgresql"
    "github.com/infrago/infra"
)

func main() {
    infra.Run()
}
```

## 配置示例

```toml
[data]
driver = "postgresql"
```

## 公开 API（摘自源码）

- `func Driver() data.Driver`
- `func (d *postgresqlDriver) Connect(inst *data.Instance) (data.Connection, error)`
- `func (c *postgresqlConnection) Open() error`
- `func (c *postgresqlConnection) Close() error`
- `func (c *postgresqlConnection) Health() data.Health`
- `func (c *postgresqlConnection) DB() *sql.DB`
- `func (c *postgresqlConnection) Dialect() data.Dialect`
- `func (postgresqlDialect) Name() string { return "pgsql" }`
- `func (postgresqlDialect) Quote(s string) string`
- `func (postgresqlDialect) Placeholder(n int) string { return fmt.Sprintf("$%d", n) }`
- `func (postgresqlDialect) SupportsILike() bool      { return true }`
- `func (postgresqlDialect) SupportsReturning() bool  { return true }`

## 排错

- driver 未生效：确认模块段 `driver` 值与驱动名一致
- 连接失败：检查 endpoint/host/port/鉴权配置
