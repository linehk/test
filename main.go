package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	var p Payload
	_ = json.Unmarshal([]byte(`{
            "field1":"somevalue1",
 	    "field2":null
        }`), &p)

	fmt.Printf(`
Field1 is defined and has a value. %+v
Field2 is defined but null. %+v
Field3 is undefined. %+v
	`, p.Field1, p.Field2, p.Field3)
}

type Optional[T any] struct {
	Defined bool
	Value   *T
}

// UnmarshalJSON is implemented by deferring to the wrapped type (T).
// It will be called only if the value is defined in the JSON payload.
func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	o.Defined = true
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	return dec.Decode(&o.Value)
	// return json.Unmarshal(data, &o.Value)
}

type Payload struct {
	Field1 Optional[any] `json:"field1"`
	Field2 Optional[any] `json:"field2"`
	Field3 Optional[any] `json:"field3"`
}
