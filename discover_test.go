package discover

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	results := Run(".")
	assert.Equal(t, []string{"."}, results["root"])
}

func TestAnalyze(t *testing.T) {
	lang := Lang{
		Key:		"lang",
		Extensions:	[]string{"z"},
		Paths:		[]string{"a/b/x.z"},
		Matchers:	[]string{"a/b/(c|d).z"},
		IgnoredDirs:	[]string{"i"},
	}
	tree := Tree{
		Files: []string{
			"/1/a/b/x.z",
			"/1/a/b/c.z",
			"/1/a/b/c.y",
			"/1/d/e.z",
			"/i/d/f.z",
		},
	}
	results := analyze(lang, tree)
	assert.Contains(t, results, "/1/a/b/x.z")
	assert.Contains(t, results, "/1/a/b/c.z")
	assert.Contains(t, results, "/1/d/e.z")
	assert.NotContains(t, results, "/1/a/b/c.y")
	assert.NotContains(t, results, "/i/d/e.z")
}

func TestFilter(t *testing.T) {
	haystack := []string{"a", "b"}
	needle := "a"
	result := filter(needle, haystack)
	assert.Equal(t, []string{"b"}, result)
}

func TestSearch(t *testing.T) {
	assert.True(t, search("pkg", "/dir1/pkg/subdir/"))
	assert.False(t, search("dir2", "/dir1/pkg/subdir/"))
}

func TestMatch(t *testing.T) {
	assert.True(t, match("manifests/init.pp", "/dir1/manifests/init.pp"))
	assert.False(t, match("manifests/init.pp", "/dir1/manifests/init.pp/"))
}

func TestContains(t *testing.T) {
	s := []string{"a","b"}
	assert.True(t, contains(s, "a"))
	assert.False(t, contains(s, "c"))
}

func TestExtension(t *testing.T) {
	assert.Equal(t, "pp", extension("/test.one/test.pp"))
	assert.NotEqual(t, "pp", extension("/test.one/test.p"))
}

