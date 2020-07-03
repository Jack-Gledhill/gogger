# Gogger - The Go Logger

Gogger is a console logger designed for simplicity and great features at the same time! It has configurable color options for a (configurable) project name and the levels. The level options are Verbose (AKA Debug), Info, Warn, Error and Fatal! 

## Example

Below is an example of how you'd use Gogger in most use-cases. 
```go
package main

import "github.com/BadgeBot/gogger"

func main() {
    log := gogger.New("MyProject", gogger.USDateTimeFormat, true) // project name, datetime format, use colors?

    // Do something here...

    log.Verbose("Hey this is a verbose message, who'da thought?")
}
```

## Documentation

See https://pkg.go.dev/github.com/BadgeBot/gogger for full documentation!