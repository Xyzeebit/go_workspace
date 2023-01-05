package main

import (
	"container/list"
	"fmt"
	// "sort"
)

func main() {
	var x list.List;
	x.PushBack(1);
	x.PushBack(2);
	x.PushBack(3);

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int));
	}
	// fmt.Println()

	// kids := []Person {
	// 	{ "Jill", 9 },
	// 	{ "Jack", 10 },
	// };
	// sort.Sort(ByName(kids));
	// fmt.Println(kids);
}