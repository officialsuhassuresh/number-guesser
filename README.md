# Number Guesser CLI

A command-line number guessing game written in Go where you try to guess a randomly generated number within a limited number of attempts.

As per the instructions in https://roadmap.sh/projects/number-guessing-game

## Prerequisites

- Go 1.22 or higher installed on your system
- Git (for cloning the repository)
- Your GOPATH should be properly configured

To check if Go is installed and GOPATH is configured:

```
go version
```

```
go env
```


## Installation

1. First, ensure your Go bin directory is in your PATH:

```
export PATH=$PATH:$(go env GOPATH)/bin
```

Add the following to your .bashrc or .zshrc file:

```
export PATH=$PATH:$(go env GOPATH)/bin
```

source the file to apply the changes:

```
source ~/.bashrc
```

2. Clone the repository:

```
git clone https://github.com/officialsuhassuresh/number-guesser.git
```

3. Navigate to the project directory:

```
cd number-guesser
``` 


Option 1:

```
go install github.com/officialsuhassuresh/number-guesser@latest
```

Option 2:

```
go install
```


## Usage

1. Start a new game:

```
number-guesser
```

2. Start a new game with a specific difficulty:

```
number-guesser play
```

3. View high scores:

```
number-guesser stats
```

4. Reset high scores:

```
number-guesser reset
``` 

5. View the help menu:

```
number-guesser --help       
```

6. View the version:

```
number-guesser --version
```

7. View the help menu for a specific command:

```
number-guesser play --help
```

## Game Rules

- The game generates a random number between 1 and 100.
- You have a limited number of attempts to guess the number.
- The difficulty level determines the number of attempts you have.
- The game will tell you if your guess is too high or too low.
- The game will tell you if you have guessed the number correctly.
- The game will keep track of your high scores.

## Features

- Interactive difficulty selection
- High scores tracking
- Reset high scores
- Help menu
- Version information
- Interactive difficulty selection
- High scores tracking
- Reset high scores
- Help menu

## Troubleshooting

- If you encounter any issues, please check the help menu for the specific command you are trying to use.

## License

This project is open-sourced under the MIT License - see the LICENSE file for details.