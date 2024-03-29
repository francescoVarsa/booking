package config

import (
	"booking/internal/models"
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

//AppConfig hold the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
