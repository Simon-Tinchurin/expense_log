agg:
	@go build -o bin/aggregator aggregator/main.go
	@./bin/aggregator