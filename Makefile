up:
	docker compose up --build

down:
	docker-compose down

rm:
	docker compose down --remove-orphans --volumes

db:
	pgcli postgres://postgres:postgres@localhost:5337/spamtube_db