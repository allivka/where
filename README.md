# where

Implementation of which(1) linux utility in golang.
It searches the target file which is not directory in all the directories listed in *PATH* environment variable

## usage

```shell
Usage:
    where [-h|--help] target <files: name1 name2 ... nameN>"
```

## build

Build with the build script that requires go compiler:

```shell
./build.sh
```

or manually:

```shell
go build -o where main.go
```

You can also build for every platform with:

```shell
./multibuild.sh
```

## installation

just move the *where* executable at the any directory set in path
