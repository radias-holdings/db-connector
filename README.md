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

and then to use the connector:

```go
db, err := connector.NewConnection(user, password, address, dbname)
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

# License

MIT License

Copyright (c) [year] [fullname]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.