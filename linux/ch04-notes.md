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