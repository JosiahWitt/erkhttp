# erkhttp

Allows [`erk`](https://github.com/JosiahWitt/erk) error kinds to embed HTTP Statuses.

[![GoDoc](https://godoc.org/github.com/JosiahWitt/erkhttp?status.svg)](https://godoc.org/github.com/JosiahWitt/erkhttp)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/JosiahWitt/erkhttp)


## Install
```bash
$ github.com/JosiahWitt/erkhttp
```


## Usage
Embed HTTP Statuses in erk errors:

```go
type ErkItemNotFound {
  erk.DefaultKind
  erkhttp.StatusNotFound
}
```

Fetch the HTTP Status using [`erkhttp.GetHTTPStatus`](https://pkg.go.dev/github.com/JosiahWitt/erkhttp?tab=doc#GetHTTPStatus):
```go
var ErrItemNotFound = erk.New(ErkItemNotFound{}, "item with key '{{.key}}' was not found")

func main() {
  fmt.Println(erkhttp.GetHTTPStatus(ErrItemNotFound)) // Prints 404, true
}
```
