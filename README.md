For run docker postgresql:
1. Pull docker image of postgres
docker pull postgres
2. Run container
docker run --name=news-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

Example of creating a migration scheme:
migrate create -ext sql -dir ./schema -seq init

Apply migration up scheme
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up