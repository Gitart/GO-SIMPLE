package mypkg

import (
	"fmt"
	"io"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)


func MarshalMyCustomBooleanScalar(b bool) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if b {
			w.Write([]byte("true"))
		} else {
			w.Write([]byte("false"))
		}
	})
}

func UnmarshalMyCustomBooleanScalar(v interface{}) (bool, error) {
	switch v := v.(type) {
	case string:
		return "true" == strings.ToLower(v), nil
	case int:
		return v != 0, nil
	case bool:
		return v, nil
	default:
		return false, fmt.Errorf("%T is not a bool", v)
	}
}
