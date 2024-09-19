db:
	docker compose up --detach

db-down:
	docker compose down

test:
	go test -v ./tests
