# Card Activity Notification Service

A simple notification service that processes card activity events and sends notifications. Built with Go using only the standard library.

## Prerequisites

- Go 1.20 or higher
- Git (for version control)

## Installation

1. Clone the repository:
```bash
git clone <your-repository-url>
cd notification-service
```

2. Initialize Go module:
```bash
go mod init notification-service
```

## Building the Application

Build the application using the Go compiler:

```bash
go build -o notification-service
```

## Running the Application

There are two ways to run the application:

1. Using `go run`:
```bash
go run main.go
```

2. Using the built executable:
```bash
./notification-service
```


## Running Tests

Run all tests with verbose output:
```bash
go test -v ./...
```

Run tests for a specific package:
```bash
go test -v ./store    # Run store package tests
go test -v ./worker   # Run worker package tests
```
**POST /event**

Example request:
```bash
curl -X POST http://localhost:8080/event \
-H "Content-Type: application/json" \
-d '{
  "orderType": "Purchase",
  "sessionId": "29827525-06c9-4b1e-9d9b-7c4584e82f56",
  "card": "4433**1409",
  "eventDate": "2023-01-04T13:44:52.835626Z",
  "websiteUrl": "https://amazon.com"
}'
