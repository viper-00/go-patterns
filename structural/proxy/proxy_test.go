package proxy

import "testing"

func TestPattern(t *testing.T) {
	var responseCode int
	var body string
	server := newNginxServer()

	testUrl := "/getStatus"
	testMethod := "GET"
	errUrl := "/gerError"

	responseCode, body = server.handleRequest(testUrl, testMethod)
	t.Logf("testUrl: %s testMethod %s responseCode: %d body: %s\n", testUrl, testMethod, responseCode, body)

	responseCode, body = server.handleRequest(testUrl, testMethod)
	t.Logf("testUrl: %s testMethod %s responseCode: %d body: %s\n", testUrl, testMethod, responseCode, body)

	responseCode, body = server.handleRequest(testUrl, testMethod)
	t.Logf("testUrl: %s testMethod %s responseCode: %d body: %s\n", testUrl, testMethod, responseCode, body)

	responseCode, body = server.handleRequest(errUrl, testMethod)
	t.Logf("testUrl: %s testMethod %s responseCode: %d body: %s\n", testUrl, testMethod, responseCode, body)
}
