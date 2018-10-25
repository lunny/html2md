package html2md

import "regexp"

const (
	Foreign = iota
	Void
)

type ReplaceFunc func(string, []string) string

type Rule struct {
	Patterns    []string
	Tp          int
	Replacement ReplaceFunc
}

func AttrRegExp(attr string) *regexp.Regexp {
	return regexp.MustCompile(attr + `\s*=\s*["']?([^"\"']*)["']?`)
}

var (
	rules    = make(map[string]*Rule)
	converts = make(map[int]func(string) string)
)

func AddRule(name string, rule *Rule) {
	rules[name] = rule
}

func AddConvert(order int, f func(string) string) {
	converts[order] = f
}
