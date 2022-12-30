Easy check exists value in allowed and denied lists

```go
package main

import (
	"fmt"

	access "github.com/r-usenko/allow-deny-list"
)

func main() {

	if err := access.Allowed("A", "B").IsAllowed("C"); err != nil {
		fmt.Println("Scenario #1:", err)
	}

	if err := access.Allowed("A", "B", "C").IsAllowed("A"); err != nil {
		fmt.Println("Scenario #2:", err)
	}

	if err := access.Denied(1, 2).IsAllowed(3); err != nil {
		fmt.Println("Scenario #3:", err)
	}

	if err := access.Denied(int64(1), int64(2), int64(3)).IsAllowed(int64(1)); err != nil {
		fmt.Println("Scenario #4:", err)
	}
}

```
