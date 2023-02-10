# API for Trainder

## Requirements

- [Go](https://go.dev)
- [MongoDB](https://www.mongodb.com/)

## Development

1. Clone the repository
2. Set up MongoDB (Both Local and MongoDB Atlas can be used)
3. Create `.env` file at the root of this repository with:

```
MONGO_URI=mongodb://<YOUR MONGO CONNECTION STRING>
MONGO_DATABASE_NAME=<YOUR DATABASE NAME>
TOKEN_HOUR_LIFESPAN=<TIME IN HOUR FOR TOKEN LIFESPAN>
API_SECRET=<YOUR API SECRET>
```

4. `go mod tidy` to get all the Go requirements
5. `go run .` to start the API

The API will be available at [localhost:8080](http://localhost:8080)

## Generating the Documentation

This repo use swagger as the documentation. The generation process of the documentation is a follows. More info at [https://github.com/swaggo/swag](https://github.com/swaggo/swag)

1. Add comments to your API source code.
2. Install swag cli tools by:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

3. Run `swag init` to generate the documentation.
4. The documentation will be available at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Test
For now, if you want to run tests, simply execute this command
```sh
go test .
```

## Development using Docker

With docker the development process can be easier. The database and the go environment will be automatically set up for you.
You will need docker and docker-compose. The processes for running using docker as as follows.

1. Build the images

```sh
docker-compose build
```

2. Start the service by

```sh
docker-compose up
```

The API will be available at [localhost:8080](http://localhost:8080)

### Notes about using docker

When you build the images all of the file of this repository will be copied to the container. Then the go compiler will be run in the container.
Which means if you made change to the code **you must re run** `docker-compose build`
