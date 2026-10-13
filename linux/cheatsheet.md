## Files & folders
| Command | What it does |
|---|---|
| `mkdir -p a/b/c` | create a folder and all missing parent folders |
| `touch file.txt` | create an empty file (won't create folders) |
| `cat file.txt` | show a file's contents |
| `nano file.txt` | edit a file (Ctrl+O save, Ctrl+X exit) |

## Writing to files
| `echo "hi" > f.txt` | write to a file (overwrites!) |
| `echo "hi" >> f.txt` | add to the end of a file |
| `cat > f.txt << 'EOF' ... EOF` | write many lines (quoted EOF = no $ expansion) |

## Environment variables
| `export NAME="x"` | set a variable that programs can see |
| `unset NAME` | remove it |
| `NAME="x" command` | set it for one command only |
| `env` | list all environment variables |

## PATH
| `echo $PATH` | show the folders the shell searches for commands |
| `export PATH="$PATH:$HOME/folder"` | ADD a folder (keep `$PATH:`!) |

**Gotchas:**
- `/something` starts from the system root, `something` is relative to where you are
- forgetting `$PATH:` replaces your whole PATH and breaks `ls`, `cat`, etc.