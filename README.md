In order to run the project, please do the following:

Create .env file in the root folder of the project (see .env.example file)  
Intall mysql on your os   
Create database with the name that you specified in the .env file  
Run migrations for db: "go run migrations/migration.go"  
From the root folder run: "go run main.go" to start the server  

For tests run: "go test ./..." from the root folder of your app
For api docs run: "swag init" from the root folder of your app
