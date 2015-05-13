package discover

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

type JsonResults struct {
	Root		[]string	`json:"root"`
	PuppetModules	[]string	`json:"puppet_modules"`
	PuppetFiles	[]string	`json:"puppet_files"`
	YAMLFiles	[]string	`json:"yaml_files"`
}

func TestDiscover(t *testing.T) {
	json_output := Discover(".")
	response := &JsonResults{}
	json.Unmarshal([]byte(json_output), &response)
	assert.Equal(t, []string{"."}, response.Root)
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

