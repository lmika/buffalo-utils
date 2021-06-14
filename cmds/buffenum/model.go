package main

import "io"

type EnumItem struct {
	Name        string
	IntValue    int
	StringValue string
}

type Enum struct {
	Package string
	Name    string
	Items   []EnumItem
}

func ParseEnum(r io.Reader) (*Enum, error) {
	var e enumAst
	if err := parser.Parse(r, &e); err != nil {
		return nil, err
	}
	return e.toModel(), nil
}

func (en *Enum) Generate(w io.Writer) error {
	return newCodeGen(en).generate(w)
}