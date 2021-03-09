package config

import "html/template"

// AppConfig holds the apps config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
