migrateup:
	migrate -path db/migration/ -database "postgresql://postgres:postgres@localhost:5432/backend_mastery?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://postgres:postgres@localhost:5432/backend_mastery?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/tdhafin/backend-mastery/db/sqlc Store