package config

import "text/template"

/// AppConfig содержит настройки приложения
type AppConfig struct {
	templateCache map[string]*template.Template
}
