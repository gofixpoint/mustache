package mustache

import (
	"fmt"
	"testing"
)

func Test_SingleCurlyBraces(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		tmpl, err := ParseString(
			"Testing {single_args} curly {braces text}.",
			WithSingleCurlyTags(),
		)
		requireNilError(t, err)
		txt, err := tmpl.Render(
			map[string]string{
				"single_args": "single",
				"braces text": "braces",
			},
		)
		requireNilError(t, err)
		requireEqual(t, "Testing single curly braces.", txt)
	})

	t.Run("escaping HTML", func(t *testing.T) {
		tmpl, err := ParseString(
			"Single braces escape HTML by default: {compareTxt}",
			WithSingleCurlyTags(),
		)
		requireNilError(t, err)
		txt, err := tmpl.Render(map[string]string{"compareTxt": "5 < 10"})
		requireNilError(t, err)
		requireEqual(t, "Single braces escape HTML by default: 5 &lt; 10", txt)
	})

	t.Run("raw", func(t *testing.T) {
		tmpl, err := ParseString(
			"Double {{braces}} are raw.",
			WithSingleCurlyTags(),
		)
		requireNilError(t, err)
		txt, err := tmpl.Render(map[string]string{
			"braces": "5 < 10",
		})
		requireNilError(t, err)
		requireEqual(t, "Double 5 < 10 are raw.", txt)
	})
}

func Test_NoEscaping(t *testing.T) {
	expected := "We disabled HTML escaping: 5 < 10"
	args := map[string]string{"compareTxt": "5 < 10"}

	type TestParams struct {
		input string
		opts []TemplateOpt
	}
	tests := []TestParams{
		{
			"We disabled HTML escaping: {{compareTxt}}",
			[]TemplateOpt{WithHTMLEscape(false)},
		},
		{
			"We disabled HTML escaping: {{{compareTxt}}}",
			[]TemplateOpt{WithHTMLEscape(false)},
		},
		{
			"We disabled HTML escaping: {compareTxt}",
			[]TemplateOpt{WithHTMLEscape(false), WithSingleCurlyTags()},
		},
		{
			"We disabled HTML escaping: {{compareTxt}}",
			[]TemplateOpt{WithHTMLEscape(false), WithSingleCurlyTags()},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tmpl, err := ParseString(test.input, test.opts...)
			requireNilError(t, err)
			txt, err := tmpl.Render(args)
			requireNilError(t, err)
			requireEqual(t, expected, txt)
		})
	}
}

func requireNilError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

func requireEqual(t *testing.T, expected string, actual string) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}
