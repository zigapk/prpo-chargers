# prpo-chargers microservice

This repository contains a source code for chargers microservice used in a demo project developed under PRPO subject at UNI LJ.

## Setup

Run a development database via Docker:

```shell
docker run -p 5433:5432 \
    -e POSTGRES_USER=prpo-chargers \
    -e POSTGRES_PASSWORD=rootroot \
    -v prpo-chargers:/var/lib/postgresql/data \
    -d postgres:14
```

Install dependencies and compile:

```shell
go mod download
make docs
make bindata
make build
```

Run the server:

```shell
./chargers serve
```

System behaviour can be configured through [configs/config.ini](configs/config.ini).

## Documentation

Generate docs using
```bash
swag init -d cmd/chargers/internal/handle
```

The docs are then available at `/docs`.
