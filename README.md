# `github.com/boostgo/errorx`

# Get started

```go
package main

import (
	"errors"
	"fmt"

	"github.com/boostgo/errorx"
)

type Person struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

func main() {
	// create error
	err := oops()

	// wrap error
	errorx.Wrap("one more type", &err, "One more message")

	fmt.Println(err)
	// out: [one more type - some type] One more message - some error: not found - conflict. Context: ctx3=3;ctx5={Johnson John};ctx1=value1;ctx2=value2;

	fmt.Println("is not found (by errorx):", errorx.Is(err, errorx.ErrNotFound))           // true
	fmt.Println("is not found (by origin):", errors.Is(err, errorx.ErrNotFound))           // true
	fmt.Println("is unauthorized (by errorx):", errorx.Is(err, errorx.ErrUnauthorized))    // false
	fmt.Println("is unauthorized (by by origin):", errors.Is(err, errorx.ErrUnauthorized)) // false
}

func oops() error {
	return errorx.
		New("some error").
		SetType("some type").
		SetContext(map[string]any{
			"ctx1": "value1",
			"ctx2": "value2",
			"ctx3": 3,
		}).
		AddContext("ctx4", nil).
		AddContext("ctx5", Person{
			LastName:  "Johnson",
			FirstName: "John",
		}).
		SetError(
			errorx.ErrNotFound,
			errorx.ErrConflict,
		)
}

```

# Error

### Output

If you want to beautify your error response and hide some inner levels of message (stack) you can use "onlyFirst" optional variable
```go
fmt.Println(errorx.Get(err).Message(1)) // "One more message"
fmt.Println(errorx.Get(err).Type(1))    // "one more type"
```

# Try

Try-Catch like in Java, C#, etc...

```go
package main

import (
	"fmt"

	"github.com/boostgo/errorx"
)

func main() {
	err := errorx.Try(func() error {
		panic("test")
		return nil
	})
	if err != nil {
		fmt.Println("err:", err)
	}

	// err: PANIC RECOVER: test. Context:
	// goroutine 1 [running]:... <TRACE>
}

```
