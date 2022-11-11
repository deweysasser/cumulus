package code_generation

import (
	"fmt"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/fatih/camelcase"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"strings"
)

var Generate = []interface{}{
	ec2.Instance{},
	ec2.Snapshot{},
	route53.HostedZone{},
	route53.ResourceRecordSet{},
	ec2.Volume{},
}

func GenerateAllYaml() {
	for _, g := range Generate {
		generateYaml(g)
	}
}

func generateYaml(t interface{}) {
	typ := reflect.TypeOf(t)

	typeToken := strings.Replace(typ.String(), ".", "_", -1)

	f, err := os.Create(typeToken + ".yaml")

	if err != nil {
		log.Error().
			Err(err).
			Str("name", typ.String()).
			Msg("error generating type")
		return
	}

	defer f.Close()

	yt := Type{Type: typ.Name()}

	log.Debug().
		Str("str", typ.String()).
		Str("type", typ.Name()).
		Str("package", typ.PkgPath()).
		Msg("Generating")

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		log.Debug().Str("field", f.Name).
			Str("type", f.Type.String()).
			Msg("field")

		fieldName := fieldName(f)

		skip := false
		category := "What"
		converter := ""
		fieldType := ""
		if f.Type.Kind() == reflect.Pointer {
			fieldType = f.Type.Elem().Name()
			converter = fmt.Sprintf("aws.%sValue", util.Capitalize(fieldType))
		}
		function := ""

		switch f.Name {
		case "Tags":
			skip = false
			function = fmt.Sprintf("%s_to_fields", strings.Replace(f.Type.Elem().Elem().String(), ".", "_", -1))
		case "Size":
			converter = "toSizeInG"
		default:
			switch f.Type.String() {
			case "*string":
			case "*time.Time":
				category = "When"
			case "*bool":
				converter = "boolToString"
			case "[]*string":
			case "*int64":
			default:
				skip = true
			}
		}

		yt.Fields = append(yt.Fields, Field{
			Name:      fieldName,
			Converter: converter,
			Function:  function,
			AWSName:   f.Name,
			Category:  category,
			Type:      fieldType,
			Skip:      skip,
		})
	}

	//fmt.Printf("}\n\n")

	bytes, err := yaml.Marshal(&yt)

	if err != nil {
		log.Error().
			Err(err).
			Msg("Error generating yaml")
	}

	f.Write(bytes)
}

func fieldName(f reflect.StructField) string {
	var r []string
	for _, s := range camelcase.Split(f.Name) {
		r = append(r, strings.ToLower(s))
	}

	return strings.Join(r, "_")
}
