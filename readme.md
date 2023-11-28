# Rest Api with GO and Gin

This project is a simple rest api with GO and Gin.
I used a Postgres database and Docker to run the project.
In the database, I stored only the user information.
The files are stored in binary format, allowing also malformed jsons. 
In fact, I didn't want to check if the json is malformed or not.
I wanted to allow all the files
All the main variables are stored in the .env file in the root of the project.
I used the environment variables for a better security.

## Requirements
- Go 1.16
- Docker
- Docker-compose

## Run
To run the project execute build.sh script in the root of the project.
Othwerwise, you can run the project manually.
In the root of the project execute the following commands:
```bash
docker-compose up -d
go build -o ./bin/myapp ./cmd/app
./bin/myapp

```
