# Advent of Code 2025

This repo is a storage for my solutions to AoC '25 puzzles. All written in Go.

https://adventofcode.com/2025

## How to run 

```
> go run . <day number>
```

## Notes

### Day 10

Day 10 task solution for part 2 uses [golp](https://github.com/draffensperger/golp) library, therefore you need to follow installation guide in its readme in order to run it.

I installed it with homebrew, and then exported following env vars:

```
export CGO_CFLAGS="-I/opt/homebrew/opt/lp_solve/include"
export CGO_LDFLAGS="-L/opt/homebrew/opt/lp_solve/lib"
```