# Deck API

- [Deck API](#deck-api)
  - [Project Overview](#project-overview)
    - [Package Structure](#package-structure)
    - [Dependencies](#dependencies)
    - [Logic Separation](#logic-separation)
    - [Build and Run](#build-and-run)
  - [⭐ **Principles for dealing with development of this repository**](#principles-for-dealing-with-development-of-this-repository)

## Project Overview

This project provides a simple RESTful API for working with a deck of cards. The API allows users to create a deck, shuffle a deck, and draw cards from a deck.

The project is structured as follows:

```md
.
├── cmd
│ └── deck
│   └── main.go
├── pkg
│ ├── deck
│ │ ├── deck.go
│ │ ├── deck_test.go
│ │ ├── draw.go
│ │ ├── draw_test.go
│ │ ├── open.go
│ │ └── open_test.go
│ ├── server
│ │ ├── server.go
│ │ └── server_test.go
│ └── service
│   ├── deck.go
│   └── deck_test.go
├── go.mod
└── README.md
```

### Package Structure

- `cmd/deck`: This package contains the `main.go` file which is the entry point of the application. It initializes the necessary dependencies, creates an instance of the server, and starts the server to listen for incoming requests.

- `pkg/deck`: This package provides the core functionality for working with a deck of cards. It includes methods for creating a deck, shuffling a deck, and drawing cards from a deck. The package also includes corresponding test files.

- `pkg/server`: This package provides the server layer of the application. It includes the HTTP server initialization and routing definitions. The package also includes corresponding test files.

- `pkg/service`: This package provides the service layer of the application. It includes methods for extracting, processing, and validating URL parameters and request bodies. Once validated, requests are forwarded to the appropriate methods in separate packages.

### Dependencies

This project depends on the following external libraries:

- `gin-gonic/gin`: A popular HTTP router and dispatcher for building RESTful APIs.
- `stretchr/testify`: A testing toolkit for writing unit tests and assertions in Go.
- `google/uuid`: A package for generating and parsing UUIDs in Go.

### Logic Separation

```mermaid
graph TD
    A[Main Package Layer] --> B[Server Package Layer]
    B --> C[Service Package Layer]
    C --> D[Dedicated Package with Separate Logic]
```

### Build and Run

- To build and run the project locally, use the following commands:

  - Build
    ```terminal
    go build ./cmd/deck
    ```
  - Run
    ```terminal
    ./deck
    ```

- To build and run the project locally, use the following commands:
  - Build
    ```terminal
    docker build -t deck-api .
    ```
  - Run
    ```terminal
    docker run -p 8080:8080 deck-api
    ```

---

## Principles for dealing with development of this server

1. Initialization of major dependencies should be performed in the main package. This ensures that the server is properly configured before it starts processing requests.
2. API routing should be carefully defined, including the endpoints that can be accessed within the server layer. Different group handlers should be registered to ensure that requests are properly handled and routed to the appropriate services.
3. The service layer should be responsible for extracting, processing, and validating URL parameters and request bodies. Once validated, requests should be forwarded to the appropriate methods in separate packages.
4. To handle incoming requests with validated parameters, an appropriate package should be built to process these requests efficiently and effectively.
