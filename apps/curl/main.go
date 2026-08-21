// Make an HTTP GET request using libcurl
// and print the response code and status.
//
// Usage example:
//
//	so run apps/curl
package main

func main() {
	resp := httpGet("https://example.com", 3)
	println("curl code:", resp.code)
	println("http status:", resp.status)
}

type httpResponse struct {
	code   int
	status int
}

func httpGet(url string, timeout int) httpResponse {
	curl_global_init(CURL_GLOBAL_DEFAULT)
	handle := curl_easy_init()
	if handle == nil {
		curl_global_cleanup()
		return httpResponse{code: CURLE_FAILED_INIT}
	}

	curl_easy_setopt_string(handle, CURLOPT_URL, url)
	curl_easy_setopt_long(handle, CURLOPT_FOLLOWLOCATION, 1)
	curl_easy_setopt_long(handle, CURLOPT_TIMEOUT, timeout)

	code := curl_easy_perform(handle)
	if code != CURLE_OK {
		curl_easy_cleanup(handle)
		curl_global_cleanup()
		return httpResponse{code: int(code)}
	}

	var status int
	curl_easy_getinfo(handle, CURLINFO_RESPONSE_CODE, &status)
	curl_easy_cleanup(handle)
	curl_global_cleanup()
	return httpResponse{code: int(code), status: status}
}
