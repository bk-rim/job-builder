# Project Installation and Execution Guide

This guide explains how to launch the project. You have the option to start the services (frontend, backend(job-service and job-processor)) separately or together using Docker Compose.

## Prerequisites

- Go 1.18 or later installed on your system
- node 18 or later
- Docker (if you choose to use Docker Compose)

### 1. Clone the Repository

Clone this repository to your machine:

```bash
git clone https://github.com/bk-rim/job-builder.git
cd job-builder
```

## Environment Configuration

The job-service requires a `.env` file containing its specific environment variables. Here's an example content for the `.env` file:

```plaintext
DB_NAME=<<DB_NAME>>
DB_USER=<<USER>>
DB_PASSWORD=<<PASSWORD>>
DB_HOST=<<HOST>>
DB_HOST_DOCKER="db"
DB_PORT="5432"
DB_SSL_MODE="disable"
```
Make sure to replace the values with configurations specific to your environment.

### 2. Compilation of Microservices

For each microservice, use the `go build` command to compile the source code:

```bash
cd job-service
go build -o job-service
cd ../job-processor
go build -o job-processor
cd ..
```

### 3. Execution of Microservices Separately

- job-service

```bash
cd job-service
source .env
./job-service
```

- job-processor

``` bash
cd job-processor
./job-processor
```

### 4. Frontend Execution

At the project root, do the following:

```bash
cd frontEnd
npm install
npm run dev
```
Now you can navigate to `http://localhost:3000`

### 5. Execute project with Docker Compose

If you prefer to use Docker Compose, ensure Docker is installed on your system, then execute the following command at the project root:

```bash
docker-compose build
docker-compose up
```
Now you can navigate to `http://localhost:3000` 
