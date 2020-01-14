# goshellshocked
[![Go Report Card](https://goreportcard.com/badge/github.com/ShiraazMoollatjie/goshellshocked?style=flat-square)](https://goreportcard.com/report/github.com/ShiraazMoollatjie/goshellshocked)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/ShiraazMoollatjie/goshellshocked)
[![Build status](https://ci.appveyor.com/api/projects/status/brw55guj4gmq0m1v/branch/master?svg=true)](https://ci.appveyor.com/project/ShiraazMoollatjie/goshellshocked/branch/master)


> Shellshocked because teenage mutant ninja turtles reference. Also, you'll be surprised at what your history files conjures up.

`goshellshocked` is a productivity helper that allows us to determine worthwhile aliases to add as part of our day to day activities.

What's unique is that it will try and search for any file that contains **_history** in your home directory to build up the frequency list. So far it supports:

* Bash (bash_history)
* ZSH (zsh_history)
* Fish (fish_history)

It will pick up other files and default to the same parsing strategy as that of **bash**. If more shells need to be supported, then add a ticket to github along with some example lines to parse.

> ### A note about passwords
> goshellshocked, in its current state, will print out passwords if they appear in your shell history. So be very careful if you decide to share your results for some reason. 

# Install and run

To run `goshellshocked`, we retrieve it with:

```sh
go get -u github.com/ShiraazMoollatjie/goshellshocked/cmd/goshellshocked/
```

Then we can run by simply running:
```sh
goshellshocked
```

With the default options, we get the following kind of output:
```
2019/12/26 13:31:58 Found history file: .bash_history
2019/12/26 13:31:58 Found history file: .zsh_history
2019/12/26 13:31:58 Frequency: 66, Command: ls
2019/12/26 13:31:58 Frequency: 33, Command: ansible-playbook -i production main.yml -K
2019/12/26 13:31:58 Frequency: 24, Command: git push origin master
2019/12/26 13:31:58 Frequency: 19, Command: code .
2019/12/26 13:31:58 Frequency: 13, Command: clear
2019/12/26 13:31:58 Frequency: 12, Command: go run cmd/goshellshocked/main.go
2019/12/26 13:31:58 Frequency: 12, Command: cd ..
2019/12/26 13:31:58 Frequency: 10, Command: ansible-lint main.yml
```

# Flags

The following flags are available for customizing `goshellshocked`:

**--exclude [comma_separated_terms]**: Exclude a list of comma separated terms. If your command uses a comma, you can wrap the command in quotes. E.g. _ls,cd ..,"echo ,"_

**--ignore [minimum_frequency]**: Filter out commands that have a minimum frequency count of the provided amount. This value is exclusive. Defaults to a value of 3.

**--output [console|json|yaml]**: Change the output format of the results. Valid values are `console`, `yaml` and `json`. Invalid values will default to `console`. Defaults to `console`.

**--outputDir [dir]**: Change the output directory of the results when using file based outputs. Defaults to the current working directory.

# Example Usage

Run with default settings and print everything to the console:
```sh
goshellshocked
```

Run and include all past commands in the output
```sh
goshellshocked --ignore=0
```

Run and include all commands that have been used at least 5 times and also ignore very common commands like `ls` and `cd ..`
```sh
goshellshocked --ignore=4 --exclude=ls,"cd ..","git status"
```

# Contributions

Contributions and ideas are always welcomed ❤️.