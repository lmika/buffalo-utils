package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"io"
)

type codeGen struct {
	enum *Enum

	unknownStringValue string

	typeName *jen.Statement
}

func newCodeGen(enum *Enum) *codeGen {
	return &codeGen{
		enum:               enum,
		unknownStringValue: "unknown",
	}
}

func (cg *codeGen) generate(w io.Writer) error {
	cg.typeName = jen.Id(cg.enum.Name)

	f := jen.NewFile(cg.enum.Package)

	f.Type().Add(cg.typeName).Int()
	f.Const().DefsFunc(func(g *jen.Group) {
		for i, e := range cg.enum.Items {
			stmt := g.Id(cg.enumItemName(e))
			if i == 0 {
				stmt = stmt.Id(cg.enum.Name)
			}
			stmt.Op("=").Lit(e.IntValue)
		}
	})

	// Array of valid values
	f.Var().Id(cg.enum.Name + "Values").Op("=").Index().Id(cg.enum.Name).ValuesFunc(func(g *jen.Group) {
		for _, e := range cg.enum.Items {
			g.Id(cg.enumItemName(e))
		}
	})

	cg.generateStringMethod(f)
	cg.generateScanMethod(f)
	cg.generateUnmarshalTextMethod(f)
	cg.generateSelectValueMethod(f)
	cg.generateSelectLabelMethod(f)

	return f.Render(w)
}

func (cg *codeGen) generateStringMethod(f *jen.File) {
	receiver := jen.Id("e").Add(cg.typeName)
	f.Func().Params(receiver).Id("String").Params().String().BlockFunc(func(g *jen.Group) {
		g.Switch(jen.Id("e")).BlockFunc(func(g *jen.Group) {
			for _, e := range cg.enum.Items {
				g.Case(jen.Id(cg.enumItemName(e))).Block(
					jen.Return(jen.Lit(e.StringValue)),
				)
			}
		})
		g.Return(jen.Lit(cg.unknownStringValue))
	})
}

func (cg *codeGen) generateScanMethod(f *jen.File) {
	receiver := jen.Id("e").Op("*").Add(cg.typeName)

	f.Func().Params(receiver).Id("Scan").Params(
		jen.Id("src").Interface(),
	).Error().BlockFunc(func(g *jen.Group) {
		g.Id("i").Op(",").Id("isInt").Op(":=").Id("src").Op(".").Params(jen.Id("int64"))

		g.If(jen.Op("!").Id("isInt")).BlockFunc(func(g *jen.Group) {
			g.Return(jen.Qual("errors", "New").Params(
				jen.Lit(fmt.Sprintf("%v is not an int", cg.enum.Name))),
			)
		})

		g.Op("*").Id("e").Op("=").Id(cg.enum.Name).Params(jen.Id("i"))
		g.Return(jen.Nil())
	})
}

func (cg *codeGen) generateUnmarshalTextMethod(f *jen.File) {
	receiver := jen.Id("e").Op("*").Add(cg.typeName)
	textParam := jen.Id("text")

	f.Func().Params(receiver).Id("UnmarshalText").Params(
		textParam.Index().Byte(),
	).Error().BlockFunc(func(g *jen.Group) {
		// TODO: allow case insensitive values
		textAsStr := jen.String().Params(jen.Id("text"))

		g.Switch(textAsStr).BlockFunc(func(g *jen.Group) {
			for _, e := range cg.enum.Items {
				g.Case(jen.Lit(e.StringValue)).BlockFunc(func(g *jen.Group) {
					g.Op("*").Id("e").Op("=").Id(cg.enumItemName(e))
					g.Return(jen.Nil())
				})
			}
		})
		g.Return(jen.Qual("errors", "New").Params(
			jen.Lit(fmt.Sprintf("invalid value for %v", cg.enum.Name))),
		)
	})
}

func (cg *codeGen) generateSelectValueMethod(f *jen.File) {
	receiver := jen.Id("e").Add(cg.typeName)
	f.Func().Params(receiver).Id("SelectValue").Params().Interface().BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id("e").Dot("String").Call())
	})
}

func (cg *codeGen) generateSelectLabelMethod(f *jen.File) {
	receiver := jen.Id("e").Add(cg.typeName)
	f.Func().Params(receiver).Id("SelectLabel").Params().String().BlockFunc(func(g *jen.Group) {
		g.Switch(jen.Id("e")).BlockFunc(func(g *jen.Group) {
			for _, e := range cg.enum.Items {
				g.Case(jen.Id(cg.enumItemName(e))).Block(
					jen.Return(jen.Lit(e.Name)),
				)
			}
		})
		g.Return(jen.Lit(cg.unknownStringValue))
	})
}

func (cg *codeGen) enumItemName(enumItem EnumItem) string {
	// TODO: configure
	return enumItem.Name + cg.enum.Name
}
