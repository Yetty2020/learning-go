## Users
Every person on a Unix-like system (Linux, and macOS too) has their own account. Each account has a home folder (~), its own files, and its own permissions. The user in rwx from earlier means “the owner’s account.

Root, the superuser: there’s one special account, root, that can do anything: read every file, change system settings, delete anything. No permission checks apply to it.

sudo = “superuser do”: run this one command as root. It’s like borrowing the master key for a single job, then handing it back.

whoami          # your username
sudo whoami     # asks for your Mac password, then prints: root

You may have seen people type sudo npm install -g something when they get a permission error.

## Permissions
Permissions control who can do what to which files and directories.
The permissions of an individual file or directory are visually represented as a 10-character string:
drwxrwxrwx
-: Regular file (e.g. -rwxrwxrwx)
d: Directory (e.g. drwxrwxrwx)
rwx: All permissions

rw-: Read and write, but not execute

r-x: Read and execute, but not write

Each letter has an octal representation
r - 4
w - 2
x - 1 
- - 0

 rwx    r-x    r-x
4+2+1  4+0+1  4+0+1
  7      5      5     →  755

  0755 - the 0 in front just shows that it is an octal number, it can be background.




## Changing permissions
This is done with chmod command which means to change mode.
chmod  -R   u=rwx,g=,o=   DIRECTORY
│      │    │             └── what to change
│      │    └── the new permissions
│      └── also apply to everything inside
└── "change mode"

## Making a script executable
You'll frequently write your own scripts, or download them from the internet, only to find that they refuse to run.
You get "Permission Denied"

+x and -x: add or remove one switch
= replaced the permissions completely. + and - only flip one switch and leave the rest alone.
With no u, g or o in front, chmod +x applies to everyone. You can target one group with u+x.


chmod -x file_name


## Changing ownership
This is when you need sudo
sudo  chown  -R   root   contacts
│     │      │    │      └── target
│     │      │    └── new owner
│     │      └── everything inside too
│     └── change owner
└── as root (the safety net is off)

The chown command, which stands for "change owner," allows you to change the owner of a file or directory, and it requires root privileges.

You can use this if you want only the root to have access to a file