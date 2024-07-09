# gRPC User Service

This repository contains a Golang gRPC service for managing user details with search functionality. The service includes endpoints for fetching user details by ID, listing users by IDs, and searching users based on specific criteria.

## Table of Contents

- [Installation](#installation)
- [Build and Run](#build-and-run)
- [Usage](#usage)
  - [GetUser](#getuser)
  - [ListUsers](#listusers)
  - [SearchUsers](#searchusers)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/grpc-user-service.git
    cd grpc-user-service
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

## Build and Run

1. Compile the protobuf file:
    ```bash
    protoc --go_out=. --go-grpc_out=. proto/user.proto
    ```

2. Build the server:
    ```bash
    go build -o server ./server
    ```

3. Run the server:
    ```bash
    ./server
    ```

4. (Optional) Build and run the client:
    ```bash
    go build -o client ./client
    ./client
    ```

## Usage

### GetUser

Fetch user details by user ID.

**Request:**
```json
{
    "id": 1
}

**Response:**
```json
{
    "user": {
        "id": 1,
        "fname": "Steve",
        "city": "LA",
        "phone": 1234567890,
        "height": 5.8,
        "married": true
    }
}
 
### ListUser
Retrieve a list of users by their IDs.

**Request:**
```json
{
   "ids": [1, 2, 3, 4]
}


### SearchUsers
Search for users based on specific criteria (e.g., city, phone number, marital status).

Request:

```json
{
    "city": "CA",
    "phone": 8978678978,
    "married": false
}