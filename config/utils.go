package config

import (
	"strings"
)

func CombinePrefix() (prefix []string) {
	for i := 0; i < len(PrefixFirst); i++ {
		for j := 0; j < len(PrefixSecond); j++ {
			prefix = append(prefix, PrefixFirst[i]+PrefixNameSplitSymbol+PrefixSecond[j])
		}
	}
	tmp := []string{}
	for i := 0; i < len(prefix); i++ {
		if strings.Contains(prefix[i], "skin") {
			for j := 1; j < len(PrefixSkinThird); j++ {
				tmp = append(tmp, prefix[i]+PrefixNameSplitSymbol+PrefixSkinThird[j])
			}
			prefix[i] = prefix[i] + PrefixNameSplitSymbol + PrefixSkinThird[0]
		}
	}
	prefix = append(prefix, tmp...)
	return
}
