# Go Simple HTTP Server

## Overview

Go Simple HTTP Server is a lightweight HTTP server built with Go (Golang). It is designed to be easy to use and extend, providing a simple way to handle HTTP requests and responses. This project serves as a foundation for building RESTful APIs and web applications.

## Features

- Lightweight and easy to set up
- Supports middleware for request handling
- Configurable logging using Logrus
- Graceful shutdown on interrupt signals

## Installation

To get started with the Go Simple HTTP Server, follow these steps:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/ralvescosta/go-simple-http-server.git
   cd go-simple-http-server
   ```

2. **Install dependencies:**
   Make sure you have Go installed on your machine. Then, run:
   ```bash
   go mod tidy
   ```

3. **Run the server:**
   You can start the server by running:
   ```bash
   go run main.go
   ```

## Usage

Once the server is running, you can access it at `http://localhost:8080`. You can customize the port and other configurations through environment variables.

## Project Structure

The project is organized as follows:

```
go-simple-http-server/
├── main.go                  # Entry point of the application
├── pkg/                     # Contains application packages
|   ├── configs              # Env Vars configs
│   ├── controllers/         # HTTP request handlers
│   ├── routes/              # Route definitions
│   ├── services/            # Business logic and services
│   └── logger/              # Logging setup and configuration
├── internal/                # Internal application code
│   ├── models/              # Data models
│   └── services/            # Business logic
│
├── properties.local.json     # Local configuration file
└── README.md                # Project documentation
```

## Contributing

Contributions are welcome! If you would like to contribute to the project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Logrus](https://github.com/sirupsen/logrus) for logging
- [Chi](https://github.com/go-chi/chi) for routing
- [Viper](https://github.com/spf13/viper) for Env Vars
