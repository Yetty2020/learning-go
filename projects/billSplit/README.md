# billsplit

A command-line tool, written in Go, that adds a tip to a bill and splits the total equally between a group of people.

## Usage

```
billsplit <bill> <people> <tip_percent>
```

```
$ billsplit 12000 4 10
You have a total bill of ₦12000.00, with a tip of 10.0 % and you are to split it amongst 4 people, which gives ₦3300.00 per person

$ billsplit 12000 0 10
need at least 1 person to split the bill

$ billsplit 12000 4 -5
tip cannot be negative

$ billsplit 12000 2.5 10
person has to be a whole number

$ billsplit abc 4 10
bill should be a number
```

Errors are printed to stderr and the program exits with code `1`.

## Rules

| Input  | Type         | Valid when        |
|--------|--------------|-------------------|
| bill   | decimal      | greater than 0    |
| people | whole number | 1 or more         |
| tip %  | decimal      | 0 or more         |

The tip is worked out as a percentage of the bill, added to the bill to get the total, and the total is split equally between everyone.

For example: ₦12000 + 10% tip (₦1200) = ₦13200, split between 4 people = ₦3300 each.

## Build and test

Requires Go.

```
go test -v                      # run the tests
go build -o ~/bin/billsplit .   # build and install
```

Make sure `~/bin` is on your `PATH` to run `billsplit` from any folder.

## How it works

- `calcTip(bill, tipPercent float64) (float64, error)` returns the tip amount, or an error if the bill is 0 or less or the tip is negative.
- `calcSplitBill(total float64, people int) (float64, error)` returns each person's share, or an error if the total is 0 or less or there are fewer than 1 person.
- `main()` reads the three arguments, converts them, adds the tip to the bill, and prints the result. All the maths lives in the two functions above, which are covered by table-driven tests.

## What I learned

- **Go doesn't mix number types.** `total / people` won't compile when one is a `float64` and the other is an `int`. I used `float64(people)` to convert it for the division.
- **Picking the right converter.** The bill and tip use `strconv.ParseFloat` because they can have decimals, while people uses `strconv.Atoi` because you can't split a bill between 2.5 people.
- **Testing one problem at a time.** I added test rows where only the total is invalid and the number of people is valid, so each guard clause is tested on its own and a missing check can't hide behind another one.
- **Passing tests don't mean the program is correct.** My first version printed ₦300 per person instead of ₦3300, because `main()` was splitting the tip instead of the total. Both functions were correct; the bug was in how `main()` connected them.
- **Checking a command name before using it.** I ran `which split` and found that `/usr/bin/split` already exists, so I called this tool `billsplit` to avoid a clash on my `PATH`.
- **Floats and money.** ₦10000 split 3 ways is ₦3333.333…, which can't be stored exactly, so I test with numbers that divide cleanly and only round when printing with `%.2f`.