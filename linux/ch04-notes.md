## Help
What it is: most CLI tools print a quick-start guide with a help option
Commands:
- `tool --help` → quick help (most common)
- `tool -h` → short form (works on many tools)
- `tool help` → some tools use this (e.g. `go help`, `git help`)
- `man tool` → the full manual (q to quit)
- use man grep in mac instead of grep -help

## Flags
What it is: options that change how a command behaves
Conventions:
- single letter → one dash: `-a`
- word → two dashes: `--help`
- combine single-letter flags: `ls -al` = `ls -a -l`
Commands:
- `ls -l` → long listing
- `ls -a` → show hidden files
- `wc -c file` → byte count
- `wc -l file` → line count
- `wc -w file` → word count
Gotcha: `-c` counts bytes, `-m` counts characters (same as Go's len() vs RuneCount)


## Positional Arguments
What it is: inputs to a command, where ORDER matters (like function arguments)
Commands:
- `cd /path` → 1 argument
- `mv FROM TO` → 2 arguments: move or rename
- `mv file.txt folder/` → move, keep the same name
- `mv old.txt new.txt` → rename
- `mv -i FROM TO` → ask before overwriting
Tried in my sandbox:
- moved key.txt from public/ to private/ and checked with ls
Gotcha: mv overwrites existing files silently. A typo in the destination name renames the file.

## Exit Codes
What it is: a number a program returns when it finishes. 0 = success, anything else = error (usually 1)
Commands:
- `echo $?` → exit code of the LAST command
- `cmd1 ; cmd2` → always run both
- `cmd1 && cmd2` → run cmd2 only if cmd1 succeeded
- `cmd1 || cmd2` → run cmd2 only if cmd1 failed
Tried in my sandbox:
- warn.sh without env vars → exit code 1
- warn.sh with env vars → exit code 0
Gotcha: $? changes after every command, so check it immediately
Go link: os.Exit(1) in my CLI tools sets the exit code

## Standard Error
What it is: a second output stream just for errors (stdout = 1, stderr = 2)
Commands:
- `cmd > file` → stdout to file
- `cmd 2> file` → stderr to file
- `cmd > out 2> err` → split them
- `cmd > all 2>&1` → both to one file
- `cmd 2> /dev/null` → discard errors
- `echo "msg" >&2` → (in a script) print to stderr

## Standard Input
What it is: the stream a program reads input from (keyboard by default). stdin = 0
Commands:
- `read -r var` → (in a script) read one line from stdin into var
- `program < file.txt` → feed a file into the program's stdin
- `wc file` vs `wc < file` → argument (program opens the file) vs stdin (shell feeds the data)

Gotcha: with `<`, the program never sees the filename, only the contents


## Piping
What it is: `|` sends one program's stdout into the next program's stdin
Commands:
- `cmd1 | cmd2` → pipe output into another command
- `grep -R "text" dir` → search all files in a folder recursively
- `--exclude-dir="name"` → skip folders with that name
- `... | wc -l` → count matching lines
- `cut -d, -f2 file | sort | uniq -c` → count occurrences per value
Tried in my sandbox:
- Bob count with exclude = 100, without = 150 (backups double-counted)
Gotcha: forgetting --exclude-dir counts duplicates. Check what folders you're searching
Pipe vs redirect: `>` sends output to a file, `|` sends it to a program

## Unix Philosophy
1. Do one thing and do it well (grep searches, wc counts, sort sorts)
2. Make programs work together (combine them with pipes)
3. Use text streams as the universal interface (everything reads/writes text)

When you build your projects, follow these so they work nicely with pipes:

Print results to stdout and errors to stderr.
Return exit code 0 on success and 1 on failure.
Print plain text (one result per line), so other tools can grep, sort or count it.
Later, accept input from stdin too, so you can do cat data.txt | yourtool.

**Why it matters:** small focused tools + a shared format = you can combine them in ways nobody planned
**In my Go code:** small single-purpose functions, logic separate from main, stdout for results, stderr for errors, exit codes
**System design link:** microservices = the same idea at a bigger scale

