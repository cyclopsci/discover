package discover

import (
	"testing"
	"github.com/onsi/gomega"
)

func TestRun(t *testing.T) {
	gomega.RegisterTestingT(t)
	results := Run(".")
	gomega.Expect(results["root"]).To(gomega.Equal([]string{"./"}))
}

func TestWalkDirectory(t *testing.T) {
	gomega.RegisterTestingT(t)
	walkDirectory(".")
	gomega.Expect(tree).To(gomega.ContainElement(".gitignore"))
	gomega.Expect(tree).To(gomega.ContainElement("README.md"))
	gomega.Expect(tree).To(gomega.ContainElement("LICENSE"))
	gomega.Expect(tree).To(gomega.ContainElement("discover/main.go"))
	gomega.Expect(tree).To(gomega.ContainElement("discover.go"))
	gomega.Expect(tree).To(gomega.ContainElement("languages.go"))
}

func TestAnalyze(t *testing.T) {
	gomega.RegisterTestingT(t)
	lang := language{
		Key:		"lang",
		Extensions:	[]string{"z"},
		Paths:		[]string{"a/b/x.z"},
		PathMatchers:	[]string{"a/b/(c|d).z"},
		IgnoredDirs:	[]string{"i"},
	}
	tree := []string{
		"/1/a/b/x.z",
		"/1/a/b/c.z",
		"/1/a/b/c.y",
		"/1/d/e.z",
		"/i/d/f.z",
		"/a/d/f.z",
	}
	results := analyzeTree([]language{lang}, tree)
	gomega.Expect(results[lang.Key]).To(gomega.ContainElement("/1/"))
	gomega.Expect(results[lang.Key]).To(gomega.ContainElement("/a/d/f.z"))
	gomega.Expect(results[lang.Key]).ToNot(gomega.ContainElement("/1/a/b/c.y"))
}

func TestSearch(t *testing.T) {
	gomega.RegisterTestingT(t)
	gomega.Expect(search("pkg", "/dir1/pkg/subdir/")).To(gomega.BeTrue())
	gomega.Expect(search("dir2", "/dir1/pkg/subdir/")).To(gomega.BeFalse())
}

func TestStringInSlice(t *testing.T) {
	gomega.RegisterTestingT(t)
	s := []string{"a","b"}
	gomega.Expect(stringInSlice("a", s)).To(gomega.BeTrue())
	gomega.Expect(stringInSlice("c", s)).To(gomega.BeFalse())
}
