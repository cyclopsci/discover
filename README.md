# Discover [![Build Status](https://circleci.com/gh/cyclopsci/discover.svg?&style=shield&circle-token=bb753ccb4400119162fb978d41adcdc601b58383)](https://circleci.com/gh/cyclopsci/discover/)
Discovery utility that returns the location of specific automation types

# Installation

    go get github.com/cyclopsci/discover/...

# Usage

    ./discover -directory <path_to_automation>

# Sample Output

    {
      "puppet_files": [
        "/etc/puppet/modules/puppet-automysqlbackup/manifests/backup.pp",
        "/etc/puppet/modules/puppet-automysqlbackup/manifests/init.pp",
        "/etc/puppet/modules/puppet-automysqlbackup/manifests/params.pp",
        "/etc/puppet/modules/puppet-flowtools/manifests/config.pp",
        "/etc/puppet/modules/puppet-flowtools/manifests/device.pp",
        "/etc/puppet/modules/puppet-flowtools/manifests/init.pp",
        "/etc/puppet/modules/puppet-flowtools/manifests/install.pp",
        "/etc/puppet/modules/puppet-flowtools/manifests/params.pp",
        "/etc/puppet/modules/puppet-flowtools/manifests/service.pp",
        "/etc/puppet/modules/stdlib/manifests/init.pp",
        "/etc/puppet/modules/stdlib/manifests/stages.pp"
      ],
      "puppet_modules": [
        "/etc/puppet/modules/puppet-automysqlbackup/",
        "/etc/puppet/modules/puppet-flowtools/",
        "/etc/puppet/modules/stdlib/"
      ],
      "root": [
        "/etc/puppet/modules/"
      ]
    }

# Types

## ```Lang```
* **Key** (_string_): name to associate the returned values in the JSON results
* **Extensions** (_slice [string]_): include all files matching any of the specified extension (omit a leading '.')
* **Paths** (_slice [string]_): include the folders that contain these items as children (assumes each item is the end of the path)
* **Matchers** (_slice [string]_): include the folders that contain these regex items as children (assumes each item is the end of the path)
* **IgnoredDirs** (_slice [string]_): ignore any path that contains these directories

# Developing

    go test -v ./
    go build ./...
    go install ./...
