package help

import (
	"math/rand"
	"time"
	"strconv"
)

type Person struct{
	name string
	age int
}

func PersonInstance() Person {
	rand.Seed(time.Now().UnixNano())
	age := rand.Intn(1000)
	return Person{name: "person" + strconv.Itoa(age) , age: age}
}

func PersonSlice(n int) PersonCompareContainer {
	res := make([]Person, n)
	for i := 0; i < n; i++ {
		res[i] = PersonInstance()
	}
	return res
}


type PersonCompareContainer []Person

func (p *PersonCompareContainer) More(i int, j int) bool {
	return (*p)[i].age > (*p)[j].age
}

func (p *PersonCompareContainer) Len() int {
	return len(*p)
}

func (p *PersonCompareContainer) Swap(i int, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PersonCompareContainer) Less(i int, j int) bool {
	return (*p)[i].age < (*p)[j].age
}









