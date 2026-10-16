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