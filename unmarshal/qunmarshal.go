// Custom scalars with user defined types

package mypkg

import (
	"fmt"
	"io"
)

type YesNo bool

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (y *YesNo) UnmarshalGQL(v interface{}) error {
	yes, ok := v.(string)
	if !ok {
		return fmt.Errorf("YesNo must be a string")
	}

	if yes == "yes" {
		*y = true
	} else {
		*y = false
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (y YesNo) MarshalGQL(w io.Writer) {
	if y {
		w.Write([]byte(`"yes"`))
	} else {
		w.Write([]byte(`"no"`))
	}
}
