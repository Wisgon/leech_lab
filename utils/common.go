package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"golang.org/x/exp/constraints"
)

func GetUniqueId(nowNano int64) string {
	randStr := fmt.Sprint(rand.Intn(1000000) + 100000) // 再加上這個以保證絕對不會重複
	data := []byte(fmt.Sprintf("%d%s", nowNano, randStr))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GetProjectRoot() string {
	// get the root path of project
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d) // this mechine is "/home/zhilong/Documents/my_projects/graph_robot"
}

func RemoveUniqueValueFromSlice[T constraints.Integer | constraints.Float | string](value T, s *[]T) {
	for i, v := range *s {
		if v == value {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
}

func SaveDataToFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
