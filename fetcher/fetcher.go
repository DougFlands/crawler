package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//var rateLimiter = time.Tick(5 * time.Millisecond)


func Fetch(url string) ([]byte, error) {
	//<- rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil,
			fmt.Errorf("wrong status code %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}