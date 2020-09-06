package staticfy

import (
	"fmt"
	"github.com/prongbang/staticfy/pkg/core"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

// Router the interface
type Router interface {
	Register()
}

type router struct {
	Handle Handler
}

func (r *router) Register() {
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

	if data.Staticfy.Routes.Upload.Path != "" {
		http.HandleFunc(data.Staticfy.Routes.Upload.Path, core.APIKeyMiddleware(core.JWTMiddleware(r.Handle.UploadHandler(data.Staticfy))))
	}
	if data.Staticfy.Routes.Delete.Path != "" {
		http.HandleFunc(data.Staticfy.Routes.Delete.Path, core.APIKeyMiddleware(core.JWTMiddleware(r.Handle.DeleteHandler(data.Staticfy))))
	}

	fmt.Println(banner)
	fmt.Printf(fmt.Sprintf("-> staticfy serving %s on :%d\n", data.Staticfy.Prefix, data.Staticfy.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", data.Staticfy.Port), nil))
}

// NewRouter is instance
func NewRouter(handle Handler) Router {
	return &router{
		Handle: handle,
	}
}
