# LazyIO Package

A simple I/O package for Go, providing functions for user input, output, and file reading.

## Functions
### Input

```go
func Input(s ...string) string
```

The Input function allows the user to input text from the standard input (keyboard).

Parameters
- `s ...string`: Variadic parameter to display prompt messages to the user.
Returns
- `string`: The text entered by the user.

Example

```go
package main

import (
	"fmt"
	"github.com/Fastiraz/lazyio"
)

func main() {
	name := lazyio.Input("Enter your name: ")
	fmt.Printf("Hello, %s!\n", name)
}
```

### Output

```go
func Output(s ...string)
```

The Output function prints text to the standard output (console).

Parameters
`s ...string`: Variadic parameter for the text to be printed.

Example

```go
package main

import (
	"github.com/Fastiraz/lazyio"
)

func main() {
	lazyio.Output("This is an example.", " Hello, World!")
}
```

### ReadFile

```go
func ReadFile(filename string) ([]string, error)
```

The ReadFile function reads the contents of a file and returns them as a slice of strings.

Parameters
`filename string`: The name of the file to be read.
Returns
`[]string`: A slice containing the lines read from the file.
`error`: An error if any occurred during file reading.

Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/Fastiraz/lazyio"
)

func main() {
	lines, err := lazyio.ReadFile("file.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	for _, line := range lines {
		lazyio.Output(line)
	}
}
```

Feel free to use and modify the LazyIO package to simplify your I/O operations in Go.
