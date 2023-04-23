package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// random number
	// rand.Seed(time.Now().UnixNano())
	// aaa := rand.Intn(30)
	// fmt.Println(aaa)

	// test delete pointer
	// aaa := AAA{"fdsfsdf"}
	// bbb := &aaa
	// DeletePointer(bbb)
	// fmt.Println(bbb == nil)
	// fmt.Println(aaa)

	// now := time.Now().UnixNano()
	// data := []byte(fmt.Sprint(now))
	// aa := fmt.Sprintf("%x", md5.Sum(data))
	// fmt.Println(aa)

	// // test map
	// aaa := TsetMap()
	// // b := aaa["111"]
	// fmt.Printf("outside:%p", &aaa)

	// ccc := CCC{Foo: "foo"}
	// ccc.Say()
	// fmt.Println(ccc.Foo)

	// s := SliceTest{
	// 	S1: []string{"11", "22", "33", "44"},
	// }
	// RemoveUniqueValueFromSlice("22", &s.S1)
	// fmt.Println("s1:", s.S1)

	// s1 := Synapse1{
	// 	NextId: "11",
	// }
	// s2 := Synapse1{
	// 	NextId: "22",
	// }
	// s3 := Synapse1{
	// 	NextId: "44",
	// }
	// s4 := Synapse1{
	// 	NextId: "55",
	// }
	// sslice := []Synapse1{s1, s2, s3, s4}
	// RemoveUniqueValueFromSynapse("44", &sslice)
	// fmt.Printf("!!!%+v", sslice)

	// test byte append
	// var result [][]byte
	// result = append(result, []byte{116, 101, 115, 116, 105, 110, 103, 95, 110, 101, 117, 114, 101, 64, 48})
	// result = append(result, []byte{116, 101, 115, 116, 105, 110, 103, 95, 110, 101, 117, 114, 101, 64, 50})
	// result = append(result, []byte{116, 101, 115, 116, 105, 110, 103, 95, 110, 101, 117, 114, 101, 64, 51})
	// fmt.Println("resul:", result)

	// test map return copy or pointer
	// m := returnMapValue()
	// fmt.Printf("outside function: %p\n", m)
	// v := (*m)["111"]
	// fmt.Printf("outside function v: %p\n", v)

	// test switch
	// a := 3
	// switch {
	// case a < 2:
	// 	fmt.Println("a < 2")
	// case a > 4:
	// 	fmt.Println("a > 4")
	// case a == 3:
	// 	fmt.Println("a == 3")
	// }

	// test string array copy
	// stringArray := []string{"11"}
	// testStringArrayCopy(stringArray)
	// fmt.Println("result:", stringArray[0])  // 22

	// seed test
	// rand.Seed(time.Now().UnixNano())
	// biggerthan := []bool{}
	// for i := 0; i < 100000; i++ {
	// 	if rand.Float32() > 0.5 {
	// 		biggerthan = append(biggerthan, true)
	// 	}
	// }
	// fmt.Println(len(biggerthan))
	// fmt.Println(float32(len(biggerthan)) / 100000)

	// rand.Seed(2)
	// go get_rand_num()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Intn(10), " in main routine")
	// 	time.Sleep(1 * time.Microsecond)
	// }
	// time.Sleep(1 * time.Second)

	// time test
	// now := time.Now()
	// t2 := 2 * time.Microsecond
	// timer1 := timer{
	// 	Abc: now.Add(t2),
	// }
	// bytedata, err := json.Marshal(timer1)
	// if err != nil {
	// 	log.Println("marshal error:" + err.Error())
	// 	panic(err)
	// }
	// fmt.Println(string(bytedata))
	// timer2 := timer{}
	// err = json.Unmarshal(bytedata, &timer2)
	// if err != nil {
	// 	log.Println("unmarshal error:" + err.Error())
	// 	panic(err)
	// }
	// fmt.Println("timer2:", timer2.Abc)
	// timeString := "2023-04-22T18:10:06.94926253+09:00"
	// timeObj, err := time.Parse(time.RFC3339Nano, timeString)
	// if err != nil {
	// 	log.Println("parse error:" + err.Error())
	// 	panic(err)
	// }
	// fmt.Println(timeObj)
	// var a time.Time
	// fmt.Println(a)
	// timeObj, err = time.Parse(time.RFC3339Nano, "0001-01-01 00:00:00 +0000 UTC")
	// if err != nil {
	// 	log.Println("parse error2:" + err.Error())
	// 	panic(err)
	// }

	// sync.Map test
	// var sm sync.Map
	// timer33 := timer{}
	// sm.Store("11", &timer{})
	// sm.Store("33", &timer33)
	// sm.Store("22", &timer{Abc: time.Now()})
	// _, ok := sm.Load("99")
	// sm.LoadOrStore("22", &timer{})
	// fmt.Println("length:", ok)
	// t3, _ := sm.Load("33")
	// t3.(*timer).Abc = time.Now()
	// sm.Range(func(key, value any) bool {
	// 	fmt.Printf("key:%s, value:%+v\n", key, value)
	// 	fmt.Println("value:", value.(*timer).Abc)
	// 	if key == "11" {
	// 		sm.Delete(key) // can delete here
	// 	}
	// 	return true
	// })
	// _, ok = sm.Load("11")
	// fmt.Println("11 not deleted:", ok)

	// test type
	// testType[*timer]()

	// test Type struct
	var m sync.Map
	m.Store("s1", S1[CCC]{})
	s1, _ := m.Load("s1")
	// s2, ok := s1.(S1[AAA])
	// fmt.Println("OK:", ok)
	// s2.PrintBBB()
	switch ss := s1.(type) {
	case S1[AAA]:
		fmt.Println("S1AAA", s1)
	case S1[CCC]:
		fmt.Println("Type CCC")
		ss.A.BBB()
	case int:
		fmt.Println("int")
	}
}

type AAA interface {
	BBB()
}

type S1[T AAA] struct {
	A T
}

func (s S1[T]) PrintBBB() {
	s.A.BBB()
}

type CCC struct {
	C1 string
}

func (c CCC) BBB() {
	fmt.Println("CCC")
}

type DDD struct{}

func (c DDD) BBB() {

}

func testType[T *timer | float32]() {
	fmt.Printf("type:%T", *new(T))
}

type timer struct {
	Abc time.Time `json:"111"`
}

// func get_rand_num() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(rand.Intn(10), " in go routine")
// 	}
// }

// func testStringArrayCopy(stringArray []string) {
// 	stringArray[0] = "22"
// }

// func returnMapValue() *map[string]*string {
// 	m := make(map[string]*string)
// 	v := "222"
// 	m["111"] = &v
// 	fmt.Printf("v point is: %p\n", &v)
// 	v = "333"
// 	fmt.Println("m[111] now is: ", m["111"])
// 	fmt.Println("m[111] point is: ", *m["111"])
// 	fmt.Printf("map in function:%p\n", &m)
// 	fmt.Printf("value in function: %p\n", &v)
// 	return &m
// }

// type AAA struct {
// 	Content string
// }

// func DeletePointer(aaa *AAA) {
// 	aaa = nil
// }

// func TsetMap() map[string]int {
// 	aaa := make(map[string]int)
// 	aaa["111"] = 111
// 	fmt.Printf("inside:%p", &aaa)
// 	return aaa
// }

// type BBB struct {
// 	Content string
// }

// func (b *BBB) Say() {
// 	fmt.Println("BBB")
// }

// type CCC struct {
// 	BBB
// 	Foo string
// }

// func NeedBBB(bbb BBB) {
// 	bbb.Say()
// }

// type SliceTest struct {
// 	S1 []string
// }

// func RemoveUniqueValueFromSlice(value string, slice *[]string) {
// 	for i, v := range *slice {
// 		if v == value {
// 			*slice = append((*slice)[:i], (*slice)[i+1:]...)
// 			break
// 		}
// 	}
// }

// type Synapse interface {
// 	GetNextId() string
// }

// type Synapse1 struct {
// 	NextId string
// }

// func (s Synapse1) GetNextId() string {
// 	return s.NextId
// }

// func RemoveUniqueValueFromSynapse[T Synapse](value string, s *[]T) {
// 	for i, v := range *s {
// 		if v.GetNextId() == value {
// 			*s = append((*s)[:i], (*s)[i+1:]...)
// 			break
// 		}
// 	}
// }
