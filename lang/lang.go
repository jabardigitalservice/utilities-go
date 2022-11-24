package lang

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Lang struct {
	localizer *i18n.Localizer
	bundle    *i18n.Bundle
}

func New(path string, locale string) Lang {
	bundle := i18n.NewBundle(language.Indonesian)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile(path)
	localizer := i18n.NewLocalizer(bundle, locale)

	return Lang{
		localizer: localizer,
		bundle:    bundle,
	}
}

func (l Lang) GetMessage(ID string, TemplateData map[string]interface{}) (string, error) {
	return l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: ID},
		TemplateData:   TemplateData,
	})
}

func (l Lang) LoadMessageFile(path string) (*i18n.MessageFile, error) {
	return l.bundle.LoadMessageFile(path)
}

func (l Lang) NewLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(l.bundle, lang)
}

func (l Lang) GetMessageByLocalize(localizer *i18n.Localizer, ID string, TemplateData map[string]interface{}) (string, error) {
	return localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: ID},
		TemplateData:   TemplateData,
	})
}
