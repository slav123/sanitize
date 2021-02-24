# Sanitize

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

    