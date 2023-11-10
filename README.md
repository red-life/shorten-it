# Shorten It
This is a REST API for shortening long URLs, written in Go. It features an in-memory cache for improved performance.

## Features
- REST API for shortening URLs
- In-memory cache to avoid hitting the data store frequently
- Postgres for persistence at the datastore layer
- Using a key generate service instead of random key generation
- Proper error handling 
- Request validating
- Fully dockerized for easy deployment

## Design & Architecture
- Hexagonal Architecture
- Separation of Concerns Principle
- Decorator Pattern
- Repository Pattern

## Endpoints
### Shorten a URL
**Request:**
```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"url": "htps://google.com"}' \
  http://localhost:5000/shorten
```
**Response:**
```json
{"key": "xyz"}
```

### Redirect to the original URL
**Request:**
```bash
curl http://localhost:5000/xyz
```
**Response:**
```
Returns 302 Found with a Location header to the original URL.
```

## Deployment
### Prerequisites
- Docker & Docker Compose
- Make
### Steps
- Copy .env.sample and change it on your own \
`cp .env.sample .env`
- Build the custom image and run the containers \
`make run`
- It will be available at `http://localhost:${PORT}`
## Run Tests
### Prerequisites
- Make
### Steps
- Just run the command below to run all the available tests \
`make run_tests`


## Next Steps
- [ ] Document the API using Swagger
- [ ] Add Logging
