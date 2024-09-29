# Event Master

**Event Master** is a backend service written in Go using the Gin web framework. It allows users to send various types of events which are stored, processed, and used to generate reports. The project supports JWT-based authentication for secure access. Event Master is designed to handle events like work sessions and tasks, applying logical corrections to ensure consistency, and storing paired events in Redis for events that require a start and end (e.g., work sessions).

## Features

- **JWT Authentication**: Secure access to the API using JSON Web Tokens (JWT).
- **Event Handling**: Accepts and stores various events, such as session starts, task starts, and more.
- **Redis Integration**: Paired events (e.g., start and end of a work session) are stored in Redis for efficient access.
- **Logical Data Correction**: Automatically corrects out-of-sequence events based on logical rules. For example, a 'task started' event will be adjusted if a 'session started' event was not previously received.
- **Support Of Paired Events**: Events like session_started and session_ended are considered paired. These pairs are stored in Redis until both events are received, at which point they are processed and stored in the main database.
- **Auto Garbage Collection**: If the client fails to send an ending pair for an event, the system recognizes the timeout and record an auto generated ending pair based on client's prefrences.
- **Set Algebra-Based Reports**: By applying set algebra operations (Union, Intersection, Complement), analysts can filter, aggregate, or isolate specific patterns or trends in timestamped event data, and generate more insightful reports.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Event Logic](#event-logic)
- [Report Generation](#report-generation)
- [Technologies Used](#technologies-used)
- [License](#license)

## Installation

1. Clone the repository:
    ```bash
    git clone git@github.com:cotopia-org/Event-Master.git
    ```

2. Navigate to the project directory:
    ```bash
    cd event-master
    ```

3. Install the necessary dependencies:
    ```bash
    go mod tidy
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

## Usage

To start using the API, you'll need to authenticate using a JWT. After authenticating, you can send various events to the server, which will store them, correct any inconsistencies, and generate reports.

You can use tools like `curl` or Postman to interact with the API.


## API Endpoints

### Authentication

- `POST /auth/login`: Authenticate and receive a JWT token.
  
### Events

- `POST /api/events`: Submit a new event. Event model includes:
  - `is_pair` bool
  - `setype` string
  - `pair_type` string
  - `meta` json

### Reports

- endpoint to get events between two points (Line Segment)
- endpoint to get union of line segments (Logical Disjunction)
- endpoint to get intersection of line segments (Logical Conjunction)
- endpoint to get complement of a line segment (Logical Negation) 

## Authentication

Event Master uses JWT tokens for authentication. To interact with the API, you must first obtain a token by sending your login credentials to the `/auth/login` endpoint.

Example request for login:
```bash
curl -X POST -d '{"username": "yourusername", "password": "yourpassword"}' http://localhost:8080/auth/login
```
## Event Logic
Event Master enforces logical rules on events to ensure data consistency:

- **Logical Sequence**: If a task_started event is received but no session_started event has been logged before it, Event Master will automatically insert the missing session_started event.
- **Paired Events**: Events like session_started and session_ended are considered paired. These pairs are stored in Redis until both events are received, at which point they are processed and stored in the database.

## Report Generation
Event Master generates reports using concepts from Algebra of Sets, allowing for complex data analysis. Reports can aggregate, filter, and compare events based on set operations such as union, intersection, and difference.

Reports can be customized based on the time period, event types, and more.

## Technologies Used
- Go: The main programming language for the backend.
- Gin: A high-performance HTTP web framework for Go.
- Redis: Used for storing paired events that are pending (e.g., session start/end).
- JWT: JSON Web Tokens for secure authentication.

## License
TBA

