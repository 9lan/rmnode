# RMNODE - Remove All Node Modules

rmnode is a command-line tool written in Golang that helps you to clean up `node_modules` directories in your PC.

## Installation

1. Install Golang if it's not already installed on your system. You can download it from the official Golang website: https://golang.org/dl/

2. Clone the repository: `git clone https://github.com/9lan/rmnode.git`

3. Navigate to the project directory: `cd rmnode`

4. Build the executable: `go build`

5. Run the program: `./rmnode`

To run the `rmnode` program without having to write `./rmnode` every time, you can add the project's directory to your system's `$PATH` environment variable.

Here's how you can do it:

1. Open your terminal and navigate to the project's directory: `cd /path/to/rmnode`

2. Run the following command to get the absolute path of the current directory: `pwd`. Note down the output of this command, as you will need it in the next step.

3. Open your shell profile file. If you're using Bash, this file will be `~/.bashrc` or `~/.bash_profile`. If you're using Zsh, this file will be `~/.zshrc`. You can open the file using any text editor of your choice, for example:

```sh
nano ~/.bashrc
# or
nano ~/.zshrc
```

4. Add the following line to the end of the file, replacing /path/to/rmnode with the absolute path you noted down in step 2:

```sh
export PATH=$PATH:/path/to/rmnode
```

5. Save and close the file.

6. Restart your terminal or run the following command to apply the changes to your current session:

```sh
source ~/.bashrc
# or
source ~/.zshrc
```

Note: If you're using a different shell profile file than ~/.bashrc, replace it in the command above.

Now, you should be able to run the `rmnode` program from any directory by simply typing `rmnode` followed by any command-line arguments.

## Usage

```sh
rmnode [list -ls --list|delete -d --delete]
```

- `list`, `-ls`, `--list`: List all node_modules directories in your project.
- `delete`, `-d`, `--delete`: Delete all node_modules directories in your project.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on the project's GitHub repository. If you want to contribute code, please fork the repository and submit a pull request.

## License

rmnode is licensed under the MIT License. See the LICENSE file for more information.
