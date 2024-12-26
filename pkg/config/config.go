package config

import "text/template"

// AppConfig holds application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
