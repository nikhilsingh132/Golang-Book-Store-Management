# 📚 Book Management System API

A robust RESTful API service built with Go for managing a book store's inventory. This system provides comprehensive CRUD (Create, Read, Update, Delete) operations for book management with a clean architecture and MySQL database integration.

## 🌟 Features

- **Complete CRUD Operations**: Full support for Create, Read, Update, and Delete operations on books
- **RESTful API Design**: Clean and intuitive API endpoints following REST principles
- **MySQL Database Integration**: Persistent data storage with GORM ORM
- **MVC Architecture**: Well-organized code structure following Model-View-Controller pattern
- **Auto-Migration**: Automatic database schema creation and updates
- **JSON Response Format**: Consistent JSON responses for all API endpoints
- **Error Handling**: Comprehensive error handling with appropriate HTTP status codes
- **Modular Design**: Separated concerns with distinct packages for models, controllers, routes, and utilities

## 🛠 Tech Stack

- **Language**: Go 1.19+
- **Web Framework**: Gorilla Mux (HTTP router and URL matcher)
- **ORM**: GORM (Go Object Relational Mapper)
- **Database**: MySQL
- **JSON Processing**: Go standard library encoding/json

## 📋 Prerequisites

Before running this application, ensure you have the following installed:

- **Go**: Version 1.19 or higher ([Download Go](https://golang.org/dl/))
- **MySQL**: Version 5.7 or higher ([Download MySQL](https://dev.mysql.com/downloads/))
- **Git**: For cloning the repository ([Download Git](https://git-scm.com/downloads))

## 🚀 Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/Book_Management_System_Golang.git
cd Book_Management_System_Golang
```

### 2. Install Dependencies

```bash
go mod download
```

If `go.mod` doesn't exist, initialize it:

```bash
go mod init book_store
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
```

### 3. Database Configuration

#### Create MySQL Database

```sql
CREATE DATABASE golang_db;
USE golang_db;
```

#### Update Database Connection

Edit the database configuration in `pkg/config/app.go`:

```go
// Update with your MySQL credentials
d, err := gorm.Open("mysql", "username:password@/golang_db?charset=utf8&parseTime=True&loc=Local")
```

### 4. Run the Application

```bash
go run cmd/main/main.go
```

The server will start on `http://localhost:8001`

## 📂 Project Structure

```
Book_Management_System_Golang/
│
├── cmd/
│   └── main/
│       └── main.go              # Application entry point
│
├── pkg/
│   ├── config/
│   │   └── app.go               # Database configuration
│   │
│   ├── controllers/
│   │   └── book-controller.go   # Business logic for book operations
│   │
│   ├── models/
│   │   └── book.go              # Book model and database operations
│   │
│   ├── routes/
│   │   └── book-routes.go       # API route definitions
│   │
│   └── utils/
│       └── book-utils.go        # Utility functions
│
├── go.mod                       # Go module dependencies
├── go.sum                       # Dependency checksums
└── README.md                    # Project documentation
```

## 🔗 API Endpoints

### Base URL
```
http://localhost:8001
```

### Endpoints Overview

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/book/` | Get all books |
| GET | `/book/{bookId}` | Get a specific book by ID |
| POST | `/book/` | Create a new book |
| PUT | `/book/{bookId}` | Update a book |
| DELETE | `/book/{bookId}` | Delete a book |

## 📝 API Documentation

### 1. Get All Books
**Endpoint**: `GET /book/`

### 2. Get Book by ID
**Endpoint**: `GET /book/{bookId}`

### 3. Create a New Book
**Endpoint**: `POST /book/`

### 4. Update a Book
**Endpoint**: `PUT /book/{bookId}`

### 5. Delete a Book
**Endpoint**: `DELETE /book/{bookId}`

## 🔍 Database Schema

The application automatically creates and manages the following table structure:

### Books Table
```sql
CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    name VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publication VARCHAR(255),
    INDEX idx_deleted_at (deleted_at)
);
```

---
**Happy Coding! 🚀**
