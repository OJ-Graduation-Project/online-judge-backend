.PHONY: init_db
run:
	@go run cmd/main.go
init_db:
	@mongod
	