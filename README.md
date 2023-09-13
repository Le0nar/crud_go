For run app create a ".env" file with password like a ".example.env" file

For run docker postgresql:
1. Pull docker image of postgres
docker pull postgres
2. Run container
docker run --name=news-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

Example of creating a migration scheme:
migrate create -ext sql -dir ./schema -seq init

Apply migration up scheme
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

Postresql via docker:

1.docker exec -it {container_id} /bin/bash
2.root@3d52773f7ff2:/# psql -U postgres
3. postgres=# \d