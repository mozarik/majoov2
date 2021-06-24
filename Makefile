migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir postgres/migration $$name

migrate:
	migrate -source file://postgres/migration \
	-database postgres://root:root@localhost:5432/root?sslmode=disable up


rollback:
	migrate -source file://postgres/migration \
	-database postgres://root:root@localhost:5432/root?sslmode=disable down

drop:
	migrate -source file://postgres/migration \
	-database postgres://root:root@localhost:5432/root?sslmode=disable drop

sqlc:
	sqlc generate