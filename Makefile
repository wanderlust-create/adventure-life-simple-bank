postgres: 
	docker run --name postgresb -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgresb createdb --username=root --owner=root adventure_life_simple_bank

dropdb:
	docker exec -it postgresb dropdb adventure_life_simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/adventure_life_simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/adventure_life_simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown