![Go](https://github.com/h2cone/lstars/workflows/Go/badge.svg)

# lstars

Lists repositories (include non-language) a user has starred.

## Installation

Download the [latest release](https://github.com/h2cone/lstars/releases).

## Usage

```shell
% ./lstars -h
Lists repositories a user has starred.

Usage:
  lstars [flags]

Flags:
      --direction string   asc or desc (default "desc")
  -h, --help               help for lstars
  -l, --lang string        language
      --num int            page num (default 1)
      --once               only request once
      --size int           page size (default 30)
      --sort string        created or updated (default "created")
  -u, --user string        username
```

### Examples

Go:

```shell
./lstars -u=h2cone -l=go
```

Non-language:

```shell
./lstars -u=h2cone -l=null
```

List all and only request once:

```shell
./lstars -u=h2cone --once
```

Paging and only request once:

```shell
./lstars -u=h2cone -l=java --num=2 --size=60 --once
```
