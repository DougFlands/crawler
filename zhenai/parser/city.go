package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `(?:<a href=")(http://m.zhenai.com/u/[0-9a-z]+)(?:" class="left-item" data-v-079f4246><div class="right-item" data-v-079f4246><div class="f-nickname" data-v-079f4246>)(?:\n)(.{0,60})`

// 获取城市
// (?:false,"linkContent":")(.{0,8})(?:","linkURL":")(http:\\u002F\\u002Fm.zhenai.com\\u002Fzhenghun\\u002F[0-9a-z]+)"

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)


	matches := re.FindAllSubmatch(contents, -1)


	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
