# grade

A command line tool that converts 0 - 100 score into letter grade.

## Usage

```
grade <score>

$ grade 70
Score 70: grade A
$ grade -1
Number must be  0 to 100
$ grade 150
Number must be  0 to 100
$ grade ban
Input a valid number
```


## Grade scale

| Score    | Grade |
|----------|-------|
| 70 – 100 | A     |
| 60 – 69  | B     |
| 50 – 59  | C     |
| 45 – 49  | D     |
| 40 – 44  | E     |
| 0 – 39   | F     |


## Build and test

```
go test -v   # the command that runs the test

go build -o ~/bin/grade .    # the command that builds the tool

```

## What I learned
- I learned how to write tests first and watched all fail before writing any line of code.
- I learned os.Args is a slice of what the user types, where index 0 is the program’s own path and the score is at index 1.
- os.Args receives string, so I had to convert it with strconv.Atoi to an integer.
- I learned how to write each function first before calling them out in the main.
- I learned that main() reports errors to stderr and exits with code 1, while the logic stays in getGrade. 


