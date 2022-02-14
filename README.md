<h1 align="center">
  <img src="https://i.ibb.co/YpvKyrK/Solution-1-Time-Stamp.png" alt="Logo" width="600">
</h1>

## Features

- Unix and UTC values from date/unix

## Tech Stack

**Server:** Golang (Gin)

## Run Locally

Here is the steps to run it with `golang`

```bash
# Move to directory
$ cd workspace

# Clone this repository
$ git clone https://github.com/madjiebimaa/fcc-timestamp-ms.git

# Move to project
$ cd fcc-timestamp-ms

# Run the application
$ go run main.go
```

Here is the steps to run it with `docker-compose`

```bash
# Move to directory
$ cd workspace

# Clone this repository
$ git clone https://github.com/madjiebimaa/fcc-timestamp-ms.git

# Move to project
$ cd fcc-timestamp-ms

# Download and setup the image
$ docker-compose up -d

# Run the application
$ go run main.go
```

## Running Tests

To run tests and get the percentage of code coverage, run the following command

```bash
  go test ./... -cover
```

## Lessons Learned

- How to create API with Golang (Gin)
- How to create middleware
- How to create gin handler
- How to create unit test for handler
- How to parse time to string

## API Reference

[Test API with REST Client Extension](https://github.com/madjiebimaa/fcc-timestamp-ms/tree/main/docs/apis)
