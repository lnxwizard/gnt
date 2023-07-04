# gnt
[gnt](https://github.com/lnxwizard/gnt) is command-line tool for creating [Go](https://github.com/golang/go) projects quickly.

![gnt-create-0 5 0-beta](https://github.com/lnxwizard/gnt/assets/91411319/d959f76a-c9d6-458e-ba66-c12657366816)

You can create Go project templates with `gnt` in your favorite terminal. For creating basic Go project just type:
```shell
gnt create myGoProject
```

# Installation
`gnt` will soon support cross-platform. So, you can use `gnt` in Linux, macOS and Windows operating systems.

## Linux
Open the `Releases` page and click to the latest release and download `gnt_{version}_linux_386.tar.gz`, `gnt_{version}_linux_amd64.tar.gz`, `gnt_{version}_linux_arm.tar.gz` `or gnt_{version}_linux_arm64.tar.gz`depending on your processor architecture. Add `gnt/bin` to the system path, open your favorite Terminal and type:
```shell
gnt --help
```
If you see the help message `gnt` has been successfully installed on your system.

## MacOS
Open the `Releases` page and click to the latest release and download the `gnt_{version}_darwin_arm64.zip` or `gnt_{version}_darwin_arm64.zip` file according to your processor architecture. Add `gnt/bin` to the system path, open your favorite Terminal and type:
```shell
gnt --help
```
If you see the help message `gnt` has been successfully installed on your system.

## Windows
Windows operating systems are not supported for now. Coming Soon...

# About
## Commands
- `bug`: Report bug
- `create`: Create Go projects
  - `--pkg`: Add pkg folder to your project.
  - `--internal`: Add internal folder to your project.
  - `--makefile`: Create Makefile.
- `help`: Help for gnt
- `repo`: Open GitHub repository
- `version`: Print the CLI version

## Flags
- `-h`, `--help`: Help for gnt
- `-v`, `--version`: Version for gnt

# Documentation
Documentation for `gnt` is coming soon...

# License
[MIT License](LICENSE)
