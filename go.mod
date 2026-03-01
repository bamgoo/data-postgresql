module github.com/infrago/data-postgresql

go 1.25.3

require (
	github.com/infrago/data v0.8.2
	github.com/infrago/infra v0.8.2
	github.com/lib/pq v1.10.9
)

require (
	github.com/infrago/base v0.8.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
)

replace github.com/infrago/infra => ../bamgoo

replace github.com/infrago/data => ../data
