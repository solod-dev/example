package main

//so:include.c <curl/curl.h>

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
