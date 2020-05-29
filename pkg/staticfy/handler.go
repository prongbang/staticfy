package staticfy

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type Handler interface {
	Register()
}

type handler struct {
}

func (h *handler) Register() {
	banner := "\n" +
		"   ______       __  _     ___    \n" +
		"  / __/ /____ _/ /_(_)___/ _/_ __\n" +
		" _\\ \\/ __/ _ `/ __/ / __/ _/ // /\n" +
		"/___/\\__/\\_,_/\\__/_/\\__/_/ \\_, / \n" +
		"                          /___/\n"

	patternDir := `
	Require pattern:
		staticfy
    	├── config.yml
    	└── statics
        	└── filename.txt`

	source, err := ioutil.ReadFile(StaticfyConfigYmlPath)
	if err != nil {
		panic(patternDir)
	}

	data := Config{}
	err = yaml.Unmarshal(source, &data)
	if err != nil {
		panic(patternDir)
	}

	pattern := fmt.Sprintf("%s/", data.Staticfy.Prefix)
	fileServer := http.FileServer(FileSystem{http.Dir(fmt.Sprintf(StaticfyStaticsPath, data.Staticfy.Directory))})
	http.Handle(pattern, http.StripPrefix(strings.TrimRight(pattern, "/"), fileServer))

	fmt.Println(banner)
	fmt.Printf(fmt.Sprintf("-> staticfy serving %s on :%d\n", data.Staticfy.Prefix, data.Staticfy.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", data.Staticfy.Port), nil))
}

func NewHandler() Handler {
	return &handler{}
}
