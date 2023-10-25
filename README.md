# Go User RestAPI

A minimalistic REST API microservice built with Go for user authentication.

## Overview

This project is a lightweight microservice focused on user authentication. Built with Go and leveraging its standard library, it serves as a demonstration of creating RESTful APIs in Go with a focus on user registration, login, and token-based authentication.

## Features

- **User Registration**: Allows new users to create an account.
- **User Login**: Existing users can log in and receive a token for authenticated routes.
- **Token-based Authentication**: Utilizes JWT (JSON Web Tokens) for secure and stateless authentication.
- **In-memory Store**: A temporary storage solution to hold user data.
- **Profile Management**: Allows users to view, update, and delete their profiles.

## Getting Started

### Prerequisites

- Go (version 1.xx or newer)

### Environment Variables
Before running the project, you need to set the following environment variables:
- `HOST`: Host address for the server (e.g., `localhost` or `0.0.0.0`). Default: `localhost`.
- `PORT`: Port on which the server will listen (e.g., `8080`). Default: `8080`.
- `JWT_KEY`: Secret key for generating and validating JWT tokens. Ensure it's a strong, unique key. No default.
- `ALLOWED_ORIGINS`: Comma-separated list of allowed origins for CORS. Default: `*` (allow all origins).

### Running the Project

1. Clone the repository:
   ```
   git clone https://github.com/mark-c-hall/GoUserRestAPI.git
   ```

2. Navigate to the project directory:
   ```
   cd GoUserRestAPI
   ```

3. Run the server:
   ```
   go run main.go
   ```

The API will start and listen on the configured port, e.g., `:8080`.

## Endpoints

- `POST /register`: Register a new user.
- `POST /login`: Login and receive a token.
- `POST /logout`: Logout the current user and invalidate the token.
- `GET /profile`: Retrieve the profile information of the authenticated user.
- `POST /profile/update`: Update user profile details.
- `POST /profile/delete`: Delete the user's profile.
- `GET /health`: Health check endpoint returning a 200 OK status, useful for monitoring and service checks.

Note: Ensure that the appropriate HTTP methods (GET, POST, etc.) are used when making requests to these endpoints.

## Future Improvements

- Replace the in-memory store with a persistent database.
- Add user profile and settings management features.
- Implement rate-limiting and additional security measures.

## Contributing

While this is a small project for educational purposes, contributions or feedback are welcome. Feel free to open an issue or submit a pull request.

## License

This project is open-source and available under the [MIT License](LICENSE).

## Acknowledgements

- [Go](https://golang.org/)
- [JWT-Go Library](https://github.com/dgrijalva/jwt-go) for token generation and validation.
