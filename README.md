# Odin

Odin is a lightweight Go package for managing environment variables within your Go applications. It provides functions to set, get, list, remove, clear, and rename environment variables stored in a SQLite database. Odin is designed to simplify the management of environment variables in your Go applications.

## Features

- Set an environment variable with a key and value.
- Get the value of an environment variable by specifying its key.
- List all environment variables.
- Remove an environment variable by specifying its key.
- Clear all environment variables.
- Rename an environment variable.

## Installation

To use Odin in your Go application, you can import it as follows:

```go
import (
    "github.com/IsaqueGeraldo/odin"
)
```

Ensure that you have the package installed in your Go environment. You can install it using the following command:

```bash
go get github.com/IsaqueGeraldo/odin
```

## Usage

Here's how you can use Odin in your Go application:

```go
package main

import (
    "fmt"
    "github.com/IsaqueGeraldo/odin"
)

func main() {
    // Initialize the Odin package
    odin.Bootstrap()

    // Set an environment variable
    err := odin.Setenv("API_KEY", "my-secret-key")
    if err != nil {
        fmt.Println("Error setting environment variable:", err)
    }

    // Get the value of an environment variable
    value, err := odin.Getenv("API_KEY")
    if err != nil {
        fmt.Println("Error getting environment variable:", err)
    } else {
        fmt.Println("Value of environment variable 'API_KEY':", value)
    }

    // List all environment variables
    envs, err := odin.Environ()
    if err != nil {
        fmt.Println("Error listing environment variables:", err)
    } else {
        fmt.Println("Environment variables:")
        for _, env := range envs {
            fmt.Printf("%s=%s\n", env.Key, env.Value)
        }
    }

    // Remove an environment variable
    err = odin.Unsetenv("API_KEY")
    if err != nil {
        fmt.Println("Error removing environment variable:", err)
    }

    // Clear all environment variables
    err = odin.Clearenv()
    if err != nil {
        fmt.Println("Error clearing environment variables:", err)
    }

    // Rename an environment variable
    err = odin.RenameKey("oldKey", "newKey")
    if err != nil {
        fmt.Println("Error renaming environment variable:", err)
    }
}
```

This example demonstrates how to use Odin's functions to manage environment variables within your Go application.

## License

Odin is released under the MIT License. See [LICENSE](LICENSE) for more information.

---

Feel free to contribute to this project or report any issues on [GitHub](https://github.com/IsaqueGeraldo/odin).
