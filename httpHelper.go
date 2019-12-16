package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"time"
)

const (
	HTTPStringFormat = "HTTP/1.1 %d %s\r\nDate: %s\r\nServer: Ev\r\nAccept-Ranges: bytes\r\nConnection: close\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n%s"
)

func NewRawHTTPResponse(status int, contentType string, body []byte) []byte {
	statusText := http.StatusText(status)
	date := time.Now().Format(time.RFC1123)
	return []byte(fmt.Sprintf(HTTPStringFormat, status, statusText, date, contentType, len(body), string(body)))
}

func GetHTTPRequest(req []byte) (httpReq *http.Request, err error) {
	return http.ReadRequest(bufio.NewReader(bytes.NewReader(req)))
}
