# Educado Transcoding Service

## Overview
This service is responsible for handling GCP trafic (uploading, downloading, etc.), transcoding and streaming.

## Getting Started

### Prerequisites
- .env file with the following variable(s):
```
GOOGLE_APPLICATION_CREDENTIALS=<path to service account key>
```
- .gcp_credentials.json (service account key).
- Go should be installed on your machine.
- (Optional) GoLand IDE.
### Installation
- Open project in GoLand.
- Insert .env file in the root of the project.
- Insert .gcp_credentials.json in the root of the project.
- Run the following command in the terminal:
```
go run ./cmd/server/main.go
```
- The server should be up and running.

### Quickstart
- Run the following command in the terminal:
```
go run ./cmd/server/main.go
```

### Docker
- Build the image:
```
docker build -t <image-name> .
```
- Run the image:
```
docker run -p 8080:8080 <image-name>
```
The service should be up and running on port 8080.

## Usage
- This service is desgined to work with Educado's backend.

## API Reference
- Description of routes (to be added)

## Acknowledgments
- Shout out to Carl Ryskov Aagesen for being a great mentor.
