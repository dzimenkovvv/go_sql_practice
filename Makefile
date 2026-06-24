include .env
export

service-run:
	@go run main.go

service-deploy:
	docker compose up -d application

service-undeploy:
	docker compose down application

service-postgres-up:
	docker run -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -p 5432:5432 -v ./out/pgdata:/var/lib/postgresql -d postgres:18.4-bookworm

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down
