# Go Challenge

The challenge has been executed in two different steps.

## Steps

### 1 Hour

`tag: 1 hour`:
As the challenge required, the tag is a snapshot of what has been done in one hour.
To be honest it is actually one hour and half, just to clean a little bit the code
and perform some tests via cURL and fix corresponding small errors.

What has been done:

- Design from scratch of the application
- Definition of Database models
- Setup of the environment including db connections, libraries and frameworks
(in particular [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io))
- Design and coding of the Backend which requires two APIs
  - GET: `/challenges` => retrieve all the available challenges in the db
  - POST: `/challenge/:id/:action` => allow a user to _accept_ a challenge
- Minimal error check and sanitization of data received from clients

For sake of simplicity:

- some sample data are automatically removed and inserted
in the database, when launching the application
- all the authorization middleware for APIs has been skipped

#### curl

- GET: `curl -s http://localhost:8080/challenges`
- POST: `curl -vv -X POST http://localhost:8080/challenge/64/accept`

## Extra tasks

For personal interest I have gone beyond and spent a couple more hours on the project.

- Written Test cases and removed initialization data from db package
- Added Dockerfile
- Added Github Actions (and corresponding code linting)
- Proposed an advanced Model for Challenge to express it via some specific constraints

## Future improvements

- Documentation & Testing
- Improved models according to more requirements
- Fine grained error handling and data sanitization
- Improved design to support lifecycle of Challenges

## Usage

### Prerequisite: DB Connection string

In order to avoid committing secrets on a public repository,
the connections string is supposed to be provided via env variable `DB_STR`

It is possible to create a text file named `.env` (dotenv) in the root directory,
which contains a line in the form:

    DB_STR=<postgres://...>

### Running

    go run main.go

### Testing

    go test -v

### Running using Docker

    docker build -t goes .
    docker run -p 8080:8080 goes
