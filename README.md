# DB Connector Package

This package provides a standardised approach for connecting to a MySQL database.

## Import the package

To import this package, navigate to the folder of your project in the command line then run:

```bash
export GOPRIVATE=github.com/fzer0zer0
```

This is to signal to the system that it should not validate the private package against the public Go checksum database otherwise it will return a 404 error.

In your main package, add the following imports:

```text
"github.com/fzer0zer0/db-connector"
"github.com/fzer0zer0/logger"
```

then run in the command line:

```bash
go get github.com/fzer0zer0
```

to use the connector, add the following variable to your project:

```go
db, err := db-connector.NewConnection()
```

## Managing environment variables

This package makes use of the [GoDotEnv library](https://github.com/joho/godotenv).

Your project will need a `.env` file with values for:

```text
DB_USER=<MySQL DB username>
DB_NAME=<MySQL DB name>
DB_ADDRESS=<address of MYSQL DB> (127.0.0.1:3306 for localhost)
```

 as well as a `.env.test` file (added to `.gitignore`) that holds:

- DB_USER
- DB_NAME
- DB_PASSWORD (be careful to not commit sensitive information)
- DB_ADDRESS (`127.0.0.1:3306` for localhost)

Set the DB_PASSWORD for your local environment

```bash
export DB_PASSWORD=<insert MySQL DB password>
```

add the following import to your main package:

```text
"github.com/joho/godotenv"
```

then run in the command line:

```bash
go get github.com/joho/godotenv
```

Add the following init function above your main function in your main package, this will use the variables in the `.env` file and the DB_PASSWORD set above:

```go
func init() {
 if err := godotenv.Load(); err != nil {
  log.Error("No .env file found", "rsp", err)
 }
}
```

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
