# Project Overview

This project consists of two primary components:

1. **Email Indexer**: Processes and indexes emails from a specified directory.
2. **Email Search API**: Provides an API for querying and managing indexed email data.

## Getting Started

### Prerequisites

Before running the project, ensure you have:

- **Golang**
- A running instance of **ZincSearch**

### Configuration

Set the following environment variables for both components:

- `ZINCSEARCH_URL` - The URL of the ZincSearch instance
- `ZINCSEARCH_USER` - Username for authentication
- `ZINCSEARCH_PASSWORD` - Password for authentication
- `ZINCSEARCH_INDEX_NAME` - Name of the index used for storing emails

For the Email Search API, an additional optional environment variable can be configured:

- `ALLOWED_ORIGINS` - Defines allowed CORS origins (default: `http://localhost:5173`)

### Running the Email Indexer

To start the email indexer, run the following command:

```sh
ZINCSEARCH_URL=YOUR_ZINCSEARCH_URL \
ZINCSEARCH_USER=YOUR_ZINCSEARCH_USER \
ZINCSEARCH_PASSWORD=YOUR_ZINCSEARCH_PASSWORD \
ZINCSEARCH_INDEX_NAME=YOUR_ZINCSEARCH_INDEX_NAME \
go run ./cmd/indexer {folderPath}
```

### Running the Email Search API

To launch the API server, use the command below:

```sh
ZINCSEARCH_URL=YOUR_ZINCSEARCH_URL \
ZINCSEARCH_USER=YOUR_ZINCSEARCH_USER \
ZINCSEARCH_PASSWORD=YOUR_ZINCSEARCH_PASSWORD \
ZINCSEARCH_INDEX_NAME=YOUR_ZINCSEARCH_INDEX_NAME \
ALLOWED_ORIGINS=YOUR_ALLOWED_ORIGINS \
go run ./cmd/api
```

## Frontend

A web-based frontend is available for browsing and searching indexed emails. It provides an intuitive user interface for querying and viewing email content seamlessly.

You can find the frontend repository here: [Emails Viewer](https://github.com/mryeibis/emails-viewer). Follow its setup instructions to integrate it with this backend service.

