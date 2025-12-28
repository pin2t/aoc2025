![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white)![Java](https://img.shields.io/badge/java-%23ED8B00.svg?style=flat&logo=openjdk&logoColor=white)

# Advent Of Code 2025

This repository contains Advent Of Code 2025 solutions https://adventofcode.com/2025

## How to run

```shell
go run 01.go < inputs/01.txt
```
```shell
java day01.java < inputs/01.txt
```
Java 25+ is required
       
### Day 10 specifics

Day 10 solution requires z3 library. Unpack suitable zip archive from https://github.com/Z3Prover/z3/releases to z3 directory, then copy bin directory to go-z3 module
```shell
cp z3/bin/* /Users/pin/go/pkg/mod/github.com/mitchellh/go-z3@v0.0.0-20191228203228-4cbedeba863f/
```
and run
```shell
CGO_CFLAGS="-I/Users/pin/aoc2025/z3/include" go run 10.go < inputs/10.txt
```
use your full path to z3 of course

Java day 10 

```shell
cd z3/bin
java -cp com.microsoft.z3.jar: ../../day10.java < ../../inputs/10.txt

```