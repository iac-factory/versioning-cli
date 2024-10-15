package output

import (
	"errors"

	"github.com/spf13/pflag"
)

// Type string that implements Cobra's Type interface for valid string enumeration values.
type Type string

const (
	JSON Type = "json"
	YAML Type = "yaml"
)

// String is used both by fmt.Print and by Cobra in help text
func (o *Type) String() string {
	return string(*o)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (o *Type) Set(v string) error {
	switch v {
	case "json", "yaml":
		*o = Type(v)
		return nil
	default:
		return errors.New("must be one of \"json\" or \"yaml\"")
	}
}

// Type is only used in help text
func (o *Type) Type() string {
	return "[\"json\"|\"yaml\"]"
}

// --> interface runtime validation.
// - See https://stackoverflow.com/questions/27803654/explanation-of-checking-if-value-implements-interface for additional
// information.
var _ pflag.Value = (*Type)(nil)
