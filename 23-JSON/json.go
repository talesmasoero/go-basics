package main

import (
	"encoding/json"
	"fmt"
)

// Adicionamos essas "tags" para o JSON ler/identificar corretamente
// os campos da struct
type Person struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	// Marshal: transforma algum tipo de dado, geralmente
	// structs ou maps, em JSON
	fmt.Println("MARSHAL:")
	ExemploJSONMarshal()

	// Unmarshal: transforma JSON em algum tipo de dado,
	// geralmente structs ou maps
	fmt.Println("UNMARSHAL:")
	ExemploJSONUnmarshal()
}

func ExemploJSONMarshal() {
	person1 := Person{"Tales", 21}
	person2 := map[string]any{"name": "Breno", "age": 15}

	// Usei MarshalIndent para melhor visualização no terminal
	person1JSON, err := json.MarshalIndent(person1, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pessoa 1: []Bytes (retorno): %v\n", person1JSON)
	fmt.Printf("Pessoa 1: String (em JSON): %v\n", string(person1JSON))

	person2JSON, err := json.MarshalIndent(person2, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pessoa 2: []Bytes (retorno): %v\n", person2JSON)
	fmt.Printf("Pessoa 2: String (em JSON): %v\n\n", string(person2JSON))
}

func ExemploJSONUnmarshal() {
	person1JSON := []byte(`{"name":"Tales","age":21}`)
	person2JSON := []byte(`{"name":"Breno","age":15}`)

	p1 := Person{}
	err := json.Unmarshal(person1JSON, &p1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pessoa 1: Struct: %v\n", p1)

	p2 := make(map[string]any)
	err = json.Unmarshal(person2JSON, &p2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pessoa 2: Map: %v\n", p2)
}
