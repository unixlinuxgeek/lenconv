### Ft to Meter and Meter to Ft  converter usage

1. Create app.go file (touch app.go):
```go
package main

import(
  "os"
  "fmt"
  "github.com/unixlinuxgeek/lenconv"
)

func main() {
  if len(os.Args) > 1 {
    m := lenconv.FtToMet(1)
    fmt.Printf("%v ft. equal %v\n", os.Args[1], m)
  }
}
```

2. Installing dependency:
```shell
go get github.com/unixlinuxgeek/lenconv
```

3. Run:
```shell
go run ./app.go
```

Example:
```shell
$ go run  ./app.go 1
1 ft. equal 0.30478512648582745 Meter
```
