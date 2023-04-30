package database

import (
	"fmt"
	"graph_robot/config"
	"log"
)

func GetSeqNum(keyPrefix string) string {
	uniqueNum, err := (*seqMap)[keyPrefix].Next()
	if err != nil {
		log.Panic(err)
	}
	return fmt.Sprint(uniqueNum)
}

func GetKeyFromPrefix(keyPrefix string) (key string) {
	uniqueNum := GetSeqNum(keyPrefix)
	key = keyPrefix + config.PrefixNumSplitSymbol + fmt.Sprint(uniqueNum)
	return
}
