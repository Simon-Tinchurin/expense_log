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
3. Install library for Postgresql connection
```
go get -u github.com/jmoiron/sqlx
go get github.com/lib/pq
```
4. Install library for uuid generation
```
go get github.com/google/uuid
```
5. Install Gin framework for API development
https://gin-gonic.com/docs/quickstart/
```
go get -u github.com/gin-gonic/gin
```

### TODO (Project Plan)

- [ ] Integrate interfaces for stores
- [ ] Integrate logging middleware?
- [ ] Create GET all expenses route