# DB Connector Package

This package provides a standardised approach for connecting to a MySQL database.

## Installation

In the command line, navigate to the folder of your project then run:

```bash
export GOPRIVATE=github.com/fzer0zer0
```

This is to signal to the system that it should not validate the private package against the public Go checksum database otherwise it will return a 404 error.

then run in the command line:

```bash
go get github.com/fzer0zer0/db-connector
```

## Usage

In your main package, add the following import:

```text
"github.com/fzer0zer0/db-connector"
```

and then to use the connector, add the following variable to your project:

```go
db, err := connector.NewConnection()
```

## Managing environment variables

Ensure your env supports the env vars below:

```text
DB_USER=<MySQL DB username>
DB_NAME=<MySQL DB name>
DB_ADDRESS=<address of MYSQL DB> (127.0.0.1:3306 for localhost)

Set the DB_PASSWORD for your local environment

```bash
export DB_PASSWORD=<insert MySQL DB password>
```

add the following import to your main package:

## Running tests in this package

To run the tests in this package, add a `.env.test` file to this package and set the following values for your local MySQL DB:

```text
DB_USER=<MySQL DB Username>
DB_PASSWORD=<MySQL DB Password>
DB_NAME=<MySQL DB Name>
```

then run in the command line:

```text
make test
```
