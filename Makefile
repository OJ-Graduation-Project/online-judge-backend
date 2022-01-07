.PHONY: init_db run stop_db init_db populate_db
run:
	@go run cmd/main.go

# starts the database, you must stop processes running on 127.0.0.1/27017 befre running this command
start_db: 
	@mkdir -p database && mongod --dbpath ./database/

# stops database, as far as i know this works for linux only
stop_db:
	@mongod --shutdown --dbpath database

# populates database with json files found in ./scripts/database/data
populate_db:
	@mongo scripts/database/mongo.js

# deletes onlineJudgeDB
drop_db:
	@mongo onlineJudgeDB --eval "db.dropDatabase()"