db:
	docker compose up --detach

db-down:
	docker compose down

tests:
	go test -v ./tests
