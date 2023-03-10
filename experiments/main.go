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

	ccc := CCC{Foo: "foo"}
	ccc.Say()
	fmt.Println(ccc.Foo)
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
