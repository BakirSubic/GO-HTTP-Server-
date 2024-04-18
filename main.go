package main

import (
	"strings"
)

type httpRequest struct {
	Method      string
	Path        string
	HTTPVersion string
	Headers     map[string]string
}

func readHTTPRequest(request string) httpRequest {
	lines := strings.Split(request, "\n")

	requestLine := lines[0]

	methodEnd := strings.IndexByte(requestLine, ' ')

	if methodEnd == -1 {
		panic("400 Bad Request")
	}

	method := requestLine[:methodEnd]

	pathAndVersion := strings.TrimSpace(requestLine[methodEnd:])

	spaceIndex := strings.IndexByte(pathAndVersion, ' ')

	if spaceIndex == -1 {
		panic("400 Bad Request")
	}

	path := pathAndVersion[:spaceIndex]

	lastSpaceIndex := strings.LastIndex(requestLine, " ")

	if lastSpaceIndex == -1 {
		panic("400 Bad Request")
	}

	httpVersion := strings.TrimSpace(requestLine[lastSpaceIndex:])

	if method == "" || path == "" || httpVersion == "" {
		panic("400 Bad Request")
	}

	requestHeaders := make(map[string]string)

	for i := 1; lines[i] != ""; i++ {
		keyName := lines[i][:strings.IndexByte(lines[i], ':')]
		requestHeaders[keyName] = strings.TrimPrefix(lines[i], keyName+": ")
		if requestHeaders[keyName] == " " {
			panic("400 Bad Request")
		}
	}

	req := httpRequest{
		Method:      method,
		Path:        path,
		HTTPVersion: httpVersion,
		Headers:     requestHeaders,
	}

	println("Method:", req.Method)
	println("Path:", req.Path)
	println("HTTP Version:", req.HTTPVersion)
	for key, value := range req.Headers {
		println(key+":", value)
	}

	return req
}

func main() {
	readHTTPRequest("GET /hello.htm HTTP/1.1\n" +
		"User-Agent: Mozilla/4.0 (compatible; MSIE5.01; Windows NT)\n" +
		"Host: www.tutorialspoint.com\n" +
		"Accept-Language: en-us\n" +
		"Accept-Encoding: gzip, deflate\n" +
		"Connection: Keep-Alive\n" +
		"")
}
