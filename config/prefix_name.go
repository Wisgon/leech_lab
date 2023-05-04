package config

// 神经元的Prefix命名原则是：尽量描述清楚功能，如果与位置有关，还要描述位置
// 不一定要这个命名原则，反正只要描述清楚神经元的功能或位置就行
var PrefixArea = map[string]string{
	"skin":          "skin",
	"sense":         "sense",
	"muscle":        "muscle",
	"valuate":       "valuate",
	"eye":           "eye",
	"mouth":         "mouth",
	"nose":          "nose",
	"consciousness": "consciousness",
	"reading":       "reading",
	"talking":       "talking",
	"hippocampus":   "hippocampus", // 海马体
}
var PrefixNeureType = map[string]string{
	"common":     "common",     // 普通神经元
	"regulate":   "regulate",   // 调节神经元，主要用于突触增强
	"inhibitory": "inhibitory", // 抑制神经元，用于抑制突触
}
var PrefixSkinAndSenseType = map[string]string{
	"normalTemperature":     "normalTemperature",
	"biggertHotTemperature": "biggertHotTemperature",
	"biggerColdTemperature": "biggerColdTemperature",
	"extremelyHotTemp":      "extremelyHotTemp",
	"extremelyColdTemp":     "eyextremelyColdTempe",
	"normalPress":           "normalPress",
	"biggerPress":           "biggerPress",
	"extremelyPress":        "extremelyPress",
}
var PrefixValuateSource = map[string]string{
	"sense": "sense",
}
var PrefixValuateLevel = map[string]string{
	// valute 0 is common, more positive, fell more better, more negative, more worse
	"valuate2":  "valuate2",
	"valuate1":  "valuate1",
	"valuate0":  "valuate0",
	"valuate-1": "valuate-1",
	"valuate-2": "valuate-2",
}

// leech config----------------------------------------------------------------------------------
var SkinAndSenseNeurePosition = map[string]string{
	"leftFrontUp":     "leftFrontUp",
	"leftFrontDown":   "leftFrontDown",
	"leftMiddleUp":    "leftMiddleUp",
	"leftMiddleDown":  "leftMiddleDown",
	"leftBackUp":      "leftBackUp",
	"leftBackDown":    "leftBackDown",
	"rightFrontUp":    "rightFrontUp",
	"rightFrontDown":  "rightFrontDown",
	"rightMiddleUp":   "rightMiddleUp",
	"rightMiddleDown": "rightMiddleDown",
	"rightBackUp":     "rightBackUp",
	"rightBackDown":   "rightBackDown",
}
var Movements = map[string]string{
	"moveLeftFrontUp":     "moveLeftFrontUp",
	"moveLeftFrontDown":   "moveLeftFrontDown",
	"moveLeftMiddleUp":    "moveLeftMiddleUp",
	"moveLeftMiddleDown":  "moveLeftMiddleDown",
	"moveLeftBackUp":      "moveLeftBackUp",
	"moveLeftBackDown":    "moveLeftBackDown",
	"moveRightFrontUp":    "moveRightFrontUp",
	"moveRightFrontDown":  "moveRightFrontDown",
	"moveRightMiddleUp":   "moveRightMiddleUp",
	"moveRightMiddleDown": "moveRightMiddleDown",
	"moveRightBackUp":     "moveRightBackUp",
	"moveRightBackDown":   "moveRightBackDown",
}
