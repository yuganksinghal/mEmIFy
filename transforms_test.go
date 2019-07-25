package mEmIFy

import "testing"

func TestSpongebobCaseSeed(t *testing.T) {
	hello := SpongebobCaseSeed("hello world", 0)
	if hello != "HeLlo WORlD" {
		t.Fail()
	}
}

func TestSpongebobCase(t *testing.T) {
	hello := SpongebobCase("hello world")
	if (hello == "HELLO WORLD") || (hello == "hello world") {
		t.Fail()
	}
}
