# go-sail

**go-sail** is a powerful CLI tool written in Go using the Cobra library. It simplifies the process of setting up Go backend frameworks like Fiber, Echo, and Gin by generating templates with built-in logging and caching features. Designed for both beginners and seasoned developers, **go-sail** reduces the time spent on repetitive coding tasks during project initialization.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Example](#example)
- [Contributing](#contributing)
- [References](#references)

## Features

- Generates templates for popular Go frameworks: Fiber, Echo, Gin.
- Make your own configuration with framework, databae and orm
- Integrated logging and caching setups.
- Reduces project setup time.
- Eliminates repetitive code for initializing Go projects.

## Installation

To install **go-sail**, follow these simple steps:

1. **Install Go**: Make sure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

2. **Clone the Repository**: Open your terminal and run the following command to clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-sail.git
3. Navigate to the Project Directory:
   ```bash
   cd go-sail
4. Install dependencies
   ```bash
   go mod tidy
5. Build and Install
   ```bash
   go build
   go install

## Usage
Once installed, you can start using go-sail to generate templates for your Go projects.
1. To check if the installation was successful, run:
   ```bash
   go-sail help

## Commands
1.  Initialize a new Go project with a chosen framework, database and orm configuration.
    ```bash
    go-sail create my-app

## Example
![image](https://github.com/user-attachments/assets/2e363533-9637-4a52-a784-ddb0aa338911)
![image](https://github.com/user-attachments/assets/4dd8542f-dc5c-4fce-bddf-82c6cf483368)
![image](https://github.com/user-attachments/assets/1797d83a-1993-430a-80c2-207d9ad52a7d)
![image](https://github.com/user-attachments/assets/9abcea55-2b5a-4c74-90cc-914f03e6dbf2)

## Contributing
Please follow these steps to contribute
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with a clear commit message.
4. Push your changes to your fork.
5. Submit a pull request to the dev branch of repository.

### Commit Message Guidelines
We follow the Conventional Commits specification for our commit messages. This leads to more readable messages that are easy to follow when looking through the project history. Here are some of the most common types of commits:

- `feat`: A new feature for the user or a significant addition to the codebase
  - Example: `feat: add support for Gin framework`

- `fix`: A bug fix
  - Example: `fix: resolve issue with Echo template generation`

- `chore`: Regular maintenance tasks, updates to build processes, or other changes that don't modify src or test files
  - Example: `chore: update dependencies to latest versions`

- `docs`: Documentation only changes
  - Example: `docs: update installation instructions in README`

- `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
  - Example: `style: format code according to Go standards`

- `refactor`: A code change that neither fixes a bug nor adds a feature
  - Example: `refactor: simplify template generation logic`

- `test`: Adding missing tests or correcting existing tests
  - Example: `test: add unit tests for Fiber template generation`

## References
1. https://github.com/spf13/cobra

