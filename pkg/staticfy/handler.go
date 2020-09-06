package staticfy

import (
	"encoding/json"
	"fmt"
	"github.com/prongbang/staticfy/pkg/core"
	"log"
	"net/http"
	"path/filepath"
)

// Handler the interface
type Handler interface {
	UploadHandler(static Staticfy) func(w http.ResponseWriter, r *http.Request)
	DeleteHandler(static Staticfy) func(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	UseCase UseCase
}

func (h *handler) UploadHandler(static Staticfy) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(static.Routes.Upload.Method, "-->", r.URL.Path)
		if err := core.AllowedMethod(static.Routes.Upload.Method, r.Method); err != nil {
			core.MethodNotAllowed(w, "Method Not Allowed")
			return
		}

		// Handle get user from JWT
		userId := core.Get(r, "sub")

		// Handle multipath file
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			core.BadRequest(w, "Cannot Parse Multipart Form")
			return
		}
		_, fileHead, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			core.BadRequest(w, "Cannot get file, please use field name `file`")
			return
		}

		ext := filepath.Ext(fileHead.Filename)
		if _, err := core.IsWhiteListExt(static.Routes.Upload.Support, ext); err != nil {
			core.BadRequest(w, err.Error())
			return
		}

		asset, err := h.UseCase.Upload(Assets{
			UserID:     userId,
			Directory:  fmt.Sprintf(StaticfyStaticsPath, static.Directory) + static.Routes.Upload.Directory,
			Path:       static.Prefix + static.Routes.Upload.Directory,
			Host:       static.Host,
			Ext:        ext,
			FileHeader: fileHead,
		})

		if err != nil {
			core.BadRequest(w, "Cannot Upload File")
			return
		}

		core.Created(w, asset)
	}
}

func (h *handler) DeleteHandler(static Staticfy) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(static.Routes.Delete.Method, "-->", r.URL.Path)
		if err := core.AllowedMethod(static.Routes.Delete.Method, r.Method); err != nil {
			core.MethodNotAllowed(w, "Method Not Allowed")
			return
		}

		// Handle get user from JWT
		userId := core.Get(r, "sub")
		if userId == "" {
			core.BadRequest(w, "Please set user id to `sub` in JWT")
			return
		}

		// Reading form values
		if err := r.ParseForm(); err != nil {
			core.BadRequest(w, "Cannot Parse Form")
			return
		}

		asset := Assets{}
		if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
			core.BadRequest(w, "Bad Request")
			return
		}

		if asset.ID <= 0 {
			core.BadRequest(w, "Bad Request")
			return
		}

		asset, err := h.UseCase.Delete(Assets{
			ID:     asset.ID,
			UserID: userId,
		})

		if err != nil {
			core.NotFound(w, "Not Found")
			return
		}

		core.Ok(w, asset)
	}
}

// NewHandler is instance
func NewHandler(useCase UseCase) Handler {
	return &handler{
		UseCase: useCase,
	}
}
