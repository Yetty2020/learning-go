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