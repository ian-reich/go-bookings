package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// App Config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
