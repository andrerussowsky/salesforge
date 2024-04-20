# SalesForge API

SalesForge is an API for managing sequences and email templates.

## Prerequisites

- Go 1.15 or higher installed on your machine.
- Docker installed on your machine.
- PostgreSQL database running locally or accessible via a connection string.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/andrerussowsky/salesforge.git
   ```

2. Change directory to the project folder:

   ```bash
   cd salesforge
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Install Goose for managing database migrations:

   ```bash
   go get -u github.com/pressly/goose/v3/cmd/goose
   ```

## Database Setup

### Docker Setup

1. Run a PostgreSQL container:

   ```bash
   docker run --name salesforge-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
   ```

2. Create the database:

   ```bash
   docker exec -it salesforge-postgres psql -U postgres -c "CREATE DATABASE salesforge"
   ```

### Database Migrations

1. Apply migrations to the database using Goose:

   ```bash
   goose -dir migrations postgres "host=localhost user=postgres password=mysecretpassword dbname=salesforge sslmode=disable" up
   ```

   This will execute the following SQL statements:

   ```sql
   CREATE TABLE sequences (
       id SERIAL PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       open_tracking_enabled BOOLEAN,
       click_tracking_enabled BOOLEAN
   );

   CREATE TABLE steps (
       id SERIAL PRIMARY KEY,
       sequence_id INT,
       email_subject VARCHAR(255) NOT NULL,
       email_content TEXT,
       FOREIGN KEY (sequence_id) REFERENCES sequences(id) ON DELETE CASCADE
   );
   ```

## Running the Application

1. Build and run the application:

   ```bash
   go run cmd/main.go
   ```

## Curl Examples

### Create a Sequence with Steps:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "New Sequence",
    "open_tracking_enabled": true,
    "click_tracking_enabled": true,
    "steps": [
        {
            "email_subject": "Subject of Step 1",
            "email_content": "Content of Step 1"
        },
        {
            "email_subject": "Subject of Step 2",
            "email_content": "Content of Step 2"
        }
    ]
}' http://localhost:8080/sequence
```

### Update a Sequence Step:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
    "email_subject": "Updated Subject of Step 1",
    "email_content": "Updated Content of Step 1"
}' http://localhost:8080/sequence/1/step/1
```

### Delete a Sequence Step:

```bash
curl -X DELETE http://localhost:8080/sequence/1/step/1
```

### Update Sequence Tracking:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
    "open_tracking": true,
    "click_tracking": false
}' http://localhost:8080/sequence/1/tracking
```
