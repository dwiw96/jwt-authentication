run:
	go run main.go

databasepg:
	docker run --rm --name postgres_jwt -p 5432:5432 -e POSTGRES_USER=db -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=authentication postgres

dockerexec:
	docker exec -it postgres_jwt psql -U db authentication