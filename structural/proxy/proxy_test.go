package proxy

import "testing"

func TestPattern(t *testing.T) {
	server := newNginxServer()

	appVersionUrl := "/v1/app/version"
	userUpdateUrl := "/v1/user/update"

	getMethod := "GET"
	postMethod := "POST"

	httpCode, body := server.handleRequest(appVersionUrl, getMethod)
	t.Logf("Url: %s HttpCode: %d Body: %s\n", appVersionUrl, httpCode, body)

	httpCode, body = server.handleRequest(userUpdateUrl, postMethod)
	t.Logf("Url: %s HttpCode: %d Body: %s\n", userUpdateUrl, httpCode, body)

	httpCode, body = server.handleRequest(userUpdateUrl, getMethod)
	t.Logf("Url: %s HttpCode: %d Body: %s\n", userUpdateUrl, httpCode, body)

}
