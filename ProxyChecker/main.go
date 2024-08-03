package main

import (
	"fmt"
	"log"
	"os"
	"proxy_checker/check"
	"proxy_checker/files"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// go run main.go -url -file_path -request_time_out -num_of_routines
func main() {

	log.Println("main runned")

	var (
		URL      string
		filePath string
		timeout  int
		routines int
		err      error
		valid    []string
		wg       sync.WaitGroup
		info     *log.Logger
		erro     *log.Logger
	)

	info = log.New(os.Stdout, "||| INFO: \t", log.Ltime)
	erro = log.New(os.Stderr, "XXX ERROR: \t", log.Ltime)

	if len(os.Args) != 5 {
		info.Println("Usage of proxy_checker:")
		info.Println("go run main.go -url -file_path -request_time_out -num_of_routines ")
		info.Println("-url :: -URL of connection target, by default :: https://httpbin.org/get")
		info.Println("-file_path :: path of file with proxies format :: ip:port")
		info.Println("-timeout :: seconds before connection timeout INT value :: default value 8")
		info.Println("-routines :: number of goroutines INT value :: default num of CPUs * 2")
		os.Exit(1)
	}

	URL = os.Args[1]
	if URL == "" {
		URL = "https://httpbin.org"
	}
	info.Println("target URL", URL)

	filePath = os.Args[2]
	if filePath == "" {
		filePath = "proxies.txt"
	}
	info.Println("file path:", filePath)

	timeout, err = strconv.Atoi(os.Args[3])
	if err != nil {
		timeout = 8
	}
	info.Println("timeout (seconds) :", timeout)

	routines, err = strconv.Atoi(os.Args[4])
	if err != nil {
		routines = runtime.NumCPU() * 2
	}
	info.Println("Number of goroutines:", routines)
	fmt.Println("=============================================================")
	fmt.Println()

	proxyList, err := files.ReadFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	proxyChan := make(chan string, len(proxyList))

	for _, proxy := range proxyList {
		proxyChan <- proxy
	}
	close(proxyChan)

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for proxy := range proxyChan {
				start := time.Now()
				duration := time.Since(start)

				isValid, err := check.CheckSocks5(URL, proxy, timeout)
				if err != nil || !isValid {
					log.Println("check socks5 FAIL", proxy, err)
				} else if isValid {
					valid = append(valid, proxy)
					log.Println("check socks5 OK", proxy, duration)
				}
				isValid, err = check.CheckHTTP(URL, proxy, timeout)
				if err != nil || !isValid {
					log.Println("check http FAIL", proxy, err)
				} else if isValid {
					log.Println("check http OK", proxy, duration)
				}
				fmt.Println("===========================================================")
			}
		}(i + 1)
	}

	wg.Wait()
	log.Println("Number of valid proxies:", len(valid))

	err = files.WriteToFile("valid.txt", valid)
	if err != nil {
		erro.Fatal(err)
	}
}
