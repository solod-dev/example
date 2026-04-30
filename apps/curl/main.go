// Make an HTTP GET request using libcurl
// and print the response code and status.
package main

import "solod.dev/so/c"

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
	curl_global_init(c.Val[int]("CURL_GLOBAL_DEFAULT"))
	handle := curl_easy_init()
	if handle == nil {
		curl_global_cleanup()
		return httpResponse{code: c.Val[int]("CURLE_FAILED_INIT")}
	}

	curl_easy_setopt_string(handle, c.Val[int32]("CURLOPT_URL"), url)
	curl_easy_setopt_long(handle, c.Val[int32]("CURLOPT_FOLLOWLOCATION"), 1)
	curl_easy_setopt_long(handle, c.Val[int32]("CURLOPT_TIMEOUT"), timeout)

	code := curl_easy_perform(handle)
	if code != c.Val[int32]("CURLE_OK") {
		curl_easy_cleanup(handle)
		curl_global_cleanup()
		return httpResponse{code: int(code)}
	}

	var status int
	curl_easy_getinfo(handle, c.Val[int32]("CURLINFO_RESPONSE_CODE"), &status)
	curl_easy_cleanup(handle)
	curl_global_cleanup()
	return httpResponse{code: int(code), status: status}
}
