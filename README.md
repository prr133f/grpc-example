# Simple gRPC service example
gRPC TODO-service with PostgreSQL, docker-compose and GitHub CI

# Download and run
Run following commands:
```sh
git clone git@github.com:prr133f/grpc-todo.git # via SSH
# OR
git clone https://github.com/prr133f/grpc-todo.git # via HTTPS
```
```sh
cd grpc-todo
docker-compose up --build
```

# Configuring
Create file `.env` and set environment variables according in list below:
- POSTGRES_USER
- POSTGRES_PASSWORD
- POSTGRES_DB
- DATABASE_HOST
- APP_PORT

