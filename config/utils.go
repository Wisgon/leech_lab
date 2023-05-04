package config

import (
	"graph_robot/utils"
	"strings"
)

func getOnePrefix(oldPrefix []string, somePrefix []string) []string {
	var prefix []string
	if len(oldPrefix) == 0 {
		for j := 0; j < len(somePrefix); j++ {
			prefix = append(prefix, somePrefix[j])
		}
	} else {
		if len(somePrefix) == 0 {
			return oldPrefix
		}
		for i := 0; i < len(oldPrefix); i++ {
			for j := 0; j < len(somePrefix); j++ {
				prefix = append(prefix, oldPrefix[i]+PrefixNameSplitSymbol+somePrefix[j])
			}
		}
	}
	return prefix
}

func GetAllPrefix() (prefix []string) {
	PrefixAreaKeys, PrefixNeureTypeKeys := utils.GetMapKeys(PrefixArea), utils.GetMapKeys(PrefixNeureType)
	prefix = getOnePrefix(prefix, PrefixAreaKeys)
	prefix = getOnePrefix(prefix, PrefixNeureTypeKeys)

	newPrefix := []string{}
	for i := 0; i < len(prefix); i++ {
		if strings.Contains(prefix[i], PrefixArea["skin"]) {
			skinPrefix := combinePrefixSkin(prefix[i])
			prefix[i] = skinPrefix[0]                        // replace the old prefix[i] to new prefix
			newPrefix = append(newPrefix, skinPrefix[1:]...) // add rest new prefix
		} else if strings.Contains(prefix[i], PrefixArea["muscle"]) {
			musclePrefix := combinePrefixMuscle(prefix[i])
			prefix[i] = musclePrefix[0]
			newPrefix = append(newPrefix, musclePrefix[1:]...)
		} else if strings.Contains(prefix[i], PrefixArea["sense"]) {
			sensePrefix := combinePrefixSense(prefix[i])
			prefix[i] = sensePrefix[0]
			newPrefix = append(newPrefix, sensePrefix[1:]...)
		} else if strings.Contains(prefix[i], PrefixArea["valuate"]) {
			valuatePrefix := combinePrefixValuate(prefix[i])
			prefix[i] = valuatePrefix[0]
			newPrefix = append(newPrefix, valuatePrefix[1:]...)
		}
	}
	prefix = append(prefix, newPrefix...)
	return
}

func combinePrefixSkin(skinPrePrefix string) (skinPrefix []string) {
	for t := range PrefixSkinAndSenseType {
		for f := range SkinAndSenseNeurePosition {
			skinPrefix = append(skinPrefix, skinPrePrefix+PrefixNameSplitSymbol+t+PrefixNameSplitSymbol+f)
		}
	}
	return
}

func combinePrefixMuscle(musclePrePrefix string) (musclePrefix []string) {
	for m := range Movements {
		musclePrefix = append(musclePrefix, musclePrePrefix+PrefixNameSplitSymbol+m)
	}
	return
}

func combinePrefixSense(sensePrePrefix string) (sensePrefix []string) {
	for x := range PrefixSkinAndSenseType {
		for z := range SkinAndSenseNeurePosition {
			sensePrefix = append(sensePrefix, sensePrePrefix+PrefixNameSplitSymbol+x+PrefixNameSplitSymbol+z)
		}
	}
	return
}

func combinePrefixValuate(valuatePrePrefix string) (valuatePrefix []string) {
	for x := range PrefixValuateSource {
		for y := range PrefixValuateLevel {
			valuatePrefix = append(valuatePrefix, valuatePrePrefix+PrefixNameSplitSymbol+x+PrefixNameSplitSymbol+y)
		}
	}
	return
}
