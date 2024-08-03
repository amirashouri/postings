DB_URL=postgresql://root:secret@localhost:5432/postings?sslmode=disable

network:
	docker network create postings_network

postgres:
	docker run --name postgres --network postings_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root postings

dropdb:
	docker exec -it postgres dropdb postings

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

tailwind-watch:
	./tailwindcss -i ./public/css/input.css -o ./public/css/style.css --watch

templ-generate:
	templ generate

tailwind-build:
	./tailwindcss -i ./public/css/input.css -o ./public/css/style.min.css --minify

templ-watch:
	templ generate --watch

tailwindcss-exec:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod +x tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 tailwindcss

.PHONY: network postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc tailwind-watch templ-generate templ-watch tailwindcss-exec tailwind-build new_migration