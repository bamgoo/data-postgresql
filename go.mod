module github.com/bamgoo/data-postgresql

go 1.25.3

require (
	github.com/bamgoo/bamgoo v0.0.0
	github.com/bamgoo/data v0.0.0
	github.com/lib/pq v1.10.9
)

replace github.com/bamgoo/bamgoo => ../bamgoo
replace github.com/bamgoo/data => ../data
