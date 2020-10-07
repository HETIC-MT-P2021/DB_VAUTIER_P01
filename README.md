# DB_VAUTIER_P01

A simple API to get informations about a shop and train on SGBD.

## Development process

Once the docker is started you'll need to import the database file into the docker database container.

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

## API documentation

A documentation is available for the API on the following link : 
https://documenter.getpostman.com/view/7544320/TVRhdA4X