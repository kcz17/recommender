# Recommender

[![](https://images.microbadger.com/badges/image/kcz17/recommender.svg)](http://microbadger.com/images/kcz17/recommender "Get your own image badge on microbadger.com")

A Sock Shop service that provides a recommendation by connecting to the catalogue service.

### To build this service

In order to build the project locally you need to make sure that dependencies are installed. Once that is in place you
can build by running:

```
go mod download
go build -o recommender
```

The result is a binary named `recommender`, in the current directory.

#### Docker
`docker-compose build`

### To run the service on port 8080

#### Go native

If you followed to Go build instructions, you should have a "recommender" binary in $GOPATH/src/github.com/kcz17/recommender/cmd/newssvc/.
To run it use:
```
./recommender --port 8080
```

#### Docker
`docker-compose up`

### Check whether the service is alive
`curl http://localhost:8080/health`

### Use the service endpoints
`curl http://localhost:8080/recommender`

### Releasing
- `docker build -t kcz17/recommender:[VERSION] -f docker/recommender/Dockerfile .`
- `docker build -t kcz17/recommender-db:[VERSION] docker/recommender-db`
