## Open a Terminal
What it is: the app that gives you a shell on your own machine (not the browser sandbox)
My setup: Ghostty on macOS (Terminal.app works the same, since both run zsh/bash)
Tried: `echo "I love the command line"` → printed it back

## Download Worldbanc
Commands:
- `curl -L URL -o file` → download a file (-L follow redirects, -o output name)
- `unzip file.zip` → extract a zip
- `rm file` → delete permanently (no Trash!)
- `mv -n old new` → rename without overwriting
Did: renamed my practice sandbox to worldbanc-practice, downloaded the real worldbanc
Gotcha: mv into an existing folder name moves INSIDE it instead of renaming

## Permanent PATH
What it is: putting `export PATH=...` in ~/.zshrc so it applies to every new shell
Commands:
- `export PATH="$PATH:$HOME/worldbanc/private/bin"` → added to the end of ~/.zshrc
- `source ~/.zshrc` or `. ~/.zshrc` → reload without restarting
- `which tool` → show where the shell finds a command
Gotchas:
- always keep `$PATH:` (or you replace your whole PATH)
- use double quotes so $PATH expands, not single quotes
- use an absolute path ($HOME/...), not a relative one

## Shell Aliases
What it is: a shortcut name for a command
Commands:
- `alias ll="ls -la"` → create (temporary, this session only)
- `alias` → list all aliases
- `alias ll` → show one alias
- `unalias ll` → remove it
- put the alias line in ~/.zshrc → permanent
My aliases: `reload` (source ~/.zshrc), `gsave` (git add + commit)
Gotcha: zsh shows aliases as `name='cmd'`, bash as `alias name='cmd'`

## Man Pages
What it is: the built-in manual for commands
Commands:
- `man grep` → open the manual
- `/text` → search, `n` next, `N` previous, `q` quit
- `Space` next page, `b` back, `G` end, `g` start
Tip: search for the flag (`/-r`) or jump to `/EXAMPLES`
Gotcha: macOS man pages are BSD versions and can differ from Linux

## Symbolic Links
What it is: a shortcut file that points to another file or folder (not a copy)
Commands:
- `ln -s TARGET LINK` → create a symlink (target first!)
- `ls -l` → shows `link -> target`, and the line starts with `l`
Did: moved tbills.txt to investments/, left a symlink in credit_cards/
Gotchas:
- a relative target is resolved from the LINK's folder, not where I ran ln
- deleting the target breaks the link
- absolute target: the link can move; relative target: link + target can move together

## top
What it is: live view of which processes use the most CPU/memory (like Activity Monitor)
Commands:
- `top` → start (sorted by CPU)
- `top -o mem` → start sorted by memory (macOS)
- inside top (macOS): `o` then `mem`/`cpu` to sort, `q` to quit
- inside top (Linux): `M` memory, `P` CPU, `q` quit
Key columns: PID, COMMAND, %CPU, MEM, TIME, STATE
Gotcha: narrow windows hide columns, so widen it and restart top

## Interrupt
What it is: stopping a running program
Commands:
- `Ctrl+C` → sends SIGINT (interrupt), most programs stop
- `echo $?` after Ctrl+C → usually 130
Use it when: a command hangs, has a typo, loops forever, or takes too long
Did: read malicious.sh with cat first, ran it, stopped it with Ctrl+C
Habit: always `cat` an unknown script before running it

