# DB_VAUTIER_P01

A simple API to get informations about a shop.

## Development process

With docker:

```sh
docker-compose up -d

docker-compose exec app /bin/sh

    go build -o main .

    # Start the server
    ./main
```

Without docker:

```sh
go build -o main .

# Start the server
./main
```