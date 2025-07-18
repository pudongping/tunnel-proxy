package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// https://www.proxy-list.download/SOCKS5
	// proxyUrl, err := url.Parse("socks5://192.111.130.2:4145")
	// ips := []string{
	// 	"37.27.253.44:8099",
	// 	"185.49.31.207:8081",
	// 	"123.141.181.1:5031",
	// 	"211.83.1.47:18000",
	// 	"118.190.210.227:3128",
	// 	"47.97.70.185:80",
	// }
	proxyUrl, err := url.Parse("http://127.0.0.1:9527")
	if err != nil {
		panic(err)
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}
	resp, err := client.Get("https://httpbin.org/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Failed to get a valid response")
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Response:", string(content))

}
