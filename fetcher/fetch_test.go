package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	result, err := Fetch("http://admin.test.17zdd.cn")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result) )
}
