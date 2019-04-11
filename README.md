# chcase

CLI changing string cases

## Usage

```
$ chcase -h
NAME:
   chcase

USAGE:
   chcase [global options] command [command options] [arguments...]

VERSION:
   0.2.0

COMMANDS:
     snake, s             Convert input to snake_case
     screaming-snake, ss  Convert input to SCREAMING_SNAKE_CASE
     kebab, k             Convert input to kebab-case
     screaming-kebab, sk  Convert input to SCREAMING-KEBAB-CASE
     camel, c             Convert input to CamelCase
     lower-camel, lc      Convert input to lowerCamelCase
     help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Example

```
$ chcase c foo_bar_baz
FooBarBaz
$ chcase k FooBarBaz
foo-bar-baz
```

## LICENSE

[MIT](./LICENSE)