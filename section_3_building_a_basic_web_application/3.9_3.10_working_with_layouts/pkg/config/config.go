package config

import "text/template"

/// AppConfig содержит настройки приложения
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
