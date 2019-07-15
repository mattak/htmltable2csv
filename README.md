# htmltable2csv 

Command line converter from html table to csv.

## Install

```shell
go install github.com/mattak/htmltable2csv
```

## Usage

```shell
$ htmltable2csv -h
NAME:
   htmltable2csv - html table structure to csv

USAGE:
   htmltable2csv [global options] [argument_filename]

VERSION:
   0.0.1

DESCRIPTION:
   argument filename is optional. default input is from stdin.

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

```shell
$ cat << __HTML__ | htmltable2csv
<table>
<tr><th>ID</th><th>NAME</th></tr>
<tr><td>1</td><td>one</td></tr>
<tr><td>2</td><td>two</td></tr>
</table>
__HTML__
ID,NAME
1,one
2,two
```

