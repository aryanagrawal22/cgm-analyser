# CGM Analyser

This project is a Continuous Glucose Monitoring (CGM) analyser. It is written in Go and aims to analyse glucose data and provide meaningful results.

## Setup

To set up the project, follow these steps:

1. Clone the repository.
2. Install the dependencies with `go mod download`.
3. Copy the `.env.example` file to `.env` and fill in your environment variables.
4. Build the application with `go build -o main ./main.go`.
5. Run the application with `./main`.

## Docker

The project includes a `docker-compose` and `Dockerfile` for building a Docker image of the application. To build and run the application, which listens on port 9876, run `docker-compoe up`.

## Cron Jobs

The `StartCron` function in `src/crons/SugarLevel.go` starts a cron job that runs every 5 seconds. The job makes a GET request to an API and logs the response. The API URL and headers are configured with environment variables.

## Environment Variables

The application uses the following environment variables:

- `API_URL`: The URL of the API to request.
- `AUTHORIZATION`: The authorization header for the API request.
- `TIMEZONE`: The timezone header for the API request.
- `API_VERSION`: The API version header for the API request.

These variables are loaded from a `.env` file at startup.
