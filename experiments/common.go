package main

import "fmt"

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
	m := returnMapValue()
	fmt.Printf("outside function: %p\n", m)
	v := (*m)["111"]
	fmt.Printf("outside function v: %p\n", v)
}

func returnMapValue() *map[string]*string {
	m := make(map[string]*string)
	v := "222"
	m["111"] = &v
	fmt.Printf("v point is: %p\n", &v)
	v = "333"
	fmt.Println("m[111] now is: ", m["111"])
	fmt.Println("m[111] point is: ", *m["111"])
	fmt.Printf("map in function:%p\n", &m)
	fmt.Printf("value in function: %p\n", &v)
	return &m
}

type AAA struct {
	Content string
}

func DeletePointer(aaa *AAA) {
	aaa = nil
}

func TsetMap() map[string]int {
	aaa := make(map[string]int)
	aaa["111"] = 111
	fmt.Printf("inside:%p", &aaa)
	return aaa
}

type BBB struct {
	Content string
}

func (b *BBB) Say() {
	fmt.Println("BBB")
}

type CCC struct {
	BBB
	Foo string
}

func NeedBBB(bbb BBB) {
	bbb.Say()
}

type SliceTest struct {
	S1 []string
}

func RemoveUniqueValueFromSlice(value string, slice *[]string) {
	for i, v := range *slice {
		if v == value {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			break
		}
	}
}

type Synapse interface {
	GetNextId() string
}

type Synapse1 struct {
	NextId string
}

func (s Synapse1) GetNextId() string {
	return s.NextId
}

func RemoveUniqueValueFromSynapse[T Synapse](value string, s *[]T) {
	for i, v := range *s {
		if v.GetNextId() == value {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
}
