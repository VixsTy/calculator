# Calculator

A simple command line calculator.

<p align="center" style="margin-bottom: 20px; width: 100px; height: 100px; margin: auto">
<img src="https://golang.org/doc/gopher/fiveyears.jpg" width="250px"/>
</p>

## Working operators

- `+` for the addition
- `-` for the substraction
- `*` for the multiplication
- `/` for the division
- `%` for the modulus
- `^` for the exponent

## TODO List

- [ ] Add function management (sin, cos, tan, cot, sec, csc, asin, acos, atan, acot, asec, acsc, sqrt, log, lg, ln, abs)
- [ ] Add constants (e, pi, π)
- [ ] Manage the negative number
- [ ] Work on cyclomatic complexity
- [ ] Expression validation to have more clear error
- [ ] Manage multiple number basis and enable calcul in base2, base10 and base16
- [ ] Infinite number by using string calculation

## Usage

### Go get

```
$ go get github.com/VixsTy/calculator/cli/calculator
// If your $GOPATH/bin is part of your $PATH
$ calculator --help
```

### Docker

```
$ docker run --rm -it vixsty/calculator --help
```

## Documentation

- Main algorithm [Shunting-yard](https://en.wikipedia.org/wiki/Shunting-yard_algorithm)
- Inspiration [alfredxing/calc](https://github.com/alfredxing/calc)
- Inspiration [mgenware/go-shunting-yard](https://github.com/mgenware/go-shunting-yard)
