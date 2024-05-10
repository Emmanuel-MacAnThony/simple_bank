createdb:
	docker exec -it postgres-simple-bank createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres-simple-bank dropdb simple_bank
postgres:
	docker run --name postgres-simple-bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
migrateup:
	migrate -path db/migrations -database "postgresql://root:feb061999@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:feb061999@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server