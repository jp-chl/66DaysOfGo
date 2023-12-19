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
