package check

import (
	"context"
	"errors"
	"github.com/fatih/color"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	infoColor  = color.New(color.FgGreen).PrintfFunc()
	errorColor = color.New(color.FgRed).PrintfFunc()
	debugColor = color.New(color.FgCyan).PrintfFunc()
	mu         sync.Mutex
)

func CheckHTTP(rawURL, rawProxy string, timeout int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	URL, err := url.Parse(rawURL)
	if err != nil {
		return false, errors.New("invalid URL")
	}

	Proxy, err := url.Parse("http://" + rawProxy)
	if err != nil {
		return false, errors.New("invalid Proxy")
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(Proxy),
		},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), nil)
	if err != nil {
		return false, errors.New("invalid request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("invalid request")
	}

	if resp.StatusCode != http.StatusOK {
		return false, errors.New(resp.Status)
	}

	return true, nil
}

func CheckSocks5(rawURL, rawProxy string, timeout int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	URL, err := url.Parse(rawURL)
	if err != nil {
		return false, errors.New("invalid URL")
	}

	Proxy, err := url.Parse("socks5://" + rawProxy)
	if err != nil {
		return false, errors.New("socks5 proxy is not valid")
	}

	dialer, err := proxy.SOCKS5("tcp", Proxy.Host, nil, proxy.Direct)
	if err != nil {
		return false, errors.New("socks5 proxy error: " + err.Error())
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), nil)
	if err != nil {
		return false, errors.New("invalid request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("invalid responce")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, errors.New(resp.Status)
	}

	return true, nil
}
