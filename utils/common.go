package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"path"
	"path/filepath"
	"runtime"

	"golang.org/x/exp/constraints"
)

type Synapse interface {
	GetNextId() string
}

func GetUniqueId(nowNano int64) string {
	randStr := fmt.Sprint(rand.Intn(1000000) + 100000) // 再加上這個以保證絕對不會重複
	data := []byte(fmt.Sprintf("%d%s", nowNano, randStr))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GetProjectRoot() string {
	// get the root path of project
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func RemoveValueFromSlice[T constraints.Integer | constraints.Float | string](value T, s *[]T) {
	for i, v := range *s {
		if v == value {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
}

func RemoveValueFromSynapse[T Synapse](value string, s *[]T) {
	for i, v := range *s {
		if v.GetNextId() == value {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
}
