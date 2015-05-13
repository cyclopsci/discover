# Discover
Discovers where all the automation lies

# Usage

    go build *.go
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
* *Key*: name to associate the returned values in the JSON results
  Type: _string_
* *Ext*: look and include all files matching this extension (omit a leading '.')
  Type: _string_
* *Paths*: include the folders that contain these items as children (assumes each item is the end of the path)
  Type: _slice [string]_
* *Matchers*: include the folders that contain these regex items as children (assumes each item is the end of the path)
  Type: _slice [string]_
* *IgnoredDirs*: ignore any path that contains these directories
  Type: _slice [string]_
