In order to run the following project follow the next steps:

Create .env file in the root folder of the project (see .env.example file).  
In the "helpers/load_env.go" change path to your .env file that you have created.  
Intall mysql on your os.  
Create database "jph".  
Run migrations for db: "go run migrations/migration.go".  
From the root folder run: "go run main.go" to start the server.  

For tests run: "go test ./..." from the root folder of your app