package main

import "fmt"

func main() {
	mapIDName := make(map[int]map[string]string)

	mapName1 := make(map[string]string)
	mapName1["Primeiro"] = "Tawe"
	mapName1["Ultimo"] = "Macawe"
	mapIDName[0] = mapName1

	mapName2 := make(map[string]string)
	mapName2["Primeiro"] = "Hello"
	mapName2["Ultimo"] = "World"
	mapIDName[1] = mapName2

	fmt.Printf("Map ID:\t\tFirst Name:\tLast Name:\n")
	// k = key (chave): mapID
	// v = value (valor): mapName
	for k, v := range mapIDName {
		fmt.Printf("%d\t\t%s\t\t%s\n", k, v["Primeiro"], v["Ultimo"])
	}
}
