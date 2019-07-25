package mEmIFy

import (
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

// SpongebobCaseSeed taKes ThE OriGiNaL stRiNg anD ReTuRns It LiKE the SpOnGeBob MocK mEMe. It TaKeS A SEed FoR Y'All LolI NaZIs.
func SpongebobCaseSeed(s string, seed int64) string {
	rand.Seed(seed)
	originalString := s
	length := len(originalString)
	spongebobSlice := make([]rune, length)
	for i, v := range originalString {
		probs := rand.Intn(100)
		if probs > 51 {
			spongebobSlice[i], _ = utf8.DecodeRuneInString(strings.ToUpper(string(v)))
		} else {
			spongebobSlice[i], _ = utf8.DecodeRuneInString(strings.ToLower(string(v)))
		}
	}
	return string(spongebobSlice)
}

// SpongebobCase iS JuSt A wRaPPeR to SpongebobCaseSeed BuT TAkeS CaRe Of SeEdiNg RaNDoM FoR You, YA NEwb
func SpongebobCase(s string) string {
	return SpongebobCaseSeed(s, time.Now().UnixNano())
}
