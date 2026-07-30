package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("Failed to read request: %w", err)
	}

	httpMessage := string(data)
	reqLine, err := parseRequestLine(httpMessage)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse requestLine: %w", err)
	}

	return &Request{RequestLine: *reqLine}, nil
}

func parseRequestLine(httpMsg string) (*RequestLine, error) {
	reqLineStr, _, isSepFound := strings.Cut(httpMsg, "\r\n")
	reqLine := strings.Fields(reqLineStr)

	if len(reqLine) != 3 || !isSepFound {
		return nil, fmt.Errorf("Invalid request line")
	}

	method, path, version := reqLine[0], reqLine[1], reqLine[2]
	if strings.Contains(method, "/") || len(method) > 6 || method != strings.ToUpper(method) {
		return nil, fmt.Errorf("Invalid http method")
	}

	if !strings.Contains(path, "/") {
		return nil, fmt.Errorf("Invalid path")
	}

	_, httpVersion, isSepFound := strings.Cut(version, "/")
	if !isSepFound || httpVersion != "1.1" {
		return nil, fmt.Errorf("Invalid http version")
	}

	return &RequestLine{
		Method:        reqLine[0],
		RequestTarget: reqLine[1],
		HttpVersion:   httpVersion,
	}, nil
}
