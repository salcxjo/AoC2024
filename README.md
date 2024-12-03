# Advent of Code 2024

This repository includes my code for the [Advent of Code 2024](https://adventofcode.com/2024/about).


*If you want to solve the puzzles yourself first, do so, **then** look for the solution to compare.*
Do **not** cheat yourself out of a great time and learning experience by just copying code!

I'll update this repository now and then to include all my solutions whenever I've got time to properly comment and upload it.

This is mostly for those that are curious.
If you want to look at the code, experiment with it, change it, etc. be my guest.

## Running from command line

To run the code for any given day, use the following line (replacing `1` with the number of the desired day):

```bash
go run ./cmd/day01
```

As an alternative, you can also compile the executables explicitly, using `build` rather than `run`:

```bash
go build ./cmd/day01
./day01
```

Inputs are expected to be in the `data` sub folder relative to the current working directory.

## Running/debugging in VS Code

I'm using [Visual Studio Code](https://code.visualstudio.com/) for writing and debugging.
This should work out of the box, if you open the repository directory directly.

You should be able to pick the source code file of any day and simply hit <kbd>F5</kbd> to run/debug it.
A matching `launch.json` configuration file is included.
