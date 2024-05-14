# API for expense manager
### Project outline

1. Create new docker container with empty data base:
```
docker run --name my_postgres_container -e POSTGRES_DB=<YOUR_DB_NAME> -e POSTGRES_USER=<YOUR_USER> -e POSTGRES_PASSWORD=<YOUR_PASSWORD> -p 5432:5432 -d postgres
```
2. Add package to load environment variables
```
go get github.com/joho/godotenv
```