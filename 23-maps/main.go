package main

import "fmt"

func main() {
	var map1 map[string]any // declared the map so map is nil

	if map1 == nil {
		println("map1 is nil")
	}

	map1 = make(map[string]any) // instantiation of the map

	map1["560086"] = "Bangalore-1"
	map1["560096"] = "Bangalore-2"
	map1["560034"] = "Bangalore-3"
	map1["560044"] = "Bangalore-4"
	map1["100001"] = 0

	for key, value := range map1 { // on map range gives key and value
		fmt.Println("Key:", key, " -- Value:", value)
	}

	val := map1["560044"]
	fmt.Println(val)

	val, ok := map1["100001"] // how do I know whether the key existed or not
	if ok {
		fmt.Println(val)
	} else {
		println("key does not exist")
	}
	delete(map1, "100001")   // delete a key-pair from the map
	val, ok = map1["100001"] // how do I know whether the key existed or not
	if ok {
		fmt.Println(val)
	} else {
		println("key does not exist")
	}

	clear(map1) // clear deletes all entries from the map

	for key, value := range map1 { // on map range gives key and value
		fmt.Println("Key:", key, " -- Value:", value)
	}

}

// map is values are mapped to keys
// map is not ordered
// Loopup is constant times O(1)
// map is not thread safe
// make is used to create a map
// what should be the key? --? key can be any thing where == can be used upon

// when a map is created few buckets are created
// initially 2 buckets are created (if not capacity is given)
// each bucket contsins two parallel arrays of 8 elemtns

// Buckets --> Parallel Arrays(8 elements)
// Load Factor -> 6.5 -- double the buckets, redistribute all the data
// Collision -->
// Overflow -->
// HashKey

//0a4d55a8d778e5022fa b701977c5d840bbc486d0
//0a4d55a8d778e5022fab701977c5d840bbc486d0
