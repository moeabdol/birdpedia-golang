postgres:
	docker container run -dt --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres postgres:12-alpine

start_db:
	docker container start postgres

stop_db:
	docker container stop postgres

create_db:
	docker exec -it postgres createdb --username=postgres --owner=postgres birds_db

drop_db:
	docker exec -it postgres dropdb --username=postgres birds_db

create_test_db:
	docker exec -it postgres createdb --username=postgres --owner=postgres birds_test_db

drop_test_db:
	docker exec -it postgres dropdb --username=postgres birds_test_db

migrate_up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/birds_db?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/birds_db?sslmode=disable" -verbose down

migrate_test_up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/birds_test_db?sslmode=disable" -verbose up

migrate_test_down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/birds_test_db?sslmode=disable" -verbose down

clean:
	go clean -cache

test:
	go test -v -cover ./...

.PHONY: postgres start_db stop_db create_db drop_db create_test_db drop_test_db migrate_up migrate_down migrate_test_up migrate_test_down clean test
