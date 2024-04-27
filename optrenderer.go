package mustache

type optRenderer struct {
	opts []TemplateOpt
}

func NewRenderer(opts... TemplateOpt) Renderer {
	return &optRenderer{opts: opts}
}

func (r *optRenderer) Render(data string, context ...interface{}) (string, error) {
	return render(r.opts, data, context...)
}

func (r *optRenderer) RenderRaw(data string, forceRaw bool, context ...interface{}) (string, error) {
	return renderRaw(r.opts, data, forceRaw, context...)
}

func (r *optRenderer) RenderPartials(data string, partials PartialProvider, context ...interface{}) (string, error) {
	return renderPartials(r.opts, data, partials, context...)
}

func (r *optRenderer) RenderPartialsRaw(data string, partials PartialProvider, forceRaw bool, context ...interface{}) (string, error) {
	return renderPartialsRaw(r.opts, data, partials, forceRaw, context...)
}

func (r *optRenderer) RenderInLayout(data string, layoutData string, context ...interface{}) (string, error) {
	return renderInLayout(r.opts, data, layoutData, context...)
}

func (r *optRenderer) RenderInLayoutPartials(data string, layoutData string, partials PartialProvider, context ...interface{}) (string, error) {
	return renderInLayoutPartials(r.opts, data, layoutData, partials, context...)
}

func (r *optRenderer) RenderFile(filename string, context ...interface{}) (string, error) {
	return renderFile(r.opts, filename, context...)
}

func (r *optRenderer) RenderFileInLayout(filename string, layoutFile string, context ...interface{}) (string, error) {
	return renderFileInLayout(r.opts, filename, layoutFile, context...)
}
