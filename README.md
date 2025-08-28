<div align="center" id="top"> 
  <img src="./.github/app.gif" alt="Car Log API" />

&#xa0;

  <!-- <a href="https://carlogapi.netlify.app">Demo</a> -->
</div>

<h1 align="center">Car Log API</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/joshmgreen/car-log-api?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/joshmgreen/car-log-api?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/joshmgreen/car-log-api?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/joshmgreen/car-log-api?color=56BEB8">

  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/joshmgreen/car-log-api?color=56BEB8" /> -->

  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/joshmgreen/car-log-api?color=56BEB8" /> -->

  <!-- <img alt="Github stars" src="https://img.shields.io/github/stars/joshmgreen/car-log-api?color=56BEB8" /> -->
</p>

<!-- Status -->

<!-- <h4 align="center">
	🚧  Car Log API 🚀 Under construction...  🚧
</h4>

<hr> -->

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/joshmgreen" target="_blank">Author</a>
</p>

<br>

## :dart: About

The Car Log API is a backend service written in Go that allows users to track and manage vehicles. It provides a RESTful interface for creating, reading, updating, and deleting vehicles, all stored in a PostgreSQL database. The API is containerized using Docker, making it easy for developers to clone the repository and run it without manually installing dependencies.

## :sparkles: Features

:heavy_check_mark: CRUD Operations
Create: Add new vehicles with year, make, model, and mileage.
Read: Retrieve all vehicles or search by model (case-insensitive).
Update: Modify an existing vehicle using its unique ID.
Delete: Remove a vehicle by ID.;\

:heavy_check_mark: Database Integration
Uses PostgreSQL for persistent storage.
Automatically migrates the Vehicle table on startup.
Optionally seeds an initial vehicle if the database is empty.;\

:heavy_check_mark: Containerization
Full Docker and Docker Compose setup.
Developers can start the API and database with a single command: docker-compose up.
No need for manual setup of Go or PostgreSQL locally.;

:heavy_check_mark: Go + Gin Framework
Handlers are implemented using Gin, a lightweight HTTP framework.
Clean separation of concerns: handlers, vehicles service, and database layer.;\

:heavy_check_mark: Error Handling & Validation
Returns standardized JSON responses with success, data, and error fields.
Validates input data for required fields and correct types.
Proper HTTP status codes for all responses (200 OK, 201 Created, 400 Bad Request, 404 Not Found, 500 Internal Server Error).;\

## :rocket: Technologies

The following tools were used in this project:

- **[Go](https://golang.org/)** – Backend programming language
- **[Gin](https://github.com/gin-gonic/gin)** – HTTP web framework for Go
- **[PostgreSQL](https://www.postgresql.org/)** – Relational database
- **[GORM](https://gorm.io/)** – ORM library for Go
- **[Docker](https://www.docker.com/)** – Containerization platform
- **[Docker Compose](https://docs.docker.com/compose/)** – Multi-container orchestration
- **[curl](https://curl.se/)** – Command-line tool for testing API endpoints
- **[Git](https://git-scm.com/)** – Version control

## :white_check_mark: Requirements

Before starting :checkered_flag:, you need to have:

- **[Docker](https://www.docker.com/)** – To run the API and database in containers
- **[Docker Compose](https://docs.docker.com/compose/)** – To orchestrate multi-container setup
- **[Go](https://golang.org/dl/) ≥ 1.24** – Required if building the API locally without Docker
- **[Git](https://git-scm.com/)** – To clone the repository
- **[PostgreSQL](https://www.postgresql.org/) ≥ 14** – Only required if running the database locally (not in Docker)
- **A terminal / command line interface** – To run Docker, Docker Compose, and curl commands
- **curl** or any API testing tool (like **[Postman](https://www.postman.com/)**) – To test API endpoints

## :checkered_flag: Starting

```bash
# 1. Clone the repository
git clone https://github.com/joshmgreen/Car-Log-API.git
cd car-log-api

# 2. (Optional) Create a .env file for database credentials
echo "DB_HOST=localhost
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=carlogdb
DB_PORT=5432" > .env

# 3. Build and start the Docker containers
docker-compose up --build

# 4. Stop the containers
docker-compose down

# 5. View logs
docker-compose logs -f

# 6. Test the API endpoints using curl

# Get all vehicles
curl -X GET http://localhost:8080/vehicles

# Get a vehicle by model
curl -X GET http://localhost:8080/vehicles/model/Civic

# Add a new vehicle
curl -X POST http://localhost:8080/vehicles \
  -H "Content-Type: application/json" \
  -d '{"Year":2023,"Make":"Toyota","Model":"Corolla","Mileage":1500}'

# Update an existing vehicle
curl -X PUT http://localhost:8080/vehicles/1 \
  -H "Content-Type: application/json" \
  -d '{"Year":2025,"Make":"Honda","Model":"Civic Type R","Mileage":5000}'

# Delete a vehicle
curl -X DELETE http://localhost:8080/vehicles/1
```

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.

Made with :heart: by <a href="https://github.com/joshmgreen" target="_blank">Josh Green</a>

&#xa0;

<a href="#top">Back to top</a>
