package config

func CombinePrefix() (prefix []string) {
	for i := 0; i < len(PrefixFirst); i++ {
		// for j := 0; j < len(PrefixSecond); j++ {
		// prefixFS := PrefixFirst[i] + PrefixNameSplitSymbol + PrefixSecond[j]
		prefixFS := PrefixFirst[i]
		if PrefixFirst[i] == "skin" {
			skinPrefix := combinePrefixSkin(prefixFS)
			prefix = append(prefix, skinPrefix...)
		} else if PrefixFirst[i] == "movement" {
			movementPrefix := combinePrefixMovement(prefixFS)
			prefix = append(prefix, movementPrefix...)
		} else {
			prefix = append(prefix, prefixFS)
		}
		// }
	}
	return
}

func combinePrefixSkin(skinPrePrefix string) (skinPrefix []string) {
	for _, t := range PrefixSkinThird {
		for _, f := range SkinNeurePosition {
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
