package config

func GetAllPrefix() (prefix []string) {
	for i := 0; i < len(PrefixArea); i++ {
		for k := 0; k < len(PrefixNeureType); k++ {
			prefixFS := PrefixArea[i] + PrefixNameSplitSymbol + PrefixNeureType[k]
			if PrefixArea[i] == "skin" {
				skinPrefix := combinePrefixSkin(prefixFS)
				prefix = append(prefix, skinPrefix...)
			} else if PrefixArea[i] == "movement" {
				movementPrefix := combinePrefixMovement(prefixFS)
				prefix = append(prefix, movementPrefix...)
			} else {
				prefix = append(prefix, prefixFS)
			}
		}
	}
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
