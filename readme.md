# 🚀 Brototype API Wrapper

Welcome to the **Brototype API Wrapper**, a Golang-Fiber-based application designed to provide a structured and easy-to-use interface around the existing Brototype API. This project simplifies accessing data, making it easier for frontend applications or data aggregation tools to interact with Brototype services efficiently.

## 🌟 Overview
This API wrapper offers seamless access to authentication, user details, review schedules, and foundational review data from the Brototype ecosystem. It is fully Dockerized for easy deployment and ensures smooth integration with modern applications.

## 🛠️ Tech Stack

- **Language**: Golang (Fiber framework)
- **Containerization**: Docker
- **Authentication**: Token-based authentication
- **API Testing**: cURL / Postman

## ✨ Features

- **🔐 Authentication** – Secure login with mobile number and password.
- **📄 User Details** – Fetch comprehensive details from Brototype.
- **📝 Review Management** – Retrieve review schedules, status, and assigned advisors.
- **📚 Foundation Reviews** – Access structured data on foundational review stages.
- **📦 Dockerized Deployment** – Easily deploy using Docker.

## 📌 API Documentation

### 🔑 Authentication
- **Endpoint:** `POST http://localhost:8080/api/auth`
- **Request:**
```json
{
  "mobile": "+910000000000",
  "password": "sample123"
}
```
- **Response:**
```json
{
   "token":"your_jwt_token_here"
}
```
- **cURL Command:**
```bash
curl -X POST http://localhost:8080/api/auth \
-H "Content-Type: application/json" \
-d '{
  "mobile": "+910000000000",
  "password": "sample123"
}'
```

### 📋 Fetch User Details
- **Endpoint:** `GET http://localhost:8080/api/details`
- **Headers:** `Authorization: Bearer <TOKEN>`
- **cURL Command:**
```bash
curl -X GET http://localhost:8080/api/details \
-H "Authorization: Bearer <TOKEN>"
```

### 🔍 Retrieve Reviews
- **Endpoint:** `GET http://localhost:8080/api/reviews`
- **cURL Command:**
```bash
curl --location --request GET 'http://localhost:8080/api/reviews' \
--header 'Authorization: Bearer <TOKEN>'
```

### 🏛️ Foundation Reviews
- **Endpoint:** `GET http://localhost:8080/api/foundations?reviewStageCode=value`
- **Query Parameters:**
  - `1` – Upcoming
  - `2` – Conducted
  - `3` – Completed
- **cURL Command:**
```bash
curl --location --request GET 'http://localhost:8080/api/foundations?reviewStageCode=3' \
--header 'Authorization: Bearer <TOKEN>'
```

## 🛠️ Installation & Setup

### 🔹 Prerequisites
- Docker installed on your system
- Golang installed (if running locally)

### 🔹 Run with Docker
```bash
git clone https://github.com/yourusername/brototype-api-wrapper.git
cd brototype-api-wrapper
docker-compose up -d
```

### 🔹 Run Locally
```bash
git clone https://github.com/yourusername/brototype-api-wrapper.git
cd brototype-api-wrapper
go mod tidy
go run main.go
```

## 🎯 Contributors Guide

### 🛠 How to Contribute
1. **Fork the Repository**: Click the fork button on GitHub.
2. **Clone Your Fork**: 
   ```bash
   git clone https://github.com/yourusername/brototype-api-wrapper.git
   ```
3. **Create a Branch**:
   ```bash
   git checkout -b feature-branch
   ```
4. **Make Changes** and **Commit**:
   ```bash
   git commit -m "Added new feature"
   ```
5. **Push to Your Fork**:
   ```bash
   git push origin feature-branch
   ```
6. **Submit a Pull Request** on GitHub.

### 📜 Code Style Guidelines
- Follow Golang best practices.
- Ensure proper error handling.
- Maintain clean and readable code.

## 📜 License
This project is licensed under the MIT License.

---

🚀 **Join the Brototype API Wrapper Community!** Let’s build something amazing together! 💡
