package markdown

import (
	"fmt"
	"strings"
	"testing"
)

func TestYAMLFrontMatter(t *testing.T) {

	// Create a new YAML frontmatter block
	fm, err := YAMLFrontMatter(map[string]interface{}{
		"title": "Awesome Feature",
		"date":  "2020-01-01",
		"tags":  []string{"feature", "major"},
	})

	fmt.Println(fm.String())

	// Check for errors
	if err != nil {
		t.Error(err)
	}

	// Check the output
	// ! The YAML string is extremely finnicky. Needs improvement.
	if fm.String() != strings.Join([]string{
		"---",
		"date: \"2020-01-01\"",
		"tags:",
		"    - feature",
		"    - major",
		"title: Awesome Feature",
		"", // TODO: This is ugly and needs to go away
		"---",
	}, "\n") {
		t.Errorf("Unexpected output: %s", fm.String())
	}

}
