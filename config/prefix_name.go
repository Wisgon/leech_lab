package config

// 神经元的Prefix命名原则是：尽量描述清楚功能，如果与位置有关，还要描述位置
// 不一定要这个命名原则，反正只要描述清楚神经元的功能或位置就行
var PrefixArea = []string{
	"eye",
	"mouth",
	"nose",
	"sense",
	"valuate",
	"skin",
	"consciousness",
	"movement",
	"reading",
	"talking",
	"hippocampus", // 海马体
}
var PrefixNeureType = []string{
	// "inborn",   // 天生的
	// "acquired", // 后天获得的
}
var PrefixSkinAndSenseType = []string{
	"normalTemperature",
	"hotTemperature",
	"coldTemperature",
	"extremelyHotTemp",
	"extremelyColdTemp",
	"normalPress",
	"biggerPress",
	"extremelyPress",
}
var PrefixSenseType = []string{
	"touchType",
	"painfulType", // painfulType had more higher weight
}

// leech config----------------------------------------------------------------------------------
var SkinAndSenseNeurePosition = []string{
	"leftFrontUp",
	"leftFrontDown",
	"leftMiddleUp",
	"leftMiddleDown",
	"leftBackUp",
	"leftBackDown",
	"rightFrontUp",
	"rightFrontDown",
	"rightMiddleUp",
	"rightMiddleDown",
	"rightBackUp",
	"rightBackDown",
}
var Movements = []string{
	"moveLeftFrontUp",
	"moveLeftFrontDown",
	"moveLeftMiddleUp",
	"moveLeftMiddleDown",
	"moveLeftBackUp",
	"moveLeftBackDown",
	"moveRightFrontUp",
	"moveRightFrontDown",
	"moveRightMiddleUp",
	"moveRightMiddleDown",
	"moveRightBackUp",
	"moveRightBackDown",
}
