# Advent of Code 2016 in Go

These are my soluions for [Advent of Code 2016](https://adventofcode.com/2016/) written in Go.

I've solved them in Coffeescript one year ago. But since I'm learning Go at the moment, and have just finished writing solutions for puzzles for years 2017 and 2018, I've decided to continue with 2016 puzzles, refactoring and cleaning up solutions as much as possible.

Goals:

* implement generic code, that can come in handy during programming contests;
* same as for years 2017/2018: learn & document quirks & tricks of Go, which are new to me.

## "Go gotchas"

Go is low-level language with built-in concurrency and garbage collection, designed as a highly efficient C++ or Java competitor. To achieve high speed (both in compilation and execution), some surprising design decisions were made. It takes time to learn them.

For list of quirks found previously, refer to [README of my Go solutions for year 2018](https://github.com/metalim/metalim.adventofcode.2018.go/blob/master/README.md#go-gotchas) and [README of my Go solutions for year 2017](https://github.com/metalim/metalim.adventofcode.2017.go/blob/master/README.md#go-gotchas).

## Puzzle inputs

Inputs are automatically retrieved from Advent of Code, provided you put at least one `<session-name>.cookie` into `inputs/` folder. To get the cookie, refer to website properties in your browser, after logging in into Advent of Code website.

## Log

Check out [LOG.md](LOG.md) for specifics of each task.
