package main

import (
	"errors"
	"fmt"
)

func main() {

	m1 := make(MyMap) // can use build in function called make on user defined type MyMap
	m1["522001"] = "Guntur-1"
	m1["522002"] = "Guntur-2"
	m1["560086"] = "RajajiNagar Bangalore"
	m1["560096"] = "Yeshvantpur Bangalore"

	m1.Display()
	keys, values := m1.GetKeyValueSlices()
	fmt.Println("keys:", keys)
	fmt.Println("values:", values)

	if err := m1.Delete("560097"); err != nil {
		println(err.Error())
	} else {
		println("key successfully deleted from map")
	}

	m2 := make(map[string]any) // not custom type but a predefined type
	m2["522001"] = "Guntur-1"
	m2["522002"] = "Guntur-2"
	m2["560086"] = "RajajiNagar Bangalore"
	m2["560096"] = "Yeshvantpur Bangalore"
	println("\ncalling a MyMap method on regular map by type casting it")
	MyMap(m2).Display()
}

type MyMap map[string]any

func (m1 MyMap) GetKeyValueSlices() (keys []string, values []any) {
	for key, value := range m1 {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}

func (m1 MyMap) Delete(key string) error {
	if m1 == nil {
		return errors.New("called on nil map")
	}
	if _, ok := m1[key]; !ok {
		return fmt.Errorf("key:%s does not exist", key)
	}

	delete(m1, key)
	return nil

}

func (m1 MyMap) Display() {
	for key, value := range m1 {
		fmt.Println("Key:", key, "Value:", value)
	}
}
