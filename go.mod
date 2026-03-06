module github.com/infrago/data-postgres

go 1.25.3

require (
	github.com/infrago/data v0.11.1
	github.com/infrago/infra v0.11.1
	github.com/lib/pq v1.10.9
)

replace github.com/infrago/data => ../data

require (
	github.com/infrago/base v0.11.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
