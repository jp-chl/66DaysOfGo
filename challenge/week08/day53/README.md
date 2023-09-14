# Day 53 of #66DaysOfGo

_Last update:  Sep 13, 2023_.

---

Today, I walked through a basic example of the SingleFlight x Package.

---

> Based on _https://levelup.gitconnected.com/optimize-your-go-code-using-singleflight-3f11a808324_

---

## Intro

The package singleflight provides a duplicate function call suppression mechanism.

It contains a `Group` struct that provides a Do function. Do executes and returns the results of the given function, making sure that only one execution is in-flight for a given key at a time. If a duplicate comes in, the duplicate caller waits for the original to complete and receives the same results. The return value shared indicates whether v was given to multiple callers.

## Demo

Create a simple server that returns a message upon a `name` query parameter.
Code available in the [normal](./normal/server.go) directory.

Create a client that makes 5 requests to the server.
Its code is also available in the [normal](./normal/client.go) directory.

```bash
cd normal
```

```bash
go run server.go
```

Now run the client. The server output must print 5 messages.

```bash
go run client.go
2023/09/13 21:49:25 Hi there! You requested something
2023/09/13 21:49:25 Hi there! You requested something
2023/09/13 21:49:25 Hi there! You requested something
2023/09/13 21:49:25 Hi there! You requested something
2023/09/13 21:49:25 Hi there! You requested something
```

```bash
# Server output
[DEBUG] processing request..
[DEBUG] processing request..
[DEBUG] processing request..
[DEBUG] processing request..
[DEBUG] processing request..
```

---

Now add the SingleFlight to the server. **It might reduce the server processing**.
Server code is in the [singleflight](./singleflight/server.go) folder (the client code remains).

SingleFlight expects a key within the `Do` function. This acts as a fast cache.

Extract:

```go
// ...
import (
  "fmt"
  "net/http"

  "golang.org/x/sync/singleflight"
)

// ...
var g = singleflight.Group{}

func main() {
  http.HandleFunc("/api/v1/get_something", func(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    response, _, _ := g.Do(name, func() (interface{}, error) {
      result := processingRequest(name)
      return result, nil
    })
// ...
```

```bash
cd ../singleflight
```

```bash
$ go mod init example.com
go: creating new go.mod: module example.com
go: to add module requirements and sums:
  go mod tidy
```

```bash
$ go get -v all
go: added golang.org/x/sync v0.3.0
```

Run the server code

```bash
go run server.go
```

```bash
$ go run client.go
2023/09/13 22:39:23 Hi there! You requested something
2023/09/13 22:39:23 Hi there! You requested something
2023/09/13 22:39:23 Hi there! You requested something
2023/09/13 22:39:23 Hi there! You requested something
2023/09/13 22:39:23 Hi there! You requested something
```

```bash
# Server output (same 5 times)
[DEBUG] processing request..
[DEBUG] processing request..
[DEBUG] processing request..
[DEBUG] processing request..
[DEBUG] processing request..
```

```bash
$ go run client.go
2023/09/13 22:39:30 Hi there! You requested something
2023/09/13 22:39:30 Hi there! You requested something
2023/09/13 22:39:30 Hi there! You requested something
2023/09/13 22:39:30 Hi there! You requested something
2023/09/13 22:39:30 Hi there! You requested something
```

```bash
# Server output (now 4)
[DEBUG] processing request..
[DEBUG] processing request..
[DEBUG] processing request..
[DEBUG] processing request..
```

---

## References

- [https://levelup.gitconnected.com/optimize-your-go-code-using-singleflight-3f11a808324](https://levelup.gitconnected.com/optimize-your-go-code-using-singleflight-3f11a808324)
- [https://pkg.go.dev/golang.org/x/sync/singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight)
