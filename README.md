# GoLiGen

[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://gitlab.com/tmaczukin/goligen/raw/master/LICENSE)
[![Build status](https://gitlab.com/tmaczukin/goligen/badges/master/build.svg)](https://gitlab.com/tmaczukin/goligen/commits/master)

Simple license file generator written in GO

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [List available license templates](#list-available-license-templates)
  - [Generate the license](#generate-the-license)
- [Default configuration](#default-configuration)
- [Own license templates](#own-license-templates)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Installation

Currently there is no source of compilet binaries. You can install the tool from source
with:

```bash
$ go install gitlab.com/tmaczukin/goligen
```

To do this you should have a configured GO environment.

Soon there should be ready a repository of compiled binaries for multiple platforms, so the installation will be much
simpler and it will not require a GO development environment.

## Usage

GoLiGen is a quite simple command line tool. It has two main commands:

- `list` - to [list the available license templates](#list-available-license-templates)
- `generate` - to [generate the license itself](#generate-the-license)

The global options description you can find by using the `help` command:

```bash
$ out/binaries/goligen help
(...)

GLOBAL OPTIONS:
   --debug              Set debug mode [$DEBUG]
   --log-level "info"   Set log level (options: debug, info, warn, error, fatal, panic) [$LOG_LEVEL]
   --help, -h           show help
   --version, -v        print the version
```

### List available license templates

```
goligen list
```

To list available license templates you need to use the `list` command. It will output the list of template's IDs on the
standard output:

```bash
$ ./goligen list
INFO[0000] Available internal license templates:
INFO[0000]   GPL-2.0
INFO[0000]   MIT
```

If you are using [user templates](#own-license-templates) the you will see also all available templates from the user's
configuration directory:

```bash
$ goligen list
INFO[0000] Available internal license templates:
INFO[0000]   MIT
INFO[0000]   GPL-2.0
INFO[0000] Available user license templates:
INFO[0000]   GPL-3.0
```

### Generate the license

```
goligen generate [command options] [ID]
```

To generate the license you need to use the `generate` command. It will generate the license text and return it on the
standard output.

`generate` command requires you to provide the `ID` argument and at least one `date`/`name` pair for the copyright:

```bash
$ goligen generate -d 2016 -n "Example Inc." MIT
INFO[0000] Generating license: MIT
INFO[0000] Generating to standard output
The MIT License (MIT)

Copyright (c) 2016 Example Inc.
(...)
```

If you will not provide the `ID` and the `copyrights` values, then GoLiGen will return an error. Also remember, that
providing `copyright` data you need to add a `date`/`name` pair. If you ommit one `date` or one `name` then an error
will be raised.

```bash
$ goligen generate
FATA[0000] You must provide 'license ID' as a first command argument

$ goligen generate MIT
FATA[0000] There must be at least one copyright-date/copyright-name pair

$ goligen generate -d 2016 MIT
FATA[0000] Copyright-date and copyright-name must be added in pairs
```

If you are using those values often, then consider creating the [configuration file](#default-configuration).

If you want to save the license into file (instead of standard output) than you can use `-o` option. However remebmer
that GoLiGen will refuse to reqrite the existing file. To force this you should use `-f` flag.

Options `-d` (for `date`) and `-n` (for `name`) you can use multiple times but always in pairs.

If you want to use a [user template](#own-license-templates) then you must also provide the `-u` flag.

Full `generate` command options list:

```bash
$ goligen help generate
NAME:
   goligen generate - Generate license

USAGE:
   goligen generate [command options] [ID]

OPTIONS:
   --copyright-date, -d [--copyright-date option --copyright-date option]       Date of copyright owner
   --copyright-name, -n [--copyright-name option --copyright-name option]       Name of the copyright owner
   -o, --output                                                                 Output file
   -f, --force-output                                                           Rewrite file if exists
   -u, --use-user-template                                                      Use user template instead of internal
```

## Default configuration

If you are creating new projects often, and you use similar data each time, then you may want to not repeat yourself
in command line. Instead you can create a configuration file with default values for the generator.

File needs to be located in your home directory: `~/.goligen/config.toml`. In the configuration file you can set
a default license `ID` and as more default `copyrights` as you want.

An example of `config.toml` file:

```toml
default_license_id = "MIT"

[[default_copyrights]]
  date = "2016"
  name = "Example Inc."
```

## Own license templates

If you want to use some Free/OpenSource license that is not included in GoLiGen, then it is highly recommended to create
a Merge Request with a new license template. I would like this software to be usable for the whole FOSS community. But
there is plenty of FOSS-like licenses. I will be adding templates in my free time, but this process may take some time.
So help is welcome.

If you want to use a license template that is not present in internal templates list (for example some not added FOSS
license, but also your own properiaty license) you may add your own _user template_.

For this you need to create a templates directory in the configuration directory: `~/.goligen/templates/`. Inside of the
directory you can add license template files. Each file must have a name same as the license `ID` that you would like
to use. The name must match the regular expression: `[A-Za-z0-9\\-\\.+]+$`.

After adding your own license template you will see it on the list:

```bash
$ goligen list
INFO[0000] Available internal license templates:
INFO[0000]   GPL-2.0
INFO[0000]   MIT
INFO[0000] Available user license templates:
INFO[0000]   GPL-3.0
```

You can use the template by adding the `-u` flag to the `generate` command.

## License

This is free sofware licensed under MIT license. See LICENSE file.
