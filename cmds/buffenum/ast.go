package main

import "github.com/alecthomas/participle"

type enumAst struct {
	Package string        `parser:"'package' @Ident"`
	Name    string        `parser:"'enum' @Ident"`
	Items   []enumItemAst `parser:"'{' @@* '}'"`
}

func (e enumAst) toModel() *Enum {
	items := make([]EnumItem, 0)
	for _, i := range e.Items {
		items = append(items, i.toModel())
	}

	return &Enum{Package: e.Package, Name: e.Name, Items: items}
}

type enumItemAst struct {
	Name     string `parser:"@Ident"`
	IntValue int    `parser:"'(' @Int ')'"`
	StrValue string `parser:"'=' @String ';'"`
}

func (e enumItemAst) toModel() EnumItem{
	return EnumItem{Name: e.Name, IntValue: e.IntValue, StringValue: e.StrValue}
}

var parser = participle.MustBuild(&enumAst{})
