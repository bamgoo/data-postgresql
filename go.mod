module github.com/infrago/data-postgresql

go 1.25.3

require (
	github.com/infrago/infra v0.0.0
	github.com/infrago/data v0.0.0
	github.com/lib/pq v1.10.9
)

require (
	github.com/infrago/base v0.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
)

replace github.com/infrago/infra => ../infra

replace github.com/infrago/data => ../data
