package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	expectedUrls := []string{
		"http://m.zhenai.com/zhenghun/aba", "http://m.zhenai.com/zhenghun/akesu", "http://m.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+"requests; but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("error: %d: %s; was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+"requests; but had %d",
			resultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("error: %d: %s; was %s", i, city, result.Items[i].(string))
		}
	}

}
