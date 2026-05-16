# gobwlib
Library to use bitwarden-cli functionality in go. The library acts as a wrapper around the bitwarden-cli. For more information on the cli check out: [bitwarden-cli](https://bitwarden.com/help/cli/).

## Attention
The project is under very heavy development. Features, design and performance may change frequently. Expect incomplete functionality, bugs, documentation and possibly data loss.

## Prerequisites
bitwarden-cli must be installed on the system ([Instructions](https://bitwarden.com/help/cli/#download-and-install)).

## Usage
### Example 1: Create Login
```go
package main

import (
	"fmt"

	"github.com/kuhweide/gobwlib/bw"
)

func main() {
	client, err := bw.NewClient(
		"https://bw.example.com",
		"example@example.com",
		"VERY_secure_AND_sAfe.p4ssw0rd",
	)

	if err != nil {
		panic(err)
	}

	item := bw.NewLoginItem("Example 1")
	item.Login.Username = "Example username"
	item.Login.Password = "Example password"
	item.Login.AddUri(bw.NewUri(bw.MatchTypeStartsWith, "https://example.com"))
	item.AddField(bw.NewField("Field name", "Field value", bw.FieldTypeHidden))

	item, err = client.CreateItem(item)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created item with id [%s]", item.Id)
}
```
