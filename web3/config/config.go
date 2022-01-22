package config

import "html/template"

// AppConfig hold app global configuration data
type AppConfig struct {
	UseCache  bool
	TmplCache map[string]*template.Template
}
