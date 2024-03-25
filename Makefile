DB_URL = postgres://root:mysecret@localhost:5432/la_candela?sslmode=disable
ntwk:
	docker network create la_candela-network

rm_ntwk:
	docker network rm la_candela-network

postgres:
	docker run --name postgres15 --network la_candela-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root la_candela

dropdb:
	docker exec -it postgres15 dropdb la_candela

startdb:
	docker start postgres15

stpdb:
	docker stop postgres15

rmdb:
	docker rm postgres15 && docker network rm la_candela-networ

migrate_up:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrate_up1:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1

migrate_down1:
	migrate -path db/migration -database "${DB_URL}" -verbose down 1

migrate_down:
	migrate -path db/migration -database "${DB_URL}" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

server:
	go run main.go

up:
	docker compose up
down:
	docker compose down


proto:
	rm -f pb/*.go
	rm -f doc/swagger/*json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=la-candela-backend-rpc \
	proto/*.proto && \
	statik -src=./doc/swagger -dest=./doc -f

redis:
	docker run --name redis --network la_candela-network -p 6379:6379 -d redis:7-alpine

startrd:
	docker start redis


.PHONY: proto up down createdb dropdb redis db_migration_down1