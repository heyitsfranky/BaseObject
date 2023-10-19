# MyBase

The **MyBase** package serves as a foundation for common utility functions for any type of struct that can be used across multiple projects (e.g. services). I find these incredibly useful, and hopefully you do as well!

Feel free to incorporate this package into your projects and adapt it to your specific needs. It's designed to simplify your development tasks and improve code quality by providing a solid base for common operations.

## Installation

To use **MyBase** in your Go project, you can simply run:

```bash
go get github.com/heyitsfranky/MyBase@latest
```

## Usage

Here's a basic example of how to use some of the functionalities of **MyBase**:
```go
package main

import (
	"fmt"
	"github.com/heyitsfranky/MyBase"
)

type SampleStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Using a sample struct
	myStruct := SampleStruct{Name: "John Doe", Age: 30, Email: "john@example.com"}

	// Convert myStruct to a map
	myMap, err := MyBase.ObjectToMap(&myStruct)
	if err != nil {
		fmt.Println(err)
		return
	}

    // Convert a map to a SampleStruct
    data := map[string]interface{}{"name": "John the Second", "age": 29, "email": "john2@example.com"}
	result, err := MyBase.MapToObject[SampleStruct](data)
	if err != nil {
		fmt.Println(err)
		return
	}

    // And many more...
}

```

## License

This package is distributed under the MIT License.
Feel free to contribute or report issues on GitHub.

Happy coding!