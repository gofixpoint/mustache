package mustache

// Contains the internal implementation functions for rendering so we can
// maintain a pretty API.

func render(opts []TemplateOpt, data string, context ...interface{}) (string, error) {
	return renderRaw(opts, data, false, context...)
}

// RenderRaw compiles a mustache template string and uses the the given data
// source - generally a map or struct - to render the template and return the
// output.
func renderRaw(opts []TemplateOpt, data string, forceRaw bool, context ...interface{}) (string, error) {
	return renderPartialsRaw(opts, data, nil, forceRaw, context...)
}

// RenderPartials compiles a mustache template string and uses the the given partial
// provider and data source - generally a map or struct - to render the template
// and return the output.
func renderPartials(opts []TemplateOpt, data string, partials PartialProvider, context ...interface{}) (string, error) {
	return renderPartialsRaw(opts, data, partials, false, context...)
}

// RenderPartialsRaw compiles a mustache template string and uses the the given
// partial provider and data source - generally a map or struct - to render the
// template and return the output.
func renderPartialsRaw(opts []TemplateOpt, data string, partials PartialProvider, forceRaw bool, context ...interface{}) (string, error) {
	var tmpl *Template
	var err error
	if partials == nil {
		tmpl, err = ParseStringRaw(data, forceRaw, opts...)
	} else {
		tmpl, err = ParseStringPartialsRaw(data, partials, forceRaw, opts...)
	}
	if err != nil {
		return "", err
	}
	return tmpl.Render(context...)
}

// RenderInLayout compiles a mustache template string and layout "wrapper" and
// uses the given data source - generally a map or struct - to render the
// compiled templates and return the output.
func renderInLayout(opts []TemplateOpt, data string, layoutData string, context ...interface{}) (string, error) {
	return renderInLayoutPartials(opts, data, layoutData, nil, context...)
}

// RenderInLayoutPartials compiles a mustache template string and layout
// "wrapper" and uses the given data source - generally a map or struct - to
// render the compiled templates and return the output.
func renderInLayoutPartials(opts []TemplateOpt, data string, layoutData string, partials PartialProvider, context ...interface{}) (string, error) {
	var layoutTmpl, tmpl *Template
	var err error
	if partials == nil {
		layoutTmpl, err = ParseString(layoutData, opts...)
	} else {
		layoutTmpl, err = ParseStringPartials(layoutData, partials, opts...)
	}
	if err != nil {
		return "", err
	}

	if partials == nil {
		tmpl, err = ParseString(data, opts...)
	} else {
		tmpl, err = ParseStringPartials(data, partials, opts...)
	}

	if err != nil {
		return "", err
	}

	return tmpl.RenderInLayout(layoutTmpl, context...)
}

// RenderFile loads a mustache template string from a file and compiles it, and
// then uses the given data source - generally a map or struct - to render the
// template and return the output.
func renderFile(opts []TemplateOpt, filename string, context ...interface{}) (string, error) {
	tmpl, err := ParseFile(filename, opts...)
	if err != nil {
		return "", err
	}
	return tmpl.Render(context...)
}

// RenderFileInLayout loads a mustache template string and layout "wrapper"
// template string from files and compiles them, and  then uses the the given
// data source - generally a map or struct - to render the compiled templates
// and return the output.
func renderFileInLayout(opts []TemplateOpt, filename string, layoutFile string, context ...interface{}) (string, error) {
	layoutTmpl, err := ParseFile(layoutFile, opts...)
	if err != nil {
		return "", err
	}

	tmpl, err := ParseFile(filename, opts...)
	if err != nil {
		return "", err
	}
	return tmpl.RenderInLayout(layoutTmpl, context...)
}