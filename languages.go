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

var yamlFiles = language{
	Key:		"yaml_files",
	Extensions:	[]string{"yml", "yaml"},
}
