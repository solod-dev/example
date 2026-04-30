# So by example

A hands-on introduction to [Solod](https://github.com/solod-dev/solod) (So) — a subset of Go that translates to C.

This repo contains small [examples](#examples) that show individual features, as well as larger [apps](#apps) that demonstrate how everything works together.

## Examples

[Hello world](./bits/01-hello-world/main.go) •
[Values](./bits/02-values/main.go) •
[Variables](./bits/03-variables/main.go) •
[Constants](./bits/04-constants/main.go) •
[For](./bits/05-for/main.go) •
[If/else](./bits/06-if-else/main.go) •
[Arrays](./bits/08-arrays/main.go) •
[Slices](./bits/09-slices/main.go) •
[Maps](./bits/10-maps/main.go) •
[Functions](./bits/11-functions/main.go) •
[Multiple returns](./bits/12-returns/main.go) •
[Variadic functions](./bits/13-variadics/main.go) •
[Recursion](./bits/15-recursion/main.go) •
[For-range](./bits/16-range/main.go) •
[Pointers](./bits/17-pointers/main.go) •
[Strings and runes](./bits/18-strings/main.go) •
[Structs](./bits/19-structs/main.go) •
[Methods](./bits/20-methods/main.go) •
[Interfaces](./bits/21-interfaces/main.go) •
[Enums](./bits/22-enums/main.go) •
[Errors](./bits/26-errors/main.go) •
[Panic](./bits/27-panic/main.go) •
[Defer](./bits/28-defer/main.go) •
[Memory](./bits/51-memory/main.go) •
[C interop](./bits/52-interop/main.go) •
[Strings](./bits/53-strings/main.go) •
[Time](./bits/54-time/main.go) •
[Files](./bits/55-files/main.go)

## Apps

[Count word frequencies](./apps/wordfreq/main.go) by @serge-hulne

[Curl interop](./apps/curl/main.go)

[FreeSWITCH module](https://github.com/rts-cn/mod_solod) by @seven1240

[Reverse lines in file](./apps/reverse/main.go)

## Running the code

To run a specific example locally, use the `so run` command. For example:

```text
so run bits/05-for
```

You'll need to have a C compiler installed and available as `cc`, or you can set a custom compiler by using the `CC` environment variable.

To see the generated C code, use the `so translate` command. For example:

```text
so translate -o bits/05-for/generated bits/05-for
```

Partially based on [Go by Example](https://gobyexample.com) by Mark McGranaghan, licensed under [CC BY 3.0](https://creativecommons.org/licenses/by/3.0/).
