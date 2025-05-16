# Go Auth System 🔐

A lightweight, scalable authentication system built in Go — perfect for MVPs, personal projects, or as a starter backend for production-ready applications.

## 🚀 Features

- User registration & login
- Password hashing with bcrypt
- JWT-based authentication
- Middleware to protect private routes
- Environment variable support
- Clean project structure
- Extensible for future features

  
## 🧪 API Endpoints

### ✅ Signup
`POST /auth/signup`

**Request Body:**
```json
{
  "username":"yourusername"
  "email": "user@example.com",
  "password": "yourpassword"
}

[Screenshot 2025-05-16 085109](https://github.com/user-attachments/assets/a1cd533c-7cde-4312-9b66-a455599aa818)

### ✅ Login
`POST /auth/login`

** Request Body **
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}

**Response**
{
  "token": "JWT_TOKEN_HERE",
  "user": "user.Email"
}
![Screenshot 2025-05-16 085137](https://github.com/user-attachments/assets/e7d78e5d-c80f-4f40-9bfd-ca095768135c)

🛠 Tech Stack
-Go (Golang)
-Gin
-Air(hot reload)
-bcrypt
-JWT
-PostgreSQL
