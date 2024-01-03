package main

import (
	"log"
)

func main() {
	meep := map[string]int{
		"foo":  42,
		"fizz": 1337,
		"bar":  31337,
	}
	//lis := ["fizz", "buzz"]
Loop:
	for k, v := range meep {
		log.Printf("Using switch/case")
		switch {
		case 42 == v:
			log.Println("Found the Forty two!")
			log.Printf("Key is type %T with a value of %v, Value is type %T with value of %d \n", k, k, v, v)
		case 31337 == v || 1337 == v:
			log.Printf("You really are elite, aren't you? %d the type is: %T", v, v)
			if v == 1337 {
				log.Printf("ERROR ITS 1337")
				break Loop
			}
		}
		//log.Println("Falling out of switch statement")
	}
	//fmt.Println(meep)
	counter := 0 // set counter to 0
	var newMap = make(map[string]int)
	for k, v := range meep {
		log.Printf("if/then block: testing value %v loop increment %d", v, counter)
		if v == 42 {
			log.Println("42 in for/if loop")
			newMap[k] = v
		}
		if v == 31337 {
			log.Println("Hit that 31337")
			newMap[k] = v
		}
		//else {
		//log.Printf("Didn't find anything for %v", v)
		//}
		counter += 1
	}
	log.Printf("printing out updated map now")
	for k, v := range newMap {
		log.Printf("%T: %v %T: %v", k, k, v, v)
	}
}
