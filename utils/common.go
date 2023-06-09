package utils

import (
	"crypto/md5"
	"fmt"
	"log"
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
	return filepath.Dir(d) // "/xxx/yyy/graph_robot"
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
		log.Panic(err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		log.Panic(err)
	}
}

func GetMapKeys[T constraints.Integer | constraints.Float | string | struct{} | interface{} | bool](m map[string]T) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

func GetUnrepeatedRandNum(max int, needNumber int) (resultIndex []int) {
	// usage: slice1 := []float32{1.1, 2.2, 3.3}  resultIndex := GetUnrepeatedRandNum(len(slice1), 2)
	var resultSlice = []int{}
	for i := 0; i < max; i++ {
		resultSlice = append(resultSlice, i)
	}
	for j := 0; j < needNumber; j++ {
		index := rand.Intn(len(resultSlice))
		randNum := resultSlice[index] // get a rand element from resultSlice
		resultIndex = append(resultIndex, randNum)
		resultSlice = append(resultSlice[:index], resultSlice[index+1:]...)
	}
	return
}
