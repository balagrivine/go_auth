module github.com/balagrivine/go_auth/config

go 1.20

replace github.com/balagrivine/go_auth/internal/database => ../internal/database

require github.com/lib/pq v1.10.9 // indirect
