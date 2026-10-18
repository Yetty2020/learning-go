# dataPlanner

A command-line tool that shows how much of your mobile data plan you've used, how much is left, and your current status, written in Go.

## Usage

```
dataPlanner <plan_gb> <used_gb>
```

```
$ dataPlanner 10 7.5
You have used up 75.0 % of your data plan. You have 2.5GB left. Your current status is: OK

$ dataPlanner 10 15
You have used up 150.0 % of your data plan. You have 0.0GB left. Your current status is: Exhausted

$ dataPlanner 0 5
You do not have an active data plan

$ dataPlanner ban 4
Plan must be a number
```

## Status rules

| Used          | Status      |
|---------------|-------------|
| under 80%     | OK          |
| 80% to 99.9%  | Running low |
| 100% or more  | Exhausted   |

Remaining data never goes below 0. If you use more than your plan, remaining shows 0.0GB and the status is Exhausted.

## Build and test

```
go test -v                        # run the tests
go build -o ~/bin/dataPlanner .   # build and install
```

## How it works

- `getStatus(percent float64) string` takes the percentage used and returns `OK`, `Running low` or `Exhausted`.
- `calcUsage(plan, used float64) (float64, float64, error)` returns the data remaining and the percentage used. It returns an error if the plan is 0 or less, or if usage is negative.
- `main()` reads the two arguments, converts them with `strconv.ParseFloat`, calls both functions, and prints the result. Any error goes to stderr and the program exits with code 1.

Both `getStatus` and `calcUsage` have table-driven tests. `main()` only handles input and output.

## What I learned


- I learned to use `float64` and `strconv.ParseFloat` for decimal numbers, because `Atoi` only works with whole numbers.
- I learned to create my own errors with `errors.New` when the input breaks the rules, like a plan of 0 or negative plan.
- I learned there are two kinds of bad input: text that isn't a number (caught by `strconv`) and numbers that don't make sense (caught by my own checks).
- I learned never to ignore an error with `_`. When I did, my program said "0% used, status OK" for a negative plan.
- I learned to test errors by turning them into true/false with `gotError := err != nil`.
- I learned that once a guard clause rejects bad input at the top, I don't need to check it again lower down.