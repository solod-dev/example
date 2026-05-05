# So by example

A hands-on introduction to [Solod](https://github.com/solod-dev/solod) (So) — a subset of Go that translates to C.

This repo contains small [examples](#examples) that show individual features, as well as larger [apps](#apps) that demonstrate how everything works together.

## Examples

Language:

[Hello world](lang/01-hello/main.go) •
[Values](lang/02-values/main.go) •
[Variables](lang/03-variables/main.go) •
[Constants](lang/04-constants/main.go) •
[For](lang/05-for/main.go) •
[If/else](lang/06-if-else/main.go) •
[Switch](lang/07-switch/main.go) •
[Arrays](lang/08-arrays/main.go) •
[Slices](lang/09-slices/main.go) •
[Maps](lang/10-maps/main.go) •
[Functions](lang/11-functions/main.go) •
[Multiple returns](lang/12-returns/main.go) •
[Variadic functions](lang/13-variadics/main.go) •
[For-range](lang/16-range/main.go) •
[Pointers](lang/17-pointers/main.go) •
[Strings and runes](lang/18-strings/main.go) •
[Structs](lang/19-structs/main.go) •
[Methods](lang/20-methods/main.go) •
[Interfaces](lang/21-interfaces/main.go) •
[Enums](lang/22-enums/main.go) •
[Errors](lang/26-errors/main.go) •
[Panic](lang/27-panic/main.go) •
[Defer](lang/28-defer/main.go)

Standard library:

[Memory](stdlib/01-memory/main.go) •
[Slices](stdlib/02-slices/main.go) •
[Maps](stdlib/03-maps/main.go) •
[C interop](stdlib/04-interop/main.go) •
[Strings](stdlib/11-strings/main.go) •
[Time](stdlib/12-time/main.go) •
[Random numbers](stdlib/13-rand/main.go) •
[Number parsing](stdlib/14-strconv/main.go) •
[Reading files](stdlib/15-read-files/main.go) •
[Writing files](stdlib/16-write-files/main.go) •
[Scanning](stdlib/17-scanners/main.go) •
[File paths](stdlib/18-paths/main.go) •
[Directories](stdlib/19-dirs/main.go) •
[Temp files](stdlib/20-temp/main.go) •
[Command line](stdlib/21-flags/main.go) •
[Env variables](stdlib/22-env/main.go) •
[Logging](stdlib/23-slog/main.go)

## Apps

Coreutils:
[cat](apps/cat/main.go),
[cut](apps/cut/main.go),
[head](apps/head/main.go),
[sort](apps/sort/main.go),
[uniq](apps/uniq/main.go),
[wc](apps/wc/main.go)

[Count word frequencies](apps/wordfreq/main.go) by [serge-hulne](https://github.com/serge-hulne)

[Curl interop](apps/curl/main.go)

[FreeSWITCH module](https://github.com/rts-cn/mod_solod) by [seven1240](https://github.com/seven1240)

[Key-value store](apps/sqlmap/main.go) with SQLite interop

[Reverse lines in file](apps/reverse/main.go)

## Running the code

To run a specific example locally, use the `so run` command. For example:

```text
so run lang/05-for
```

```text
so run apps/head -n 4 data/jenny.txt
```

You'll need to have a C compiler installed and available as `cc`, or you can set a custom compiler by using the `CC` environment variable.

To see the generated C code, use the `so translate` command. For example:

```text
so translate -o lang/05-for/generated lang/05-for
```

Partially based on [Go by Example](https://gobyexample.com) by Mark McGranaghan, licensed under [CC BY 3.0](https://creativecommons.org/licenses/by/3.0/).
