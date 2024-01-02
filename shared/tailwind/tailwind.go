package tw

import (
	"strings"

	"github.com/iancoleman/orderedmap"
)

type (
	TwString struct {
		Input string
	}
	TwObject struct {
		Input map[string]bool
	}
	TwSlice struct {
		Input []string
	}
	TwClasses struct {
		classes *orderedmap.OrderedMap
	}
	MergeArg interface {
		String() string
		parse(classes *orderedmap.OrderedMap)
	}
)

func NewTwClasses() *TwClasses {
	return &TwClasses{classes: orderedmap.New()}
}

func NewString(input string) *TwString {
	return &TwString{Input: input}
}

func NewObject(input map[string]bool) *TwObject {
	return &TwObject{Input: input}
}

func NewSlice(input []string) *TwSlice {
	return &TwSlice{Input: input}
}

func (tc *TwClasses) AddClasses(arg MergeArg) {
	arg.parse(tc.classes)
}

func (tc *TwClasses) String() string {
	result := strings.Builder{}

	for _, v := range tc.classes.Keys() {
		result.WriteString(v)
		result.WriteString(" ")
	}

	if result.Len() > 0 {
		return result.String()[0 : result.Len()-1]
	}

	return ""
}

func (s TwString) parse(classes *orderedmap.OrderedMap) {
	slice := strings.Split(s.Input, " ")

	for _, v := range slice {
		(*classes).Set(v, true)
	}
}

func (s TwString) String() string {
	return s.Input
}

func (o TwObject) parse(classes *orderedmap.OrderedMap) {
	for k, v := range o.Input {
		if v {
			(*classes).Set(k, true)
		}
	}
}

func (o TwObject) String() string {
	result := strings.Builder{}

	for k, v := range o.Input {
		if v {
			result.WriteString(k)
			result.WriteString(" ")
		}
	}

	if result.Len() > 0 {
		return result.String()[0 : result.Len()-1]
	}
	return ""
}

func (s TwSlice) parse(classes *orderedmap.OrderedMap) {
	for _, v := range s.Input {
		(*classes).Set(v, true)
	}
}

func (s TwSlice) String() string {
	result := strings.Builder{}

	for _, v := range s.Input {
		result.WriteString(v)
		result.WriteString(" ")
	}

	if result.Len() > 0 {
		return result.String()[0 : result.Len()-1]
	}
	return ""
}

func Merge(args ...MergeArg) string {
	tc := NewTwClasses()

	if len(args) == 0 {
		return ""
	}

	for _, arg := range args {
		tc.AddClasses(arg)
	}

	result := tc.String()

	return result
}
