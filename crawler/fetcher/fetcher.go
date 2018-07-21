package fetcher

import (
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"

	"bufio"
	"golang.org/x/net/html/charset"
	"net/http"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
	"go_projects/go_crawler_in_action/crawler_distributed/config"
)

var rateLimiter = time.Tick(time.Second / config.QPS)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetching url %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// gbk => utf-8
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
