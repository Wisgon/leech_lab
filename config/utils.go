package config

import "strings"

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
	prefix = getOnePrefix(prefix, PrefixArea)
	prefix = getOnePrefix(prefix, PrefixNeureType)

	newPrefix := []string{}
	for i := 0; i < len(prefix); i++ {
		if strings.Contains(prefix[i], "skin") {
			skinPrefix := combinePrefixSkin(prefix[i])
			prefix[i] = skinPrefix[0]
			newPrefix = append(newPrefix, skinPrefix[1:]...)
		} else if strings.Contains(prefix[i], "movement") {
			movementPrefix := combinePrefixMovement(prefix[i])
			prefix[i] = movementPrefix[0]
			newPrefix = append(newPrefix, movementPrefix[1:]...)
		} else if strings.Contains(prefix[i], "sense") {
			sensePrefix := combinePrefixSense(prefix[i])
			prefix[i] = sensePrefix[0]
			newPrefix = append(newPrefix, sensePrefix[1:]...)
		}
	}
	prefix = append(prefix, newPrefix...)
	return
}

func combinePrefixSkin(skinPrePrefix string) (skinPrefix []string) {
	for _, t := range PrefixSkinAndSenseType {
		for _, f := range SkinAndSenseNeurePosition {
			skinPrefix = append(skinPrefix, skinPrePrefix+PrefixNameSplitSymbol+t+PrefixNameSplitSymbol+f)
		}
	}
	return
}

func combinePrefixMovement(movementPrePrefix string) (movementPrefix []string) {
	for _, m := range Movements {
		movementPrefix = append(movementPrefix, movementPrePrefix+PrefixNameSplitSymbol+m)
	}
	return
}

func combinePrefixSense(sensePrePrefix string) (sensePrefix []string) {
	for _, x := range PrefixSkinAndSenseType {
		for _, y := range SkinAndSenseNeurePosition {
			for _, z := range PrefixSenseType {
				sensePrefix = append(sensePrefix, sensePrePrefix+PrefixNameSplitSymbol+x+PrefixNameSplitSymbol+y+PrefixNameSplitSymbol+z)
			}
		}
	}
	return
}
