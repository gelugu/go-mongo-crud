db:
	docker compose up --detach

db-down:
	docker compose down

tests: db
	go test -v ./tests
	make db-down
