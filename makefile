createdb:
	docker exec -it postgres-simple-bank createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres-simple-bank dropdb simple_bank
postgres:
	docker run --name postgres-simple-bank --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
migrateup:
	migrate -path db/migrations -database "postgresql://root:feb061999@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migrations -database "postgresql://root:feb061999@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migrations -database "postgresql://root:feb061999@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migrations -database "postgresql://root:feb061999@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb --destination db/mock/store.go  github.com/Emmanuel-MacAnThony/simple_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown1 migrateup1 migratedown sqlc test server