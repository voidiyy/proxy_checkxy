# ProxyChecker

ProxyChecker is a simple and efficient command-line tool written in Go for validating HTTP and SOCKS5 proxies. It supports concurrent checking of proxies using goroutines, allowing for fast and scalable performance.
```
.
├── check
│   └── check.go
├── files
│   └── read_write.go
├── go.mod
├── go.sum
├── logs
│   └── info.go
├── main.go
├── proxies.txt
├── proxy_checker
└── valid.txt
```

## Features

- **Concurrent Proxy Checking**: Utilizes goroutines to check multiple proxies simultaneously, improving the speed and efficiency of the tool.
- **Customizable Parameters**: Allows users to specify the target URL, the file containing the list of proxies, the connection timeout, and the number of goroutines to use.
- **Support for HTTP and SOCKS5 Proxies**: Validates both HTTP and SOCKS5 proxies, ensuring they are functional.
- **Detailed Logging**: Provides detailed logging for each proxy check, including information about success or failure, error messages, and response times.
- **Output of Valid Proxies**: Saves the valid proxies to a separate file for further use.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher

### Clone the Repository

```bash
git clone https://github.com/yourusername/ProxyChecker.git
cd ProxyChecker

go build -o proxy_checker .

```

## Usage
```
./proxy_checker <target_url> <file_path> <timeout_seconds> <num_of_goroutines>
```

## Example
```
./proxy_checker https://httpbin.org/get proxies.txt 8 10
  
```

## Hey!

Its just learning project, but you can commit changes to make it more advanced



