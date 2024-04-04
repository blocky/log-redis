# Log Redis Set Events
A small go service for logging redis key values changed with `set`.

# Dependencies

### GO

This project uses Go 1.22 (installation instructions
[here](https://go.dev/doc/install)).

### Mage

The program `mage` is used for build tasks (installation instructions
[here](https://magefile.org/)).

### Redis

This program is designed to watch for specific Redis events. For development and
testing, you will want to have a Redis instance running (installation
instructions [here](https://redis.io/docs/install/install-redis/)).

If you have docker installed, you can run the following command to start a Redis
container:

```bash
docker run --name blocky-redis-dev -p 6379:6379 -d redis
```

# Running the program

Below, we assume that you have a Redis server running on `localhost:6379` with
no password and would like to write to database 0.

The server relies on the following environment variables to configure redis:
```bash
REDIS_ADDRESS=":6379"
REDIS_PASSWORD=""
REDIS_DATABASE=0
```

To run the server, run the following command:
```bash
 REDIS_ADDRESS=":6379" REDIS_PASSWORD="" REDIS_DATABASE=0 go run .
```

# Binaries

To build the executable, run the following command:
```bash
mage build
```
If you wish to cross compile, you can specify your desired OS and architecture:
```bash
GOOS=linux GOARCH=amd64 mage build
```
For available OS and architecture options, see
[here](https://golang.org/doc/install/source#environment)

Alternatively, you can run the following command to cross compiles a number of
binaries:
```bash
mage buildAll
```

To clean up the dist directory run
```bash
mage clean
```
