package markdown

import (
	"testing"
)

func runMarkdown(input string) string {
	return string(MarkdownCommon([]byte(input)))
}

func doTests(t *testing.T, tests []string) {
	// catch and report panics
	var candidate string
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("\npanic while processing [%#v]: %s\n", candidate, err)
		}
	}()

	for i := 0; i+1 < len(tests); i += 2 {
		input := tests[i]
		candidate = input
		expected := tests[i+1]
		actual := runMarkdown(candidate)
		if actual != expected {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nActual  [%#v]",
				candidate, expected, actual)
		}

		// now test every substring to stress test bounds checking
		if !testing.Short() {
			for start := 0; start < len(input); start++ {
				for end := start + 1; end <= len(input); end++ {
					candidate = input[start:end]
					_ = runMarkdown(candidate)
				}
			}
		}
	}
}

func TestDocument(t *testing.T) {
	var tests = []string{
		// Empty document.
		"",
		"",

		" ",
		"",

		// This shouldn't panic.
		// https://github.com/russross/blackfriday/issues/172
		"[]:<",
		"<p>[]:&lt;</p>\n",

		// This shouldn't panic.
		// https://github.com/russross/blackfriday/issues/173
		"   [",
		"<p>[</p>\n",
	}
	doTests(t, tests)
}
