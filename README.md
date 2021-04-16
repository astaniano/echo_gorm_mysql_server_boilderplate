in order to run the following project follow the next steps:

create .env file in the root folder of the project (see .env.example file)
in the "helpers/load_env.go" change path to your .env file that you have created
intall mysql on your os
create database "jph"
run migrations for db: "go run migrations/migration.go"
from the root folder run: "go run main.go" to start the server

for tests run: "go test ./..." from the root folder of your app