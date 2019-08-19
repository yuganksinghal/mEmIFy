package mEmIFy

import (
	"testing"
)

func TestSpongebobCaseSeed(t *testing.T) {
	hello := SpongebobCaseSeed("hello world", 0) // normal case
	if hello != "HeLlo WORlD" {
		t.Fail()
	}

	empty := SpongebobCaseSeed("", 0) // empty string
	if empty != "" {
		t.Fail()
	}

	capital := SpongebobCaseSeed("HELLO WORLD", 0) // all caps for good measurs
	if capital == "HELLO WORLD" {
		t.Fail()
	}
}

func TestSpongebobCase(t *testing.T) {
	hello := SpongebobCase("hello world")
	if (hello == "HELLO WORLD") || (hello == "hello world") { // make sure they aren't fully capital or fully lower case
		t.Fail()
	}

	empty := SpongebobCaseSeed("", 0) // empty string
	if empty != "" {
		t.Fail()
	}
}

func TestCCify(t *testing.T) {

	hello, err := CCify("hello world") // testing hello world
	if hello != "hello world" || err != nil {
		t.Fail()
	}

	helloagain, err := CCify("hello world!") // testing hello world
	if helloagain != "hello world!" || err != nil {
		t.Fail()
	}

	protecc, err := CCify("Protect and Attack") // standard test case
	if protecc != "Protecc and Attacc" || err != nil {
		t.Fail()
	}

	puncctuation, err := CCify("Protect! and Attack!") // standard test case without punctuation
	if puncctuation != "Protecc! and Attacc!" || err != nil {
		t.Fail()
	}

	empty, err := CCify("") // empty string, guess what, this actually returns an error. _suprise surprise_
	if err == nil || empty != "" {
		t.Fail()
	}
}

func TestSpacity(t *testing.T) {
	hello, err := Spacity("hello") // testing hello
	if hello != "h e l l o" || err != nil {
		t.Fail()
	}

	hi, err := Spacity("hi again") // testing 2 words, should error out
	if err == nil || hi != "" {
		t.Fail()
	}

	emptyString, err := Spacity("") // empty string, should error out
	if err == nil || emptyString != "" {
		t.Fail()
	}
}
