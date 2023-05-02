package main

import (
	"fmt"
	"math/rand"
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
	// 	log.Panic(err)
	// }
	// fmt.Println(string(bytedata))
	// timer2 := timer{}
	// err = json.Unmarshal(bytedata, &timer2)
	// if err != nil {
	// 	log.Println("unmarshal error:" + err.Error())
	// 	log.Panic(err)
	// }
	// fmt.Println("timer2:", timer2.Abc)
	// timeString := "2023-04-22T18:10:06.94926253+09:00"
	// timeObj, err := time.Parse(time.RFC3339Nano, timeString)
	// if err != nil {
	// 	log.Println("parse error:" + err.Error())
	// 	log.Panic(err)
	// }
	// fmt.Println(timeObj)
	// var a time.Time
	// fmt.Println(a)
	// timeObj, err = time.Parse(time.RFC3339Nano, "0001-01-01 00:00:00 +0000 UTC")
	// if err != nil {
	// 	log.Println("parse error2:" + err.Error())
	// 	log.Panic(err)
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
	// var m sync.Map
	// m.Store("s1", S1[CCC]{})
	// s1, _ := m.Load("s1")
	// // s2, ok := s1.(S1[AAA])
	// // fmt.Println("OK:", ok)
	// // s2.PrintBBB()
	// switch ss := s1.(type) {
	// case S1[AAA]:
	// 	fmt.Println("S1AAA", s1)
	// case S1[CCC]:
	// 	fmt.Println("Type CCC")
	// 	ss.A.BBB()
	// case int:
	// 	fmt.Println("int")
	// }

	// test sync.map
	// m := sync.Map{}
	// ccc := CCC{}
	// fmt.Printf("first pointer:%p\n", &ccc)
	// m.Store("ccc", &ccc) // 保存指针无法保证并发安全
	// cccv, _ := m.Load("ccc")
	// cccp := cccv.(*CCC)
	// cccp.BBB()
	// var wg sync.WaitGroup
	// wg.Add(3)
	// AddC1 := func(ccc *CCC) {
	// 	for i := 0; i < 100000; i++ {
	// 		ccc.AddC1()
	// 	}
	// 	wg.Done()
	// }
	// go AddC1(cccp)
	// go AddC1(cccp)
	// go AddC1(cccp)
	// wg.Wait()
	// fmt.Println("c1 value:!!!!", cccp.C1)

	// cccv1, _ := m.Load("ccc")
	// cccp1 := cccv1.(*CCC)
	// fmt.Println("cccp1 value:&&&&", cccp1.C1)
	// fmt.Printf("cccp1 pointer: %p\n", cccp1)
	// cccv2, _ := m.Load("ccc")
	// cccp2 := cccv2.(*CCC)
	// fmt.Printf("cccp2 pointer: %p\n", cccp2)
	// fmt.Println("cccp2 ccc1:", cccp2.C1)

	// test main routine panic
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()
	// go func() {
	// 	// even main panic to recover, go routine still running
	// 	for i := 0; i < 20; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println("i:", i)
	// 	}
	// }()
	// log.Panic("main panic")

	// test wait group in routine
	// wg := &sync.WaitGroup{}
	// wg.Add(1)
	// go R1(wg)
	// wg.Wait()

	// tset sync map store channel
	// sm := sync.Map{}
	// sm.Store("c1", make(chan bool, 1))
	// go testChan(&sm)
	// go endC1(&sm) // according to this test, channel work correctly
	// time.Sleep(7 * time.Second)

	// test map array
	// aaa := make(map[string][]string)
	// fmt.Println("aaa", len(aaa["111"]))

	// test write map[string]interface to file
	// aaa := make(map[string]interface{})
	// var bbb []map[string]interface{}
	// aaa["111"] = 33
	// aaa["222"] = "fdsfsdf"
	// m := make(map[string]interface{})
	// m["000"] = 111
	// bbb = append(bbb, m)
	// aaa["333"] = bbb
	// prefix := []byte("var neures = ")
	// jsonData, err := json.Marshal(aaa)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// jsonFile, err := os.Create("./test_create.js")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer jsonFile.Close()
	// prefix = append(prefix, jsonData...)
	// _, err = jsonFile.Write(prefix)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// test go routine in go routine
	// go routineOuter() //测试结果是，只有main退出，inner才会中断，outer退出并不会中断inner
	// time.Sleep(6 * time.Second)
	// fmt.Println("main return")

	// test of empty map length
	// m := make(map[string]interface{})
	// // m["111"] = 3
	// fmt.Println("len:", len(m))
	// mb, err := json.Marshal(m)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("len byte:", len(mb))

	// // contain empty
	// a := "dfdfdf"
	// fmt.Println(strings.Contains(a, "")) // true

	// // test null type
	// b := make(map[string]interface{})
	// c := b["22"].(string)  // will panic
	// fmt.Println("c:", c)

	// remove map element:
	ccc1 := CCC{C1: 1}
	ccc2 := CCC{C1: 2}
	aaa := make(map[string]*CCC)
	aaa["222"] = &ccc1
	aaa["333"] = &ccc2
	aaa["444"] = &ccc1
	delete(aaa, "333")
	fmt.Printf("ccc1: %+v\n", ccc1)
	fmt.Printf("ccc2: %+v\n", ccc2)
	fmt.Println("aaa: ", aaa)

	bbb := []map[string]*CCC{
		aaa,
	}
	delete(bbb[0], "444")
	fmt.Println("aaa2: ", aaa)
}

func routineOuter() {
	go routineInner()
	time.Sleep(3 * time.Second)
	fmt.Println("outer return")
}

func routineInner() {
	time.Sleep(5 * time.Second)
	fmt.Println("inter return")
}

func endC1(sm *sync.Map) {
	c, _ := sm.Load("c1")
	c1 := c.(chan bool)
	time.Sleep(5 * time.Second)
	fmt.Println("true to c1")
	c1 <- true
}

func testChan(sm *sync.Map) {
	c, _ := sm.Load("c1")
	c1 := c.(chan bool)
	<-c1
	fmt.Println("get c1")
}

func R1(wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		//实验表明，即使R1运行完了，这个go routine仍然在运行
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("in R1 go")
	}(wg)
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
	mu sync.Mutex
	C1 int
}

func (c *CCC) AddC1() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.C1++
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

func get_rand_num() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10), " in go routine")
	}
}

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
