package html

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/foolin/pagser"
)

func Int(node *goquery.Selection, args ...string) (out interface{}, err error) {
	return strconv.Atoi(node.Text())
}

func Html2Struct(rawHtml string, v interface{}) error {
	var p = pagser.New()
	p.RegisterFunc("Int", Int)
	return p.Parse(v, rawHtml)
}
