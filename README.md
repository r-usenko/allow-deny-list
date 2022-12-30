```go
package main

import (
	"fmt"

	access "github.com/r-usenko/allow-deny-list"
)

func main() {
	
	if err := access.Allowed("A", "B").IsAllowed("C"); err != nil {
		fmt.Print("Allowed", err)
	}
	
	if err := access.Denied(1, 2).IsAllowed(3); err != nil {
		fmt.Print("Denied", err)
	}
}
```
