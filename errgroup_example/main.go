package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/xpzouying/gofake/errgroup"
)

func main() {
	var urls = []string{
		"http://www.baidu.com",
		"https://bbs.hupu.com",
	}

	g := errgroup.New()

	client := http.Client{Timeout: 3 * time.Second}

	for _, url := range urls {

		url := url

		g.Go(func() error {
			resp, err := client.Get(url)
			if err != nil {
				log.Printf("http get [%s] error: %v", url, err)
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Println("status code is not ok")
				return errors.New("status is not 200")
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// valid domain
			ss := strings.Split(url, ".")
			domain := strings.Join(ss[len(ss)-2:], ".")

			if valid := strings.Contains(string(data), domain); valid {
				log.Printf("request url succ: %s", url)
			} else {
				log.Printf("request url failed: %s", url)
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Printf("http get error: %v", err)
		return
	}

	log.Printf("http get succ")
}
