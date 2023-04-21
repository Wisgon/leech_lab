package test

import (
	"graph_robot/config"
	"strings"
	"testing"
)

func TestCombinePrefix(t *testing.T) {
	prefix := config.GetAllPrefix()
	skinPrefix := []string{}
	for _, v := range prefix {
		if strings.Contains(v, "movement") {
			skinPrefix = append(skinPrefix, v)
			t.Logf("prefix:%s\n", v)
		}
	}
	t.Log("len:", len(skinPrefix))
}
