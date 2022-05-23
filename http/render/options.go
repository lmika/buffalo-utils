package render

import "html/template"

type ConfigOption func(config *Config)

func WithFuncs(funcs template.FuncMap) ConfigOption {
	return func(config *Config) {
		config.funcMaps = funcs
	}
}

// WithFrame adds a frame template, which is used to render the contents of template used
// in the request within another template.  Frame templates are evaluated in reverse order,
// meaning any later fame templates are rendered in earlier order.
func WithFrame(frameTemplate string) ConfigOption {
	return func(config *Config) {
		config.frameTemplates = append(config.frameTemplates, frameTemplate)
	}
}
