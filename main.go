package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
)

const (
	defaultDir  = "."
	defaultPort = "8080"
)

func main() {
	// Get the current IP address
	ips, err := getCurrentIPs()
	if err != nil {
		fmt.Println(err)
		return
	}

	dir, port := parseFlags()

	dir, err = filepath.Abs(dir)
	if err != nil {
		fmt.Println("Can't find", dir)
		return
	}

	fmt.Printf("Serving files from %s\n\n", dir)

	addressesMessage := makeAddresses(ips, port)
	if len(ips) == 1 {
		fmt.Println("Available at:")
		fmt.Println(addressesMessage)
	} else {
		fmt.Println("Try these links, one of them should share you the folder:")
		fmt.Println(addressesMessage)
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getCurrentIPs() ([]string, error) {
	// Get the local IP address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0)
	// Extract the IP address
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	if len(ips) == 0 {
		return nil, errors.New("Unable to determine local IP address (not connected to network?)")
	}

	return ips, nil
}

func parseFlags() (string, string) {
	dir := defaultDir
	port := defaultPort

	// Parse flags
	flag.Usage = func() {
		fmt.Println("Usage: serve [options] [dir]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	dir = flag.Arg(0)

	flag.StringVar(&port, "port", defaultPort, "Port to serve on")
	flag.Parse()

	return dir, port
}

func makeAddresses(ips []string, port string) string {
	addresses := ""
	for _, ip := range ips {
		addresses += fmt.Sprintf("http://%s:%s\n", ip, port)
	}
	return addresses
}
