package urlx

import . "github.com/pkg4go/assert"
import "testing"
import "fmt"

type Case struct {
	From   string
	To     string
	Expect string
}

var cases = []Case{{
	From:   "http://example.com",
	To:     "/a",
	Expect: "http://example.com/a",
}, {
	From:   "http://example.com",
	To:     "/a/b",
	Expect: "http://example.com/a/b",
}, {
	From:   "http://example.com",
	To:     "/a/../b",
	Expect: "http://example.com/b",
}, {
	From:   "http://example.com",
	To:     "/a/..",
	Expect: "http://example.com/",
}, {
	From:   "http://example.com",
	To:     "a",
	Expect: "http://example.com/a",
}, {
	From:   "http://example.com/a",
	To:     "b",
	Expect: "http://example.com/b",
}, {
	From:   "http://example.com/a",
	To:     "/b",
	Expect: "http://example.com/b",
}, {
	From:   "http://example.com/a",
	To:     "b/c",
	Expect: "http://example.com/b/c",
}}

func TestResolve(t *testing.T) {
	a := A{t}

	for _, c := range cases {
		result, _ := Resolve(c.From, c.To)
		a.Equal(result, c.Expect,
			fmt.Sprintf("from: %s, to: %s, result: %s, expect: %s",
				c.From, c.To, result, c.Expect))
	}
}
