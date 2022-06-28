package config

import (
	"html/template"
	"log"
)

// AccConfig
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
