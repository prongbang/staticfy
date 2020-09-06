package staticfy

import (
	"mime/multipart"
	"time"
)

const (
	StaticfyStaticsPath   = "./staticfy%s"
	StaticfyConfigYmlPath = "./staticfy/config.yml"
)

// Config is a model
type Config struct {
	Staticfy Staticfy `yaml:"staticfy"`
}

// Staticfy is a model
type Staticfy struct {
	Port      int    `yaml:"port"`
	Host      string `yaml:"host"`
	Directory string `yaml:"directory"`
	Prefix    string `yaml:"prefix"`
	Routes    struct {
		Upload Route `yaml:"upload"`
		Delete Route `yaml:"delete"`
	}
}

// Route is a model
type Route struct {
	Path      string   `yaml:"path"`
	Method    string   `yaml:"method"`
	Support   []string `yaml:"support"`
	Directory string   `yaml:"directory"`
}

type Assets struct {
	ID         int64                 `json:"id" db:"id"`
	URL        string                `json:"url" db:"url"`
	CreatedAt  time.Time             `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time             `json:"updatedAt" db:"updated_at"`
	UserID     string                `json:"userId,omitempty" db:"user_id"`
	FilePath   string                `json:"-" db:"file_path"`
	Host       string                `json:"-"`
	Directory  string                `json:"-"`
	Path       string                `json:"-"`
	Ext        string                `json:"-"`
	FileHeader *multipart.FileHeader `json:"-"`
}
