package main

func main() {
	map1 := map[string]string{"522001": "Guntur-1", "522002": "Guntur-2", "560086": "Bangalore-1"}
	key1, key2 := "522001", "522002"
	// keys, _ := GetKeysAndValues(map1)
	// slices.Contains(keys, key1)
	_, ok1 := map1[key1]
	_, ok2 := map1[key2]

	if ok1 && ok2 {
		println("both keys", key1, key2, "exists")
	}

	count := 0
	for key := range map1 {
		if key == key1 {
			count++
		}
		if key == key2 {
			count++
		}
	}

	if count == 2 {
		println("both keys", key1, key2, "exists")
	} else if count == 1 {
		println("one of the keys are available")
	}

	// num := 29

	// println(num % 8)
	// println(num & 7) // mod operation

	// num % 2 power 3 = num & (2 power 3-1)

}

func GetKeysAndValues(map1 map[string]string) (keys []string, values []string) {
	for key, value := range map1 {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}
