# Gocommerce – Simple E-Commerce API (Go + Fiber)

Gocommerce is a clean, modular, and production-ready e-commerce REST API built using Go, Fiber, GORM, and MySQL. Designed for educational and internship purposes, this project adheres to the Rakamin Evermos Virtual Internship specifications.

## ✨ Features

- JWT-based authentication (Login & Register)
- Product, Category, and Transaction management
- Store auto-created on user registration
- File upload (product photos & store logos)
- Address book with filtering
- Pagination & filtering for listing endpoints
- Ownership and role-based access control (admin only category access)
- Product logs saved during transactions

## 🧱 Tech Stack

- **Go** with **Fiber** framework
- **GORM** ORM with **MySQL**
- **Redis** for caching JWT
- **Clean Architecture** with flat-layered structure
- File uploads handled using `multipart/form-data`


## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- MySQL
- Redis

### Setup

```bash
git clone https://github.com/nurhidaylma/gocommerce.git
cd gocommerce
cp .env.example .env
go mod tidy
go run main.go
