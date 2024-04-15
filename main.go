package main

import (
	"fmt"
	"strings"
)

func readHTTPRequest(request string) {
	lines := strings.Split(request, "\n")

	requestLine := lines[0]

	requestHeader := make(map[string]string)

	for i := 1; lines[i] != ""; i++ {
		keyName := lines[i][:strings.IndexByte(lines[i], ':')]
		requestHeader[keyName] = strings.TrimPrefix(lines[i], keyName+": ")
	}

	fmt.Println("RequestLine: ")
	fmt.Println(requestLine, "\n")

	fmt.Println("Request Header:")
	for key, value := range requestHeader {
		fmt.Println(key, ":", value, "\n")
	}

}

func main() {
	fmt.Println("Test")

	readHTTPRequest("GET /hello.htm HTTP/1.1\n" +
		"User-Agent: Mozilla/4.0 (compatible; MSIE5.01; Windows NT)\n" +
		"Host: www.tutorialspoint.com\n" +
		"Accept-Language: en-us\n" +
		"Accept-Encoding: gzip, deflate\n" +
		"Connection: Keep-Alive\n" +
		"")
}
