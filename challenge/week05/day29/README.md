# Day 29 of #66DaysOfGo

_Last update:  Aug 17, 2023_.

---

Today, I've continued with the Design Patterns series, with the Proxy.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Proxy pattern

This pattern involves using a proxy class to provide an interface to another object, which could be a network connection, memory-intensive object, or expensive resource. The proxy acts as a wrapper or agent and can offer additional functionality such as caching or precondition checks. It allows controlled access to objects and enhances their capabilities.

### UML diagram

<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/W3sDesign_Proxy_Design_Pattern_UML.jpg" alt="Proxy Pattern UML example" width="550"/>

### Code example

```go
package main

import (
    "fmt"
    "time"
)

const (
    HTTP_GET                       = "GET"
    HTTP_POST                      = "POST"
    HTTP_PUT                       = "PUT"
    HTTP_PATCH                     = "PATCH"
    HTTP_DELETE                    = "DELETE"
    HTTP_RESPONSE_CODE_OK          = 200
    HTTP_RESPONSE_CODE_BAD_REQUEST = 400
    HTTP_RESPONSE_CODE_FORBIDDEN   = 403
    TEST_LOG_TEMPLATE              = "Hitting url: [%s], with method: [%s]...\n"
    TEST_RESPONSE_TEMPLATE         = "RESPONSE: [%s]-[%s]"
)

type RestAPI interface {
    request(string, string) (int, string)
}

type MicroserviceAPI struct {
}

func NewMicroserviceAPI() *MicroserviceAPI {
    return &MicroserviceAPI{}
}

func (m *MicroserviceAPI) request(url string, method string) (int, string) {
    switch method {
    case HTTP_GET:
        m.simulateTimeout(2)
    case HTTP_POST:
        m.simulateTimeout(3)
    case HTTP_PUT:
        m.simulateTimeout(4)
    case HTTP_PATCH:
        m.simulateTimeout(5)
    case HTTP_DELETE:
        response := fmt.Sprintf("Method [%s] is not supported", method)
        fmt.Println(response)
        return HTTP_RESPONSE_CODE_FORBIDDEN, response
    }

    response := fmt.Sprintf(TEST_RESPONSE_TEMPLATE, url, method)
    return HTTP_RESPONSE_CODE_OK, response
}

func (m *MicroserviceAPI) simulateTimeout(timeoutInSeconds uint8) {
    defer func(t time.Time) {
        time.Sleep(time.Duration(timeoutInSeconds) * time.Second)
        fmt.Printf("Operation took %v\n", time.Since(t))
    }(time.Now())
}

type APIGateway struct {
    backend RestAPI
    cache   map[string]string
}

func NewAPIGateway() *APIGateway {
    return &APIGateway{
        backend: NewMicroserviceAPI(),
        cache:   make(map[string]string),
    }
}

func (a *APIGateway) request(url string, method string) (int, string) {
    if method == HTTP_DELETE {
        return HTTP_RESPONSE_CODE_BAD_REQUEST, fmt.Sprintf("%s method is not supported", HTTP_DELETE)
    }

    cacheKey := fmt.Sprintf("[%s][%s]", url, method)
    if cacheValue, present := a.cache[cacheKey]; present {
        defer func(t time.Time) {
            fmt.Printf("Retrieving (%s) from cache...\n", cacheKey)
            fmt.Printf("Operation took %v\n", time.Since(t))
        }(time.Now())
        return HTTP_RESPONSE_CODE_OK, cacheValue
    }

    response_code, response_str := a.backend.request(url, method)

    fmt.Printf("Saving (%s) into cache...\n", cacheKey)
    a.cache[cacheKey] = response_str

    return response_code, response_str
}

func main() {
    demo(NewMicroserviceAPI())

    fmt.Printf("\n==================\nUsing proxy\n==================\n\n")

    demo(NewAPIGateway())
}

func demo(api RestAPI) {
    urlAndMethodsArray := [][2]string{
        {"url1", HTTP_GET},
        {"url2", HTTP_GET},
        {"url1", HTTP_GET},
        {"url2", HTTP_GET},
        {"url3", HTTP_POST},
        {"url4", HTTP_PUT},
        {"url5", HTTP_PATCH},
        {"url6", HTTP_DELETE},
    }

    for i := 0; i < len(urlAndMethodsArray); i++ {
        url, method := urlAndMethodsArray[i][0], urlAndMethodsArray[i][1]
        fmt.Printf(TEST_LOG_TEMPLATE, url, method)
        response_code, response_str := api.request(url, method)
        fmt.Printf("Response: code [%d], string [%s]\n", response_code, response_str)
        fmt.Println()
    }
}
```

```bash
$ go run proxy.go
Hitting url: [url1], with method: [GET]...
Operation took 2.001201242s
Response: code [200], string [RESPONSE: [url1]-[GET]]

Hitting url: [url2], with method: [GET]...
Operation took 2.000285159s
Response: code [200], string [RESPONSE: [url2]-[GET]]

Hitting url: [url1], with method: [GET]...
Operation took 2.000657336s
Response: code [200], string [RESPONSE: [url1]-[GET]]

Hitting url: [url2], with method: [GET]...
Operation took 2.001184656s
Response: code [200], string [RESPONSE: [url2]-[GET]]

Hitting url: [url3], with method: [POST]...
Operation took 3.001053317s
Response: code [200], string [RESPONSE: [url3]-[POST]]

Hitting url: [url4], with method: [PUT]...
Operation took 4.001196175s
Response: code [200], string [RESPONSE: [url4]-[PUT]]

Hitting url: [url5], with method: [PATCH]...
Operation took 5.001208502s
Response: code [200], string [RESPONSE: [url5]-[PATCH]]

Hitting url: [url6], with method: [DELETE]...
Method [DELETE] is not supported
Response: code [403], string [Method [DELETE] is not supported]


==================
Using proxy
==================

Hitting url: [url1], with method: [GET]...
Operation took 2.001193937s
Saving ([url1][GET]) into cache...
Response: code [200], string [RESPONSE: [url1]-[GET]]

Hitting url: [url2], with method: [GET]...
Operation took 2.000979249s
Saving ([url2][GET]) into cache...
Response: code [200], string [RESPONSE: [url2]-[GET]]

Hitting url: [url1], with method: [GET]...
Retrieving ([url1][GET]) from cache...
Operation took 3.299µs
Response: code [200], string [RESPONSE: [url1]-[GET]]

Hitting url: [url2], with method: [GET]...
Retrieving ([url2][GET]) from cache...
Operation took 3.867µs
Response: code [200], string [RESPONSE: [url2]-[GET]]

Hitting url: [url3], with method: [POST]...
Operation took 3.00115966s
Saving ([url3][POST]) into cache...
Response: code [200], string [RESPONSE: [url3]-[POST]]

Hitting url: [url4], with method: [PUT]...
Operation took 4.000150678s
Saving ([url4][PUT]) into cache...
Response: code [200], string [RESPONSE: [url4]-[PUT]]

Hitting url: [url5], with method: [PATCH]...
Operation took 5.001186136s
Saving ([url5][PATCH]) into cache...
Response: code [200], string [RESPONSE: [url5]-[PATCH]]

Hitting url: [url6], with method: [DELETE]...
Response: code [400], string [DELETE method is not supported]
```

---

## References

- [https://refactoring.guru/design-patterns/proxy](https://refactoring.guru/design-patterns/proxy)
