package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application configuration parameters
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
