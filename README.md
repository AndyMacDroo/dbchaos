# DB-Chaos #
 ![](https://github.com/AndyMacDroo/dbchaos/workflows/unit-tests/badge.svg)
 
A work in progress tool to wreak controlled havoc on a connected relational datasource.

Currently supports:

* MySQL
* PostgreSQL

## Docker Image

The application has a docker image available here:

```
docker pull andymacdonald/dbchaos
```

## Running Tests

You can run the unit tests with:

```shell script
docker-compose up
```

You can use the following command to have docker-compose shutdown services when the `dbchaostests` have ran to completion, and to return the exit code from the completed tests:

```shell script
docker-compose up --abort-on-container-exit --exit-code-from dbchaostests
```

## Environment Variables

The following environment variables can be used to configure the application.

`CHAOS_MODE` - What chaos mode to use, one of `CONNECTION_LEAK` or (tba) `QUERY_BURST`.
`DATABASE_HOST` - Host/IP of a given database.
`DATABASE_PORT` - Port for a given database.
`DATABASE_PASSWORD` - Password for a given database.
`DATABASE_USERNAME` - Username for a given database.
`DATABASE_TYPE` - One of `MySQL` or `PostgreSQL`.
`DATABASE_NAME` - The name of the database.
`MAX_CONNECTIONS_TO_LEAK` - Number of connections to leak.
`CONNECTION_CREATION_WAIT_MS` - Wait in milliseconds between creation of connections.
`CONNECTION_LEAK_HOLD_MS` - Wait time in milliseconds to keep hold of leaked connections.

## Kubernetes Deployment

The service can be ran as a kubernetes deployment and an example can be found below:

* [`examples/deployment.yml`](examples/deployment.yml)