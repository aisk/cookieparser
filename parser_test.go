package cookieparser

import (
	"testing"
)

func TestParseAndDump(t *testing.T) {
	origin := "yummy_cookie=choco; tasty_cookie=strawberry"
	cookies := Parse(origin)
	if len(cookies) != 2 {
		t.Fatal("excepted two cookies been parsed")
	}
	if Dump(cookies) != origin {
		t.Fatal("unexcepted dump value")
	}
}
