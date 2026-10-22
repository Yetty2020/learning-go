## Package Managers in Linux

Brew (Homebrew) is like an app store for the terminal.
It is a package manager software tool that allows you to install other software. It goes to a system folder and is available everywhere on your system.

npm is also an example of a software manager. It for installing JS libraries from the terminal. It goes it node/modules in your project.

Basically,You never have to go to a website, download a file and set it up yourself. In TS terms, it’s npm for your whole Mac instead of for one project.

Neovim is a text editor that runs inside the terminal. It’s like VS Code, but with no windows or mouse. You edit files using only the keyboard, right where you type commands.
Neovim works the way nano works but is more complex.

nano is like editing a file in Notepad, while a configured Neovim can work like VS Code with the TS extension (autocomplete, errors, go-to-definition), all inside the terminal.

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
This is to install brew in the terminal

brew install neovim
nvim --version

This is to install neovim and check the version


## Using Neovim

To open an existing file
nvim pathName

When you open this, there are usually two modes: normal mode and insert mode.
Normal mode is the default mode and this is used for commands.
When you want to switch to insert mode, there are keys used to do this:
Key	Enters Insert mode...
i	before the cursor
a	after the cursor
o	on a new line below
s	after deleting one character
c (+ a motion)	after deleting some text

How to tell which mode you’re in: look at the bottom-left of the screen.

-- INSERT -- means Insert mode
blank means Normal mode

The safe reset: press Esc once or twice. That always puts you back in Normal mode.

If you think a stray keypress changed the file: press Esc, type :q!, and press Enter. That quits without saving, so the file is untouched. Then open it again with the file path:

Cheat sheet for this lesson

i	go into Insert mode (start typing)
Esc	back to Normal mode
:w + Enter	write (save)
:q + Enter	quit
:wq + Enter	save and quit in one go
:q! + Enter	quit without saving (the emergency exit if you mess up)