package mEmIFy

import (
	"math/rand"
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
