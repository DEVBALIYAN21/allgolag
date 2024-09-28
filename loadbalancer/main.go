package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type LoadBalancer struct {
	servers []*url.URL
	counter uint32
}

func NewLoadBalancer(targets []string) *LoadBalancer {
	var servers []*url.URL
	for _, target := range targets {
		url, err := url.Parse(target)
		if err != nil {
			panic(err)
		}
		servers = append(servers, url)
	}

	return &LoadBalancer{servers: servers}
}

func (lb *LoadBalancer) getNextServer() *url.URL {
	server := lb.servers[atomic.AddUint32(&lb.counter, 1)%uint32(len(lb.servers))]
	fmt.Println(server)
	return server
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := lb.getNextServer()
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func main() {
	// List of backend servers (URL shortener instances)
	servers := []string{
		"http://localhost:8081", // Instance 1
		"http://localhost:8082", // Instance 2
	}

	lb := NewLoadBalancer(servers)
	http.ListenAndServe(":8080", lb) // Load balancer running on port 8080
}
