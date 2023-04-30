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
	"muscle",
	"reading",
	"talking",
	"hippocampus", // 海马体
}
var PrefixNeureType = []string{
	"common",     // 普通神经元
	"regulate",   // 调节神经元，主要用于突触增强
	"inhibitory", // 抑制神经元，用于抑制突触
}
var PrefixSkinAndSenseType = []string{
	"normalTemperature",
	"biggertHotTemperature",
	"biggerColdTemperature",
	"extremelyHotTemp",
	"extremelyColdTemp",
	"normalPress",
	"biggerPress",
	"extremelyPress",
}
var PrefixSenseType = []string{
	"senseType",   // 连接所有skin的所有类型神经元
	"painfulType", // 连接skin的extremely类型的神经元
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
