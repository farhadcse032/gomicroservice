
postgres:
	docker run --name micropostgres -p 5436:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
createdb:
	docker exec -it micropostgres createdb --username=root --owner=root  simple_bank
checkdb:
	docker exec -it micropostgres /bin/sh
dropdb:
	docker exec -it micropostgres dropdb --username=root --owner=root  simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5436/simple_bank?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5436/simple_bank?sslmode=disable" -verbose down 
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY: postgres createdb checkdb dropdb migrateup migratedown sqlc test