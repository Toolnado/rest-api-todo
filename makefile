all:dockerPostgres sleep migrate

dockerPostgres:
	docker run --name todo_db -e POSTGRES_PASSWORD='qwerty' --rm -d -p 5436:5432 postgres

sleep:
	sleep 5

migrate:
	migrate -path schema -database 'postgres://postgres:qwerty@localhost:5436?sslmode=disable' up