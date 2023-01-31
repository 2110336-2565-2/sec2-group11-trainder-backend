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
