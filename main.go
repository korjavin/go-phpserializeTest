package main

import (
	"bufio"
	"fmt"
	"github.com/korjavin/go-php-serialize"
	"golang.org/x/text/encoding/unicode"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		val, _ := phpserialize.DecodeWithEncoding(text, unicode.UTF8)
		switch vtype := val.(type) {
		default:
			fmt.Printf("unexpected type %T", vtype)
		case map[interface{}]interface{}:
			map1 := val.(map[interface{}]interface{})
			for k, v := range map1 {
				if map2, ok := v.(map[interface{}]interface{}); ok {
					for k, v := range map2 {
						fmt.Printf("%v %v \n ", k, v)
					}

				} else {
					fmt.Printf("%v %v \n ", k, v)
				}
			}
		}
	}
}
func traverseMap(key interface{}, val interface{}) {
	switch val.(type) {
	default:
		fmt.Printf("%v %v \n", key, val)
	case map[interface{}]interface{}:
		map1 := val.(map[interface{}]interface{})
		for k, v := range map1 {
			traverseMap(k, v)
		}
	}
}
