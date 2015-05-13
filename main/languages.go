package discover

var PuppetFiles = Lang{
	Key:		"puppet_files",
	Extensions:	[]string{"pp"},
	IgnoredDirs:	[]string{"spec", "pkg", "tests", "test"},
}

var PuppetModule = Lang{
	Key:		"puppet_modules",
	Paths:		[]string{"manifests/init.pp", "Modulefile", "metadata.json"},
	IgnoredDirs:	[]string{"spec", "pkg", "tests", "test"},
}

var Yaml = Lang{
	Key:		"yaml_files",
	Extensions:	[]string{"yml", "yaml"},
}
