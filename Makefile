agg:
	@go build -o bin/aggregator aggregator/main.go
	@./bin/aggregator

seed:
	@go build -o bin/seed seed/main.go
	@./bin/seed

.PHONY: agg seed