# ZEXD - URL Shortening Service

**ZEXD** is a URL shortening service designed for speed along with reliability, scalability, and monitoring in mind. This project uses a Go-based backend, with PostgreSQL as the database, DragonflyDB for caching, and integrated monitoring via Prometheus and Grafana.

## Features

- URL shortening and redirection
- Click tracking and analytics
- URL expiration management
- Health monitoring with Prometheus and Grafana
- API documentation via Swagger

## Benchmarks

The performance of ZEXD was tested using the [k6](https://k6.io/) load testing tool. The tests were run on a Lenovo Ideapad Flex 5 with an 11th Gen Intel Core i7 processor and 16GB of RAM.

- Tests for Url Creation
  - Test 1
    - 1 VU for 30 seconds
    - 14845 requests
    - Avg. response time: 1.92ms
  - Test 2
    - 100 VU for 2 minutes
    - 183464 requests
    - Avg. response time: 65.25ms
  - Test 3
    - 200 VU for 2 minutes
    - 356080 requests
    - Avg. response time: 67.23ms

- Tests for Url Redirection
  - Test 1
    - 1 VU for 30 seconds
    - 10713 requests
    - Avg. response time: 2.7ms
  - Test 2
    - 100 VU for 2 minutes
    - 165830 requests
    - Avg. response time: 72.03ms
  - Test 3
    - The server was not able to handle the load of 200 VUs for 2 mins.

See the [benchmark results](loadtesting/benchmark.md) for the performance of ZexD.

## Prerequisites

Make sure you have the following installed:

- **Docker** and **Docker Compose**
- **Go 1.23** or later (optional, if you wish to run or develop outside Docker)

## Getting Started

### Cloning the Repository

```bash
git clone https://github.com/AvaterClasher/zexd-backend.git
cd zexd-backend
```

### Environment Variables

Create an `.env` file in the root directory based on the `.env.example` provided:
Change with whatever DB or redis service you are using if not using docker compose.

```bash
DB_HOST=postgres
DB_PORT=5432
DB_USER=zexd
DB_PASSWORD=zexd
DB_NAME=zexd_db
DB_SSL=disable
TABLE_NAME=shortened_url # dont change this
EXPIRY_TIME=1440  # Expiration time in minutes
REDIS_DOMAIN=redis://user@dragonflydb:6379/
SERVER_DOMAIN=http://localhost:8080/
```

### Docker Compose Setups

#### Local Development

This setup is for local development and testing. Any changes made to the code will be reflected here.
The environment variables are already set in the docker compose file.
To start the service in a local development environment, use:

```bash
docker-compose -f compose.local.yml up --build
```

#### Production Setup

This uses a pre-built image from Docker Hub. This setup is for production use.
The environment variables are already set in the docker compose file.
For a production-ready setup:

```bash
docker-compose -f compose.prod.yml up -d --build
```

This setup configures all services (backend, PostgreSQL, DragonflyDB, Prometheus, and Grafana) with production-level configurations.

### Running the Containers Individually

If you want to run the services independently, here are some examples:

#### Backend (ZEXD)

```bash
docker run -p 8080:8080 --env-file .env avaterclasher/zexd-backend:v1
```

### Swagger API Documentation

- **Production**: Hosted Swagger documentation is available at [https://zexd.onrender.com/swagger/index.html](https://zexd.onrender.com/swagger/index.html)
- **Local**: Once the service is running locally, access Swagger documentation at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

## Monitoring with Prometheus and Grafana

This project uses Prometheus and Grafana for monitoring. When started, these services will be available at:

- **Prometheus**: [http://localhost:9090](http://localhost:9090)
- **Grafana**: [http://localhost:3000](http://localhost:3000)

To log into Grafana, use:

- **Username**: `admin`
- **Password**: `admin`

You can find the Grafana dashboard JSON file in `grafana_dashboard.json` for pre-configured metrics related to this project.

## Dockerfile Information

The Dockerfile is a multi-stage build with two stages:

1. **Builder**: Builds the Go application.
2. **Alpine**: Copies the compiled application and runs it on a minimal Alpine image for lightweight production usage.

### Build and Run Manually

To build and run the image manually:

```bash
docker build -t zexd-backend .
docker run -p 8080:8080 --env-file .env zexd-backend
```

## API Endpoints

The API endpoints are as follows:

| Method | Endpoint | Body | Description |
| --- | --- | --- | --- |
| POST | /api/create | {"url": "https://www.youtube.com", "user_id": "user1"} | Generates a shortened URL |
| POST | /api/delete | {"url": "https://www.youtube.com"} | Deletes the shortened URL |
| GET | /api/list/{user_id} || Returns the urls of the user |
| GET | /{shortened_url} || Redirects to the original URL |
| GET | /health || Returns the health status of the server |

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any suggestions or improvements.

## Maintainer

- Github: [Avater Clasher](https://github.com/AvaterClasher)
- Linkedin: [Soumyadip Moni](https://www.linkedin.com/in/soumyadip-moni/)

## License

This project is licensed under the [MIT License](LICENSE).
