package discover

import (
	"testing"

	"github.com/onsi/gomega"
)

var (
	root string = "./testing/fixtures/"
	displayRoot string = "./testing/fixtures/"
)

func TestLangPuppetFile(t *testing.T) {
	gomega.RegisterTestingT(t)

	tree := walkDirectory(root)
	results := analyzeTree(root, displayRoot, []language{puppetFile,puppetModule}, tree)

	gomega.Expect(results[puppetFile.Key]).To(gomega.Equal([]string{
		"testing/fixtures/puppet/manifest.pp",
		"testing/fixtures/puppet/modulefail/manifests/class.pp",
	}))
}

func TestLangPuppetModule(t *testing.T) {
	gomega.RegisterTestingT(t)

	tree := walkDirectory(root)
	results := analyzeTree(root, displayRoot, []language{puppetModule}, tree)

	gomega.Expect(results[puppetModule.Key]).To(gomega.Equal([]string{"testing/fixtures/puppet/module/"}))
}

func TestLangAnsibleRole(t *testing.T) {
	gomega.RegisterTestingT(t)

	tree := walkDirectory(root)
	results := analyzeTree(root, displayRoot, []language{ansibleRole}, tree)

	gomega.Expect(results[ansibleRole.Key]).To(gomega.Equal([]string{"testing/fixtures/ansible/role/"}))
}

func TestLangAnsiblePlaybook(t *testing.T) {
	gomega.RegisterTestingT(t)

	tree := walkDirectory(root)
	results := analyzeTree(root, displayRoot, []language{ansiblePlaybook}, tree)

	gomega.Expect(results[ansiblePlaybook.Key]).To(gomega.Equal([]string{"testing/fixtures/ansible/playbook.yml"}))
}
