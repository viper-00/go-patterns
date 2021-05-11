package proxy

// wiki: https://en.wikipedia.org/wiki/Proxy_pattern

/**
 * Proxy is a structural design pattern that provide a substitute or placeholder for another object to control acccess to it.
 */

// subject - server
type server interface {
	handleRequest(string, string) (int, string)
}

// proxy - nginx
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
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
		return true
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] += 1
	return true
}

// application - real subject
type application struct{}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/v1/app/version" && method == "GET" {
		// to do something
		return 200, "OK"
	} else if url == "/v1/user/update" && method == "POST" {
		// to do something
		return 200, "user created"
	}

	return 404, "has error with request url and method"
}
