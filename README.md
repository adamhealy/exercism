# exercism
A place for my [exercism](https://exercism.io/profiles/adamhealy) exercises.


## Track Go

### Main track
- Hello World:
    - [Exercism](https://exercism.io/tracks/go/exercises/hello-world/solutions/ed878512cf97408d966c6b8be301c5a4) solution
    - [GitHub](https://github.com/adamhealy/exercism/blob/main/go/hello-world/hello_world.go) solution
- Two Fer:
    - [Exercism](https://exercism.io/tracks/go/exercises/two-fer/solutions/a2b62369f69b4e6bbbbb4d7866e4499a) solution
    - [GitHub](https://github.com/adamhealy/exercism/blob/main/go/two-fer/two_fer.go) solution
- Hamming: So much to learn from feedback here (hence the list)
    - Lessons:
        - Strings / Runes / slices
            - [Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)
            - [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
        - Errors return the [zero value](https://golang.org/ref/spec#The_zero_value)
        - [Go proverbs](https://go-proverbs.github.io/) (zen of go)
        - No blank lines in functions
        - for range to iterate over slice ([Go by Example: Range](https://gobyexample.com/range))
    - [Exercism](https://exercism.io/tracks/go/exercises/hamming/solutions/57febc1f040649f1af34970e7313faeb) solution
    - [GitHub](https://github.com/adamhealy/exercism/blob/main/go/hamming/hamming.go) solution
- Raindrops:
    - Lessons:
        - To declare an empty variable, use `var s string` instead of `s := ""`. Reason as explain by the mentor:
            > This is because the first approach doesn't allocate any memory until it's assigned a value. While the other approach assigns it a value, which in this case is the default value.
    - [Exercism](https://exercism.io/tracks/go/exercises/raindrops/solutions/6b0e70d3bda2476f95a71b72e0821c03) solution
    - [GitHub](https://github.com/adamhealy/exercism/blob/main/go/raindrops/raindrops.go) solution



### Extra exercises
-  Space Age: I learnt from the community solutions on this one, specifically about [maps](https://blog.golang.org/maps)
from corbyshaner's [solution](https://exercism.io/tracks/go/exercises/space-age/solutions/9ab75f550d924657ba27f12c3f539e7b). 
    - [Exercism](https://exercism.io/tracks/go/exercises/space-age/solutions/ea11a36254854236a9a7c8eaae3f3884) solution
    - [GitHub](https://github.com/adamhealy/exercism/blob/main/go/space-age/space_age.go) solution
- Leap: I could probably condense conditionals into single statement
    - [Exercism](https://exercism.io/tracks/go/exercises/leap/solutions/50c55031e6354a2d8e35af31a0458ce0) solution
    - [Github](https://github.com/adamhealy/exercism/blob/main/go/leap/leap.go) solution



## Resources

- [Buf](https://buf.build/docs/introduction) - cli tool for working with Protobuf's
- [Gophercises](https://courses.calhoun.io/courses/cor_gophercises) - A series of exercise to learn go
- [golang-migrate](https://github.com/golang-migrate/migrate) - Go Migrate cli tool & go library
- [pgx](https://github.com/jackc/pgx) - PostgreSQL Driver and Toolkit
- [pq](https://github.com/lib/pq) - A pure Go postgres driver for Go's database/sql package
