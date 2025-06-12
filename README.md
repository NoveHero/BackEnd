# Go-based Authentication and Story Management Backend

This repository contains the backend component of a larger application, implemented in **Golang**, with a focus on user authentication, authorization using JWT, and story management endpoints. The project follows a modern microservice architecture and is deployed using Docker.

---

## 🚀 Technology Stack

### 🔹 Language: Go (Golang)

The backend is developed using **Go**, known for its simplicity, performance, and support for concurrency. Go is a great choice for building scalable backend services with high throughput.

### 🔹 Web Framework: [Fiber](https://github.com/gofiber/fiber)

Fiber is a fast, Express.js-inspired web framework built on top of Go. Its syntax and simplicity make it ideal for writing clean and performant HTTP services.

### 🔹 ORM: [GORM](https://gorm.io)

GORM is used for interacting with the relational database. It provides a robust ORM layer to simplify database operations such as user management and story storage.

---

## 🔐 JWT-based Authentication

The project uses **JWT (JSON Web Tokens)** for stateless authentication. Benefits of this approach include:

- **Statelessness**: No need to store sessions on the server.
- **Secure**: Tokens are signed using HMAC or RSA algorithms.
- **Portable**: Tokens can be passed in HTTP headers, cookies, or URLs.
- **Flexible**: Payload can contain custom claims such as roles or expiration time.
- **Widely Supported**: Easily integrated across platforms.
- **Access Control**: Role-based authorization is supported.
- **Token Expiry**: Tokens include expiration timestamps for improved security.

### 🧩 Middleware for Token Validation

A middleware handles token verification:
- Extracts JWT from `Authorization` header.
- Strips the `Bearer` prefix.
- Parses and validates the token using a secret key.
- If valid, stores the token claims in Fiber’s context for downstream access.

---

## 📌 API Endpoints

### Auth Endpoints
- `POST /api/auth/signup`: Register a new user.
- `POST /api/auth/login`: User login and token generation.
- `POST /api/auth/change-password`: Change user password.

### User Profile
- `GET /api/me/`: Retrieve current user profile info.
- `PATCH /api/me/change-nickname`: Change user nickname.

### Story Management
- `GET /api/me/stories`: Get list of user stories.
- `GET /api/me/stories/:uuid`: Get story details by UUID.
- `POST /api/me/stories`: Create a new story.
- `PUT /api/me/stories/:uuid/update`: Update story content by UUID.
- `PATCH /api/me/stories/:uuid/change-title`: Update story title.
- `DELETE /api/me/stories/:uuid/`: Delete a story by UUID.
- `PATCH /api/me/stories/meta/:uuid/`: Update story metadata.

### LLM Permission
- `GET /api/llm-permission`: Request permission to access the LLM service.

---

## 🐳 Deployment with Docker

The backend is containerized using Docker for portability and ease of deployment.

### 🧱 Dockerfile Highlights

- **Base Image**: `golang:1.23.2-alpine` (lightweight & fast).
- **Working Directory**: `/app`.
- **Dependency Management**: Copies `go.mod` and `go.sum`, installs dependencies.
- **Build**: Compiles the project to a binary named `main`.
- **Final Image**: Uses `alpine:latest` to keep the container lightweight.
- **Expose Port**: 3000 is exposed for external access.
- **Run**: Application is started via `CMD ["./main"]`.

---

## 📦 Getting Started

```bash
# Clone the repository
git clone https://github.com/NoveHero/BackEnd.git
cd your-repo-name

# Build the Docker image
docker build -t go-auth-backend .

# Run the container
docker run -p 3000:3000 go-auth-backend
