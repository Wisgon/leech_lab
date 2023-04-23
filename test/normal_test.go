package test

import (
	"graph_robot/config"
	"strings"
	"testing"
)

func TestGetAllPrefix(t *testing.T) {
	prefix := config.GetAllPrefix()
	i := 0
	for _, v := range prefix {
		if strings.Contains(v, "sense") && strings.Contains(v, "normalTemperature") {
			t.Log(v)
			i++
		}
	}
	t.Log(i, "~~~~")
}
