build:
	go build -o bin/piglatin ./api

run-db:
	docker compose up

run: build
	bin/piglatin
