
DB_URL=postgres://root:password@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name postgres17 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17.4-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down -all

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
cleardb:
	make migratedown && make migrateup
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go simple-bank/db/sqlc Store
proto:
	rm pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto
evans:
	evans --host localhost --port 9090 --path proto -r repl
.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test cleardb server mock proto evans
