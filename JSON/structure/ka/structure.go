
// *****************************************************************************
// On line convertor
// https://app.quicktype.io/#l=go
// https://app.quicktype.io/?share=oaHrFJextwTffp0PBlR5
// https://app.quicktype.io?share=Xiunr9VRbyY7qh74W2gd
// https://gobyexample.com/json
// 
// / This file was generated from JSON Schema using quicktype, do not modify it directly.
// // To parse and unparse this JSON data, add this code to your project and do:
// //
// //    welcome, err := UnmarshalWelcome(bytes)
// //    bytes, err = welcome.Marshal()
// *****************************************************************************

// package main

// import "encoding/json"

// func UnmarshalWelcome(data []byte) (Welcome, error) {
// 	var r Welcome
// 	err := json.Unmarshal(data, &r)
// 	return r, err
// }

// func (r *Welcome) Marshal() ([]byte, error) {
// 	return json.Marshal(r)
// }


type Chk struct {
	F []F `json:"F,omitempty"`
}

type F struct {
	C *C `json:"C,omitempty"`
	N *N `json:"N,omitempty"`
	S *S `json:"S,omitempty"`
}

type C struct {
	CM *string `json:"cm,omitempty"`
}

type N struct {
	CM   *string `json:"cm,omitempty"`  
	Attr *string `json:"attr,omitempty"`
}

type S struct {
	Code  *int64   `json:"code,omitempty"` 
	Price *float64 `json:"price,omitempty"`
	Name  *string  `json:"name,omitempty"` 
	Qty   *int64   `json:"qty,omitempty"`  
}
