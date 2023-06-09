postgres: 
	docker run --name postgresb -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

start:
	docker start 4a9eb2d61aa2

createdb:
	docker exec -it postgresb createdb --username=root --owner=root adventure_life_simple_bank

dropdb:
	docker exec -it postgresb dropdb adventure_life_simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/adventure_life_simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/adventure_life_simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/wanderlust-create/adventure-life-simple-bank/db/simple_bank_db Store

.PHONY: postgres start createdb dropdb migrateup migratedown sqlc test server mock