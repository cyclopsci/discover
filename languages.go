package discover

var puppetFiles = language{
	Key:		"puppet_files",
	Extensions:	[]string{"pp"},
	IgnoredDirs:	[]string{"spec", "pkg", "tests", "test"},
}

var puppetModule = language{
	Key:		"puppet_modules",
	Paths:		[]string{"manifests/init.pp", "Modulefile", "metadata.json"},
	IgnoredDirs:	[]string{"spec", "pkg", "tests", "test"},
}

var ansibleRoles = language{
	Key:		"ansible_roles",
	Paths:		[]string{"tasks/main.yml", "meta/main.yml"},
}

var ansiblePlaybooks = language{
	Key:		"ansible_playbooks",
	ContentRegex:	[]string{"^(\\s+)?([-]\\s)?hosts: \\S+"},
}
