package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe1 = `cityData.+footerData`
const cityListRe2 = `(?:linkContent:")(.{0,8})(?:",linkURL:")(http://m.zhenai.com/zhenghun/[0-9a-z]+)"`

func ParseCityList(contents []byte) engine.ParseResult {
	re1 := regexp.MustCompile(cityListRe1)
	re2 := regexp.MustCompile(cityListRe2)

	matches1 := re1.FindAll(contents, -1)[0]
	matches2 := re2.FindAllSubmatch(matches1, -1)

	result := engine.ParseResult{}
	for _, m := range matches2 {
		result.Items = append(result.Items, string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[2]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
