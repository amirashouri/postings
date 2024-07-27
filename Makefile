DB_URL=postgresql://root:secret@localhost:5432/postings?sslmode=disable

network:
	docker network create postings_network

postgres:
	docker run --name postgres --network postings_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root postings

dropdb:
	docker exec -it postgres dropdb postings

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

.PHONY: network postgres createdb dropdb migrateup migrateup1 migratedown migratedown1