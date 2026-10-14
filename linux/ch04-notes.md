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