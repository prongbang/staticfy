package staticfy

type Config struct {
	Staticfy struct {
		Directory string `yaml:"directory"`
		Prefix    string `yaml:"prefix"`
		Port      int    `yaml:"port"`
	} `yaml:"staticfy"`
}

const (
	StaticfyStaticsPath   = "./staticfy%s"
	StaticfyConfigYmlPath = "./staticfy/config.yml"
)
