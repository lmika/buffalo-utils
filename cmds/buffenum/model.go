package main

import "io"

type EnumItem struct {
	Name        string
	IntValue    int
	StringValue string
	IsDefault bool
	SelectLabel string
}

type Enum struct {
	Name    string
	Items   []EnumItem
}

type File struct {
	Package string
	Enums	[]*Enum
}

func ParseFile(r io.Reader) (*File, error) {
	var e fileAst
	if err := parser.Parse(r, &e); err != nil {
		return nil, err
	}
	return e.toModel(), nil
}

func (en *File) Generate(w io.Writer) error {
	return newCodeGen(en).generate(w)
}


func (en *Enum) DefaultItem() (EnumItem, bool) {
	for _, i := range en.Items {
		if i.IsDefault {
			return i, true
		}
	}
	return EnumItem{}, false
}