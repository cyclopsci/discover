package discover

import (
	"os"
	"fmt"
	"regexp"
	"strings"
)

var (
	Languages = []Lang{
		PuppetFiles,
		PuppetModule,
		YamlFiles,
	}
	Results = make(map[string][]string)
)

type Lang struct {
	Key		string
	Extensions	[]string
	Paths		[]string
	Matchers	[]string
	IgnoredDirs	[]string
}

func Run(root string) map[string][]string {
	os.Chdir(root)
	tree, _ := walk(root)
	for _, lang := range Languages {
		Results[lang.Key] = analyze(lang, tree)
	}
	Results["root"] = []string{root}
	return Results
}

func analyze(l Lang, t Tree) []string {
	var base_path string
	var re_matcher *regexp.Regexp
	var matches []string
	candidates := t.Files
	for _, value := range l.IgnoredDirs {
		candidates = filter(value, candidates)
	}
	for _, value := range candidates {
		for _, p := range l.Paths {
			base_path = strings.Replace(value, p, "", 1)
			if match(p, value) && !contains(matches, base_path) {
				matches = append(matches, base_path)
			}
		}
		for _, m := range l.Matchers {
			re_matcher = regexp.MustCompile("(" + m + ")")
			if match(m, value) {
				base_path = re_matcher.ReplaceAllString(value, "")
				matches = append(matches, base_path)
			}
		}
		for _, e := range l.Extensions {
			if e == extension(value) {
				matches = append(matches, value)
			}
		}
	}
	return matches
}

func filter(needle string, haystack []string) []string {
	var result []string
	for _, value := range haystack {
		if !search(needle, value) {
			result = append(result, value)
		}
	}
	return result
}

func search(matcher string, value string) bool {
	exp := fmt.Sprintf("^.*(/)?%s(/)?", matcher)
	matched, _ := regexp.MatchString(exp, value)
	if matched {
		return true
	}
	return false
}

func match(matcher string, value string) bool {
	exp := fmt.Sprintf("^.*(/)?%s$", matcher)
	matched, _ := regexp.MatchString(exp, value)
	if matched {
		return true
	}
	return false
}

func contains(s []string, v string) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}
	return false
}

func extension(v string) string {
	sl := strings.Split(v, ".")
	return sl[len(sl)-1]
}
