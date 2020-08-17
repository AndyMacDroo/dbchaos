# DB-Chaos #
 ![](https://github.com/AndyMacDroo/dbchaos/workflows/unit-tests/badge.svg)
 
A work in progress tool to wreak controlled havoc on a connected relational datasource.

Currently supports:

* MySQL
* PostgreSQL

## Running Tests

You can run the unit tests with:

```shell script
docker-compose up
```

You can use the following command to have docker-compose shutdown services when the `dbchaostests` have ran to completion, and to return the exit code from the completed tests:

```shell script
docker-compose up --abort-on-container-exit --exit-code-from dbchaostests
```

