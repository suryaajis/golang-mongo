```markdown
# golang-mongo Development Patterns

> Auto-generated skill from repository analysis

## Overview
This skill covers the core development patterns, coding conventions, and workflows observed in the `golang-mongo` repository. The codebase is written in Go and is structured to interact with MongoDB, following idiomatic Go practices with a focus on clarity and maintainability. This guide will help you quickly understand how to contribute code, structure files, and run tests in this repository.

## Coding Conventions

### File Naming
- Use **camelCase** for file names.
  - Example: `userRepository.go`, `mongoClient.go`

### Import Style
- Use **relative imports** within the module.
  - Example:
    ```go
    import (
        "context"
        "myproject/database"
        "myproject/models"
    )
    ```

### Export Style
- Use **named exports** for functions and types that should be accessible outside the package.
  - Example:
    ```go
    // Exported function
    func ConnectToMongo() (*mongo.Client, error) {
        // ...
    }
    ```

### Commit Message Patterns
- Commit messages are **freeform** and concise, averaging 26 characters.
  - Example: `add user repository logic`

## Workflows

### Add a New Repository Layer
**Trigger:** When you need to add a new data access layer for a MongoDB collection.  
**Command:** `/add-repository-layer`

1. Create a new file using camelCase, e.g., `orderRepository.go`.
2. Define a struct for the repository, e.g., `type OrderRepository struct { ... }`.
3. Implement exported methods for CRUD operations.
4. Use relative imports to include necessary models and database utilities.
5. Write corresponding test files following the test file pattern.

### Update MongoDB Connection Logic
**Trigger:** When you need to change how the application connects to MongoDB.  
**Command:** `/update-mongo-connection`

1. Locate the connection logic, typically in `mongoClient.go` or similar.
2. Modify the exported function (e.g., `ConnectToMongo`) as needed.
3. Update any configuration or environment variable usage.
4. Test the connection by running the application or relevant tests.

### Run Tests
**Trigger:** When you want to verify code changes or ensure stability.  
**Command:** `/run-tests`

1. Identify test files matching the `*.test.*` pattern.
2. Use Go’s testing tool to run tests:
    ```sh
    go test ./...
    ```
3. Review the output for any failures or errors.

## Testing Patterns

- **Test Framework:** Not explicitly specified, but Go’s built-in testing is implied.
- **Test File Pattern:** Files are named with the pattern `*.test.*`.
  - Example: `userRepository.test.go`
- **Test Structure:** Use Go’s `testing` package.
  - Example:
    ```go
    import "testing"

    func TestInsertUser(t *testing.T) {
        // test logic here
    }
    ```

## Commands
| Command                  | Purpose                                             |
|--------------------------|-----------------------------------------------------|
| /add-repository-layer    | Scaffold a new repository layer for a collection    |
| /update-mongo-connection | Update MongoDB connection logic                     |
| /run-tests               | Execute all tests in the codebase                   |
```
