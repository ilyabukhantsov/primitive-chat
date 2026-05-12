test:
	go test ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

up:
	docker compose up

upd:
	docker compose up -d

build:
	docker compose build

down:
	docker compose down

restart:
	docker compose down
	docker compose up -d

logs:
	docker compose logs -f

ps:
	docker compose ps

app:
	docker compose exec app sh

postgres:
	docker compose exec postgres psql -U postgres

clean:
	docker compose down -v
