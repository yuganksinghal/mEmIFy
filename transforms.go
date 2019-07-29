package mEmIFy

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// SpongebobCaseSeed taKes ThE OriGiNaL stRiNg anD ReTuRns It LiKE the SpOnGeBob MocK mEMe. It TaKeS A SEed FoR Y'All CoNTrOl FrEAks.
func SpongebobCaseSeed(s string, seed int64) string {
	rand.Seed(seed)
	spongebobString := ""
	for _, v := range s {
		probs := rand.Intn(100)
		if probs > 51 {
			spongebobString = spongebobString + strings.ToUpper(string(v))
		} else {
			spongebobString = spongebobString + strings.ToLower(string(v))
		}
	}
	return string(spongebobString)
}

// SpongebobCase iS JuSt A wRaPPeR to SpongebobCaseSeed BuT TAkeS CaRe Of SeEdiNg RaNDoM FoR You, YA NEwb
func SpongebobCase(s string) string {
	return SpongebobCaseSeed(s, time.Now().UnixNano())
}

// CCify translates words like protect to protecc
func CCify(s string) (string, error) {
	if s == "" {
		return "", errors.New("The string passed in is empty")
	}
	trimmed := strings.TrimSpace(s)
	split := strings.Split(trimmed, " ")
	regular, err := regexp.Compile(`[a-z]+\W`)
	if err != nil {
		return "", err
	}
	newString := make([]string, 0)
	for _, v := range split {
		keepSafe := v
		if regular.MatchString(v) {
			keepSafe = v[:len(v)-3] + "cc" + string(v[len(v)-1])
		} else if string(v[len(v)-2]) == "c" {
			keepSafe = v[:len(v)-2] + "cc"
		}
		newString = append(newString, keepSafe)
	}
	return strings.Join(newString, " "), nil
}
