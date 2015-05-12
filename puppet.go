package main

var PuppetFiles = Lang{
	Key:		"puppet_files",
	Ext:		"pp",
	IgnoredDirs:	[]string{"spec", "pkg", "tests", "test"},
}

var PuppetModule = Lang{
	Key:		"puppet_modules",
	Paths:		[]string{"manifests/init.pp", "Modulefile", "metadata.json"},
	RequiredDirs:	[]string{"manifests"},
	IgnoredDirs:	[]string{"spec", "pkg", "tests", "test"},
}

