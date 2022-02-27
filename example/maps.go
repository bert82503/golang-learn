package main

import "fmt"

// Maps are Go’s built-in associative data type
// (sometimes called hashes or dicts in other languages).
/*
Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
(关联数据类型，哈希表/散列表或字典)

To create an empty map, use the builtin make: make(map[key-type]val-type).

Set key/value pairs using typical name[key] = val syntax.

Printing a map with e.g. fmt.Println will show all of its key/value pairs.

Get a value for a key with name[key].

The builtin len returns the number of key/value pairs when called on a map.

The builtin delete removes key/value pairs from a map.

The optional second return value when getting a value from a map indicates if the key was present in the map.
This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
Here we didn’t need the value itself, so we ignored it with the blank identifier _.

You can also declare and initialize a new map in the same line with this syntax.

Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
*/
func main() {
	// create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	// map: map[k1:7 k2:13]
	// map[k:v k:v]

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	// v1:  7

	fmt.Println("len:", len(m))
	// len: 2

	delete(m, "k2")
	fmt.Println("map:", m)
	// map: map[k1:7]

	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	// prs: false

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	// map: map[bar:2 foo:1]
}
