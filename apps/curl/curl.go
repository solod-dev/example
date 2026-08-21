package main

//so:include.c <curl/curl.h>
//so:link curl

//so:extern CURL_GLOBAL_DEFAULT
var CURL_GLOBAL_DEFAULT int = 0

//so:extern CURLE_FAILED_INIT
var CURLE_FAILED_INIT int = 0

//so:extern CURLE_OK
var CURLE_OK int32 = 0

//so:extern CURLINFO_RESPONSE_CODE
var CURLINFO_RESPONSE_CODE int32 = 0

//so:extern CURLOPT_FOLLOWLOCATION
var CURLOPT_FOLLOWLOCATION int32 = 0

//so:extern CURLOPT_TIMEOUT
var CURLOPT_TIMEOUT int32 = 0

//so:extern CURLOPT_URL
var CURLOPT_URL int32 = 0

//so:extern CURL
type curlHandle struct{}

func curl_global_init(flags int) int32
func curl_easy_init() *curlHandle
func curl_easy_perform(handle *curlHandle) int32
func curl_easy_getinfo(handle *curlHandle, info int32, value any) int32
func curl_easy_cleanup(handle *curlHandle)
func curl_global_cleanup()

//so:extern curl_easy_setopt
func curl_easy_setopt_long(handle *curlHandle, option int32, value ...int) int32

//so:extern curl_easy_setopt
func curl_easy_setopt_string(handle *curlHandle, option int32, value ...string) int32
