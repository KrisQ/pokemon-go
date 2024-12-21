# pokemon-go

Small project to dive into Golang 🐻

## Todos

- [x] map command
- [x] make a new service package -> manages your PokeAPI interactions

Example:
```go
// File: internal/utils/helpers.go
package utils

import "fmt"

// HelperFunction is an internal utility function
func HelperFunction() {
    fmt.Println("This is a helper function")
}
```

```go
// File: main.go
package main

import (
    "myapp/internal/utils" // Accessible
    "myapp/pkg/service"
)

func main() {
    utils.HelperFunction()
    service.DoService()
}
```

- [x] displays the names of 20 location areas
- [ ] each subsequent call to map should display the next 20 locations, and so on

Example usage:

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

Pointers:
- [pokemon location-area endpoint](https://pokeapi.co/docs/v2#location-areas)
- update all commands to accept a pointer to a "config" struct as param. this will contain the next and previous URL that we need to paginate locations
- [get request example](https://pkg.go.dev/net/http#example-Get)
- [unmarshal \[\]byte](https://blog.boot.dev/golang/json-golang/?_gl=1*1r9ptw8*_gcl_au*NTIyMDY1NTk4LjE3MzQ2MTMxNjM.*_ga*MjY2MzgxMjM0LjE3MzQ2MTMxNjM.*_ga_M7P2PBGN8N*MTczNDc4MDMzNS40LjEuMTczNDc4MTA4Ny42MC4wLjMyOTcyMDk4NA..#example-unmarshal-json-to-struct-decode)
- use curl to make GET request, very convenient

- [x] mapb command
- [ ] displays the 20 previous locations
- [ ] edge case: if still on the first page -> print "you're on the first page"

Tips:
- https://jsonlint.com/
- https://mholt.github.io/json-to-go/ -> use it to generate the structs you'll need to parse the PokeAPI response. Keep in mind it sometimes can't know the exact type of a field that you want, because there are multiple valid options. For nullable strings, use *string.
