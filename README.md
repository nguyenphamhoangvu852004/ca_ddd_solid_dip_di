# Clean Architecture & SOLID Principles in Go

This project demonstrates the application of **Clean Architecture**, **Domain-Driven Design (DDD)**, and **SOLID Principles** in a Go application using Gin web framework.

## Table of Contents

- [Project Overview](#project-overview)
- [Clean Architecture Layers](#clean-architecture-layers)
- [SOLID Principles Applied](#solid-principles-applied)
- [Dependency Injection Pattern](#dependency-injection-pattern)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [How to Build & Run](#how-to-build--run)

---

## Project Overview

This application follows the **Clean Architecture** pattern where the codebase is organized into distinct layers, each with a specific responsibility. The User module serves as a complete example of how these principles are implemented.

### Key Goals:
✅ **Independence of Frameworks** - Business logic is not dependent on web frameworks  
✅ **Testability** - Each layer can be tested independently  
✅ **Maintainability** - Clear separation of concerns makes code easy to understand and modify  
✅ **Flexibility** - Easy to swap implementations without affecting other layers  

---

## Clean Architecture Layers

### 1️⃣ **Presentation Layer (Controller)**
**Location:** `internal/controller/`

Responsible for handling HTTP requests and returning responses.

```
├── user_controller.go
│   ├── GetUsers() - HTTP GET handler
│   └── CreateUser() - HTTP POST handler
```

**Characteristics:**
- Handles HTTP protocol concerns (status codes, request/response)
- Converts HTTP requests to use case input
- Converts use case output to HTTP responses
- **Does NOT** contain business logic
- **Depends on** Presentation (UseCase layer)

**Example:**
```go
func (uc *UserController) CreateUser(c *gin.Context) {
    // Call use case (business logic)
    uc.createUserUseCase.Execute()
    
    // Return HTTP response
    c.JSON(http.StatusCreated, gin.H{
        "message": "User created successfully",
    })
}
```

---

### 2️⃣ **Application/Use Case Layer**
**Location:** `internal/usecase/`

Orchestrates the application workflows and implements business logic.

```
├── create_user.go
│   ├── CreateUser struct (use case)
│   └── Execute() - business operation
```

**Characteristics:**
- Contains application-specific business rules
- Orchestrates use of domain entities and repositories
- **Does NOT** know about HTTP or database details
- **Depends on** Domain layer (interfaces)
- **Depends on** infrastructure only through interfaces

**Example:**
```go
type CreateUser struct {
    userRepo domain_repository.IUserRepository // Interface dependency
}

func (c *CreateUser) Execute() {
    c.userRepo.Create() // Call repository through interface
}
```

---

### 3️⃣ **Domain Layer**
**Location:** `internal/domain/`

Contains core business logic and domain models.

```
├── entities/
│   └── user.entity.go (User business entity)
└── repository/
    └── user.repository.go (Repository interface)
```

**Characteristics:**
- **Core business rules** - User entity with `IsTeenager()` method
- **Domain interfaces** - `IUserRepository` defines contracts
- **No external dependencies** - Pure business logic
- **Framework agnostic** - Can be used in any context

**Example:**
```go
type User struct {
    ID   string
    Name string
    Age  int
}

// Pure business logic
func (u *User) IsTeenager() bool {
    return u.Age >= 16 && u.Age < 20
}
```

---

### 4️⃣ **Infrastructure Layer**
**Location:** `internal/infrastructure/`

Implements technical concerns like database access, external APIs, etc.

```
├── repository/
│   └── user_repository.go (Repository implementation)
```

**Characteristics:**
- Implements domain interfaces (`IUserRepository`)
- Handles database operations, external APIs, etc.
- **Encapsulates external dependencies** - Database drivers, ORM, HTTP clients
- **Replaceable** - Can swap implementations without affecting layers above
- **Depends on** Domain layer (implements interfaces)

**Example:**
```go
type UserRepositoryImpl struct {
    // MySQL, PostgreSQL, etc.
}

func (u *UserRepositoryImpl) Create() {
    // Actual database implementation
    panic("unimplemented")
}
```

---

### 5️⃣ **Dependency Injection Layer**
**Location:** `initialize/dependency_injection/`

Wires up all dependencies - the composition root.

```
├── user.go
│   └── InitUserDependencies() - Dependency container
```

**Characteristics:**
- **Single point of assembly** - All dependencies initialized here
- **Follows Dependency Inversion** - Object creation isolated from business logic
- **Easy to test** - Mock dependencies by creating different DI functions
- **Main benefit** - Changes to dependencies don't cascade through codebase

**Example:**
```go
func InitUserDependencies() *UserDI {
    // Layer 4: Infrastructure
    userRepository := infrastructure_repository.NewUserRepostoryImpl()
    
    // Layer 2: Use Case (depends on repository interface)
    createUserUseCase := usecase.NewCreateUser(userRepository)
    
    // Layer 1: Presentation (depends on use case)
    userController := controller.NewUserController(createUserUseCase)
    
    return &UserDI{UserController: userController}
}
```

---

## SOLID Principles Applied

### **S** - Single Responsibility Principle
Each class/struct has ONE reason to change:

| Component | Responsibility |
|-----------|---|
| `UserController` | Handle HTTP requests/responses |
| `CreateUser` | Orchestrate user creation workflow |
| `User` (Entity) | Business rules related to users |
| `UserRepositoryImpl` | Database operations |

```go
// ✅ Good - Single responsibility
type CreateUser struct {
    userRepo IUserRepository
}

// ❌ Bad - Mixed responsibilities
type UserHandler struct {
    userRepo IUserRepository
    httpClient *http.Client
    logger Logger
}
```

---

### **O** - Open/Closed Principle
Classes are **OPEN** for extension, **CLOSED** for modification:

- **Closed:** Don't modify existing layers
- **Open:** Extend by creating new implementations

```go
// Existing interface - CLOSED for modification
type IUserRepository interface {
    Get()
    Create()
}

// ✅ Easy to extend - Just implement the interface
type PostgresUserRepository struct { }
func (p *PostgresUserRepository) Get() { /* ... */ }

// ✅ No changes needed to use case or controller
```

---

### **L** - Liskov Substitution Principle
Subclasses must be substitutable for their parent classes:

```go
// Any implementation of IUserRepository can replace another
var repo IUserRepository

// Use case works with ANY implementation
createUserUseCase := NewCreateUser(repo)
```

---

### **I** - Interface Segregation Principle
Clients should NOT depend on interfaces they don't use:

```go
// ✅ Good - Specific interface
type IUserRepository interface {
    Get()
    Create()
}

// Use case only needs what it uses
type CreateUser struct {
    userRepo IUserRepository
}

// ❌ Bad - Fat interface
type IRepository interface {
    Get()
    Create()
    Update()
    Delete()
    GetAll()
    Search()
    Paginate()
}
```

---

### **D** - Dependency Inversion Principle
High-level modules should NOT depend on low-level modules. Both should depend on abstractions:

```
Without Dependency Inversion (❌ Bad):
Controller → UseCase → RepositoryImpl (concrete database code)

With Dependency Inversion (✅ Good):
Controller → UseCase → IUserRepository (interface)
                        ↓ (implements)
                    RepositoryImpl (concrete database code)
```

**Implementation:**
```go
// UseCase depends on INTERFACE (abstraction)
type CreateUser struct {
    userRepo IUserRepository // Interface, not concrete implementation
}

// DI layer injects concrete implementation
repo := NewUserRepostoryImpl()
useCase := NewCreateUser(repo) // Passing interface-compatible implementation
```

---

## Dependency Injection Pattern

### The Dependency Chain (Bottom-Up Construction)

```
Step 1: Create Infrastructure (Lowest Level)
↓
userRepository := NewUserRepostoryImpl()

Step 2: Create Use Case (Depends on Repository Interface)
↓
createUserUseCase := NewCreateUser(userRepository)

Step 3: Create Controller (Depends on Use Case)
↓
userController := NewUserController(createUserUseCase)

Result: Complete dependency chain ready to use
```

### Benefits of DI

| Benefit | Explanation |
|---------|---|
| **Loose Coupling** | Components don't know how to create their dependencies |
| **Testability** | Easy to inject mock objects for testing |
| **Flexibility** | Swap implementations without code changes |
| **Maintainability** | Dependency flow is clear and explicit |

---

## API Endpoints

### User Module

| Method | Endpoint | Handler | Layer |
|--------|----------|---------|-------|
| `GET` | `/api/v1/users` | `UserController.GetUsers()` | Presentation |
| `POST` | `/api/v1/users` | `UserController.CreateUser()` | Presentation |

**Request Flow:**
```
HTTP Request
    ↓
Controller (Layer 1: Presentation)
    ↓
UseCase (Layer 2: Application)
    ↓
Repository Interface (Layer 3: Domain)
    ↓
RepositoryImpl (Layer 4: Infrastructure)
    ↓
Database
```

---

## Project Structure

```
ca_ddd_solid_dip_di/
├── cmd/
│   ├── server/
│   │   └── main.go                 # Entry point
│   └── test/
├── internal/                        # Private application code
│   ├── controller/
│   │   └── user_controller.go       # Layer 1: Presentation
│   ├── domain/
│   │   ├── entities/
│   │   │   └── user.entity.go       # Layer 3: Domain Models
│   │   └── repository/
│   │       └── user.repository.go   # Layer 3: Domain Interfaces
│   ├── infrastructure/
│   │   └── repository/
│   │       └── user_repository.go   # Layer 4: Infrastructure
│   └── usecase/
│       └── create_user.go           # Layer 2: Application Logic
├── initialize/
│   ├── application.go               # App initialization
│   ├── router.go                    # Route registration
│   └── dependency_injection/
│       └── user.go                  # Layer 5: DI Container
├── go.mod
└── README.md
```

### Layer Assignment Summary

| Directory | Layer | Purpose |
|-----------|-------|---------|
| `cmd/` | **Entry Point** | Application startup |
| `internal/controller/` | **1 - Presentation** | HTTP handlers |
| `internal/usecase/` | **2 - Application** | Business workflows |
| `internal/domain/` | **3 - Domain** | Core business logic & interfaces |
| `internal/infrastructure/` | **4 - Infrastructure** | External dependencies (DB, APIs) |
| `initialize/` | **5 - Composition** | Dependency assembly |

---

## How to Build & Run

### Prerequisites
- Go 1.25.1 or higher
- Gin framework

### Install Dependencies
```bash
go mod download
```

### Run the Application
```bash
go run ./cmd/server/main.go
```

The server will start on `http://localhost:8080`

### Test the API
```bash
# Health check
curl http://localhost:8080/ping

# Get users
curl http://localhost:8080/api/v1/users

# Create user
curl -X POST http://localhost:8080/api/v1/users
```

---

## Key Takeaways

### Why Clean Architecture?

1. **Scalability** - Easy to add new features without affecting existing code
2. **Maintainability** - Clear structure makes the codebase easy to understand
3. **Testability** - Each layer can be tested independently
4. **Flexibility** - Swap implementations (e.g., MySQL → PostgreSQL) without logic changes
5. **Professional Code** - Follows industry best practices

### The Golden Rule

> **Business logic should be separated from infrastructure concerns.**

- ✅ Business logic lives in: Entities, UseCases, Domain interfaces
- ✅ Infrastructure concerns live in: Controllers, Repositories, Database code
- ✅ Connection managed by: Dependency Injection

### Dependency Direction

```
                 ↓ Depends on (points to)
Presentation → Application → Domain ← Infrastructure
(Controller)   (UseCase)    (Entity)  (Repository)
```

All dependencies point **INWARD** toward the Domain layer - this is the essence of **Dependency Inversion Principle**.

---

## References

- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Domain-Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design)
- [SOLID Principles](https://en.wikipedia.org/wiki/SOLID)
- [Go Best Practices](https://golang.org/doc/effective_go)
