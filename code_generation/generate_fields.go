package code_generation

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
)

func GenerateFieldCode(pkg string, file string) error {
	target := strings.TrimSuffix(file, filepath.Ext(file)) + ".go"

	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	typ := &Type{}

	err = yaml.Unmarshal(bytes, typ)
	if err != nil {
		return err
	}

	if len(typ.Fields) < 1 {
		return errors.New("No fields found")
	}

	out, err := os.Create(target)
	if err != nil {
		return err
	}

	defer out.Close()

	f := jen.NewFile(pkg)
	f.Comment(fmt.Sprintf("Code generated. DO NOT EDIT.  Code generated from %s", file))
	f.Line()

	f.Func().
		Params(
			jen.Id("i").Id(strings.ToLower(typ.Type)),
		).
		Id("GeneratedFields").
		Params(
			jen.Id("builder").Qual("github.com/deweysasser/cumulus/cumulus", "IFieldBuilder"),
		).
		BlockFunc(func(g *jen.Group) {
			for _, f := range typ.Fields {
				field := jen.Id("i").Dot("obj").Dot(f.AWSName)

				hidden := jen.Qual("github.com/deweysasser/cumulus/cumulus", "DefaultHidden")

				if f.ShowByDefault {
					hidden = nil
				}

				converter := jen.Id(f.Converter)
				if strings.HasPrefix(f.Converter, "aws.") {
					converter = jen.Qual("github.com/aws/aws-sdk-go/aws", f.Converter[4:])
				}

				converter = converter.Call(field)

				if f.Type != "string" && f.Type != "Time" {
					converter = jen.Qual("fmt", "Sprint").Call(converter)
				}

				switch {
				case f.Skip:
				case f.Function != "":
					g.Id(f.Function).Call(jen.Id("builder"), jen.Id("i").Dot("Ctx").Call(), field)
				case f.Category == "GID" || f.Category == "Name" || f.Category == "Description":
					g.Id("builder").Dot(f.Category).Call(
						converter,
					)
				default:
					g.If(field.Clone().Op("!=").Nil()).Block(
						jen.Id("builder").Dot(f.Category).Call(
							jen.Lit(f.Name),
							converter,
							hidden,
						),
					)
				}
				g.Line()
			}
		},
		)

	err = f.Render(out)
	if err != nil {
		return err
	}

	return nil
}
