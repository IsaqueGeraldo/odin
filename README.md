# Odin

The `odin` package provides a simple and lightweight environment variable management system using SQLite as a backend. It allows you to store, retrieve, and manipulate environment variables programmatically, with support for sanitization and expansion.

---

## Features

- **Environment Variable Storage**: Store key-value pairs in a SQLite database.
- **Sanitization**: Automatically sanitizes keys to ensure they are valid.
- **CRUD Operations**: Create, retrieve, update, and delete environment variables.
- **Clear All Variables**: Clear all stored environment variables in one command.
- **Environment Expansion**: Expand strings with placeholders for environment variables.
- **SQLite Backend**: Uses SQLite for persistent storage.

---

## Installation

To use the `odin` package, import it into your Go project:

```go
import "github.com/IsaqueGeraldo/odin"
```

Ensure you have the required dependencies installed:

```bash
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

---

## Usage

### 1. Initialize the Database

Before using the package, call the `Bootstrap` function to initialize the SQLite database and create the necessary table:

```go
package main

import "github.com/IsaqueGeraldo/odin"

func main() {
    odin.Bootstrap()
}
```

---

### 2. Set an Environment Variable

Use `Setenv` to store a key-value pair:

```go
err := odin.Setenv("APP_PORT", "8080")
if err != nil {
    panic(err)
}
```

---

### 3. Get an Environment Variable

Retrieve the value of a stored key using `Getenv`:

```go
value := odin.Getenv("APP_PORT")
fmt.Println("APP_PORT:", value)
```

---

### 4. Delete an Environment Variable

Remove a specific key using `Unsetenv`:

```go
err := odin.Unsetenv("APP_PORT")
if err != nil {
    panic(err)
}
```

---

### 5. Clear All Environment Variables

Delete all stored environment variables with `Clearenv`:

```go
odin.Clearenv()
```

---

### 6. List All Environment Variables

Retrieve all stored key-value pairs as a slice of strings:

```go
envs := odin.Environ()
for _, env := range envs {
    fmt.Println(env)
}
```

---

### 7. Expand Environment Variables in Strings

Expand placeholders in a string using `ExpandEnv`:

```go
odin.Setenv("APP_PORT", "8080")
result := odin.ExpandEnv("Server is running on port $APP_PORT")
fmt.Println(result) // Output: Server is running on port 8080
```

---

### 8. Lookup an Environment Variable

Check if a key exists and retrieve its value using `LookupEnv`:

```go
value, exists := odin.LookupEnv("APP_PORT")
if exists {
    fmt.Println("APP_PORT:", value)
} else {
    fmt.Println("APP_PORT not set")
}
```

---

## Key Sanitization

The `odin` package automatically sanitizes keys to ensure they are valid. Keys are converted to uppercase, and invalid characters are replaced with underscores (`_`).

Example:

```go
odin.Setenv("app-port", "8080")
value := odin.Getenv("APP_PORT") // Sanitized key
fmt.Println(value) // Output: 8080
```

---

## Database File

The SQLite database file is named odin.db and is created in the current working directory. Ensure the application has write permissions to this directory.

---

## Dependencies

- [GORM](https://gorm.io/) - The ORM used for database interactions.
- [SQLite Driver](https://gorm.io/docs/connecting_to_the_database.html#SQLite) - SQLite driver for GORM.

---

## License

This package is open-source and available under the MIT License.

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the package.

---

## Author

Developed by [Isaque Geraldo](https://github.com/IsaqueGeraldo).
