package proxy

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Proxy_pattern
//
// Provides a surrogate for an object to control it's actions.

type Server interface {
	handleRequest(string, string) (int, string)
}

type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	if allowed := n.checkRateLimiting(url); !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	fmt.Print(n.rateLimiter[url])
	if _, ok := n.rateLimiter[url]; !ok {
		n.rateLimiter[url] = 1
		return true
	}
	if n.rateLimiter[url]+1 > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] += 1
	return true
}

type application struct{}

func (a *application) handleRequest(url, method string) (int, string) {
	var testUrl, testMethod string = "/getStatus", "GET"
	if url == testUrl && method == testMethod {
		return 200, "OK"
	}

	return 404, "Not Found"
}
