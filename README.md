# Portfolio Auth-Service

#### The Portfolio Auth-Service is responsible for managing user authentication and authorization. This service provides APIs for user registration, login, and token management, ensuring secure access to the portfolio management system.

## Features

***• User Registration: Create new user accounts.***

***• User Login: Authenticate users and issue tokens.***

***• Token Validation: Validate tokens for secure access to other services.***

***• Password Management: Secure password storage and verification.***

## Installation

### 1. Clone the repository:

```sh
git clone git@github.com:Javokhdev/Portfolio-AuthService.git
cd Portfolio-AuthService
```

### 2. Set up the environment:

***• Go (1.16+)***

***• PostgreSQL***

### 3. Install dependencies:

```
go mod tidy
```

#### 4. Set up the database:

***• Create a PostgreSQL database.***

***• Run the migrations located in the migrations directory.***


### 5. Configuration:

Create a .env file in the root directory with the following environment variables:

```.env
HTTP_PORT=:8081

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=root
POSTGRES_DATABASE=auth_service

DEFAULT_OFFSET=1
DEFAULT_LIMIT=10

TOKEN_KEY=my_secret_key
```
## Usage

#### Run the service:

```sh
go run server/server.go
go run main.go
```

## Contributing

Contributions are always welcome!

See contributing.md for ways to get started.

Please adhere to this project's code of conduct.

## License

This project is licensed under the MIT License.

## Acknowledgement

#### Javoxir Xasanov 
[![telegram](https://img.shields.io/badge/telegram-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://t.me/javohir_khasanov)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/javohir-xasanov/)




