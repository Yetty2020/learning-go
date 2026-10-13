## Help
What it is: most CLI tools print a quick-start guide with a help option
Commands:
- `tool --help` → quick help (most common)
- `tool -h` → short form (works on many tools)
- `tool help` → some tools use this (e.g. `go help`, `git help`)
- `man tool` → the full manual (q to quit)
- use man grep in mac instead of grep -help

## Flags
**What it is:** options that change how a command behaves
**Conventions:**
- single letter → one dash: `-a`
- word → two dashes: `--help`
- combine single-letter flags: `ls -al` = `ls -a -l`
**Commands:**
- `ls -l` → long listing
- `ls -a` → show hidden files
- `wc -c file` → byte count
- `wc -l file` → line count
- `wc -w file` → word count
**Gotcha:** `-c` counts bytes, `-m` counts characters (same as Go's len() vs RuneCount)
