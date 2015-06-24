package discover

var puppetManifest = language{
	Key:		     "puppet_manifests",
	Extensions:	 []string{"pp"},
	IgnoredDirs: []string{"spec", "pkg", "tests", "test"},
}

var puppetModule = language{
	Key:		     "puppet_modules",
	Paths:		   []string{"manifests/init.pp", "Modulefile", "metadata.json"},
	IgnoredDirs: []string{"spec", "pkg", "tests", "test"},
}

var ansibleRole = language{
	Key:   "ansible_roles",
	Paths: []string{"tasks/main.yml", "meta/main.yml"},
}

var ansiblePlaybook = language{
	Key:		      "ansible_playbooks",
	ContentRegex: []string{
		"^(\\s+)hosts:\\s\\S+$",
		"^-(\\s+)hosts:\\s\\S+$",
	},
}
