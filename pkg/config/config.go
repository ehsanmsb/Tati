package config

import "text/template"

// AppConfig holds application configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
