# Sanitize
=====================

[![Go Report Card][goreport-svg]][goreport-link]
[![License][license-svg]][license-link]

Super simple function to clean up string from regional characters:

Sławomir Jasiński -> Slawomir Jasinski

##Usage

    package main

    import (
        "github.com/slav123/sanitize"
        "fmt"
    )

    func main() {
        fmt.Println(sanitize.Accents("gąska"))
    }




gives you

    gaska

Checkout on
    The [Go Playground](https://play.golang.org/p/wdAz8sxYx09)

[goreport-svg]: https://goreportcard.com/badge/github.com/slav123/sanitize
[goreport-link]: https://goreportcard.com/report/github.com/slav123/sanitize
[license-svg]: https://img.shields.io/github/license/SLAV123/sanitize
[license-link]: https://github.com/slav123/sanitize/blob/main/LICENSE
