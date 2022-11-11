package code_generation

type Type struct {
	Type   string  `yaml:"type"`
	Fields []Field `yaml:"fields"`
}

type Field struct {
	Name          string `yaml:"name"`
	Type          string `yaml:"type,omitempty"`
	AWSName       string `yaml:"awsname"`
	Category      string `yaml:"category"`
	Converter     string `yaml:"converter,omitempty"`
	Function      string `yaml:"function,omitempty"`
	Skip          bool   `yaml:"skip"`
	ShowByDefault bool   `yaml:"show_by_default"`
}
