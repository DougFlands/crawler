package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	result, err := Fetch("http://m.zhenai.com/zhenghun/aba")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result) )
}
