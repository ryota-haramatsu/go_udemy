package main

/*
omitemptyで受け取ったjsonの値が空値であれば表示しない
MarshalJSON()
UnmarshalJSON()
でマーシャルやアンマーシャルをするときに、
修正をすることができる
*/

import (
	"encoding/json"
	"fmt"
)

type T struct{}

type Person struct {
	// omitempty : 空を除外する
	Name      string   `json:"name,omitempty"` // jsonでマーシャルするときの名前を設定できる
	Age       int      `json:"age,omitempty"`  // マーシャルするときに文字列型にすることもできる
	Nicknames []string `json:"nicknames,omitempty"`
	T         *T       `json: "T,omitempty"` // Tのポインタ型にしないと
}

// func (p Person) MarshalJSON() ([]byte, error) {
// 	// a := struct{Name string}{Name: "test"}
// 	v, err := json.Marshal(&struct {
// 		Name string
// 	}{
// 		Name: "Mr." + p.Name,
// 	})
// 	return v, err
// }

// func (p *Person) UnmarshalJSON(b []byte) error {
// 	// 独自のPerson2でアンマーシャルし、
// 	// 成功したらPersonのNameに独自のPerson2のName + "!!"を入れる
// 	type Person2 struct {
// 		Name string
// 	}
// 	var p2 Person2
// 	err := json.Unmarshal(b, &p2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	p.Name = p2.Name + "!!"
// 	return err
// }

func main() {
	// jsonのUnmarshal json→structへ
	b := []byte(`{"name": "ryota", "age": 10}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// jsonへMarshal
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
