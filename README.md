In order to run the project, please do the following:

Create .env file in the root folder of the project (see .env.example file)  
Intall mysql on your os   
Create database with the name that you specified in the .env file  
Run migrations for db: "go run migrations/db_migration.go"  
Run go mod tidy (this will get all the dependencies for the your project)  
Run "swag init" (this will generate docs for the project)  
To start the server run: "go run main.go"  

For tests run: "go test ./..." from the root folder of your app  
