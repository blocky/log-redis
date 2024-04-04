# Log Redis Set Events
A small go service for logging redis key values changed with `set`.

# Dependencies

### GO
This project uses Go 1.22 (installation instructions [here](https://go.dev/doc/install))

### Mage
`mage` is used for build tasks (installation instructions [here](https://magefile.org/))

### Redis

This project uses Redis for storing key value pairs. You can install Redis however
you like. If you have docker installed, you can run the following command to start a Redis container:

```bash
docker run --name blocky-redis-dev -p 6379:6379 -d redis
```

# Running the server

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
The server will start on port 8080.

# Binaries

To build the executable, run the following command
(replace 'linux' and 'amd64' with your desired OS and architecture):
```bash
GOOS=linux GOARCH=amd64 mage build
```
For available OS and architecture options, see [here](https://golang.org/doc/install/source#environment)

Alternatively, you can run the following command, which depends only on golang, and cross compiles a number of binaries:
```bash
./cross_compile.sh
```
