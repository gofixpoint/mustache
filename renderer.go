package mustache

type Renderer interface {
	// Render compiles a mustache template string and uses the the given data
	// source - generally a map or struct - to render the template and return
	// the output.
	Render(data string, context ...interface{}) (string, error)

	// RenderRaw compiles a mustache template string and uses the the given data
	// source - generally a map or struct - to render the template and return
	// the output.
	RenderRaw(data string, forceRaw bool, context ...interface{}) (string, error)

	// RenderPartials compiles a mustache template string and uses the the given
	// partial provider and data source - generally a map or struct - to render
	// the template and return the output.
	RenderPartials(data string, partials PartialProvider, context ...interface{}) (string, error)

	// RenderPartialsRaw compiles a mustache template string and uses the the
	// given partial provider and data source - generally a map or struct - to
	// render the template and return the output.
	RenderPartialsRaw(data string, partials PartialProvider, forceRaw bool, context ...interface{}) (string, error)

	// RenderInLayout compiles a mustache template string and layout "wrapper"
	// and uses the given data source - generally a map or struct - to render
	// the compiled templates and return the output.
	RenderInLayout(data string, layoutData string, context ...interface{}) (string, error)

	// RenderInLayoutPartials compiles a mustache template string and layout
	// "wrapper" and uses the given data source - generally a map or struct - to
	// render the compiled templates and return the output.
	RenderInLayoutPartials(data string, layoutData string, partials PartialProvider, context ...interface{}) (string, error)

	// RenderFile loads a mustache template string from a file and compiles it,
	// and then uses the given data source - generally a map or struct - to
	// render the template and return the output.
	RenderFile(filename string, context ...interface{}) (string, error)

	// RenderFileInLayout loads a mustache template string and layout "wrapper"
	// template string from files and compiles them, and  then uses the the
	// given data source - generally a map or struct - to render the compiled
	// templates and return the output.
	RenderFileInLayout(filename string, layoutFile string, context ...interface{}) (string, error)
}
