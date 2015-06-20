package discover

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestLangPuppetFile(t *testing.T) {
	gomega.RegisterTestingT(t)

	walkDirectory("./testing/fixtures/")
	results := analyzeTree([]language{puppetFile,puppetModule}, tree)

	gomega.Expect(results[puppetFile.Key]).To(gomega.Equal([]string{
		"puppet/manifest.pp",
		"puppet/modulefail/manifests/class.pp",
	}))
}

func TestLangPuppetModule(t *testing.T) {
	gomega.RegisterTestingT(t)

	walkDirectory("./testing/fixtures/")
	results := analyzeTree([]language{puppetModule}, tree)

	gomega.Expect(results[puppetModule.Key]).To(gomega.Equal([]string{"puppet/module/"}))
}

func TestLangAnsibleRole(t *testing.T) {
	gomega.RegisterTestingT(t)

	walkDirectory("./testing/fixtures")
	results := analyzeTree([]language{ansibleRole}, tree)

	gomega.Expect(results[ansibleRole.Key]).To(gomega.Equal([]string{"ansible/role/"}))
}

func TestLangAnsiblePlaybook(t *testing.T) {
	gomega.RegisterTestingT(t)

	walkDirectory("testing/fixtures")
	results := analyzeTree([]language{ansiblePlaybook}, tree)

	gomega.Expect(results[ansiblePlaybook.Key]).To(gomega.Equal([]string{"ansible/playbook.yml"}))
}
