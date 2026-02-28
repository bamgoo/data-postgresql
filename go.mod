module github.com/bamgoo/data-postgresql

go 1.25.3

require (
	github.com/bamgoo/bamgoo v0.0.0
	github.com/bamgoo/data v0.0.0
	github.com/lib/pq v1.10.9
)

require (
	github.com/bamgoo/base v0.0.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
)

replace github.com/bamgoo/bamgoo => ../bamgoo

replace github.com/bamgoo/data => ../data
