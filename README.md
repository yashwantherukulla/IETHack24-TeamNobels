# IETHack24-TeamNobels

Welcome to the **IETHack24-TeamNobels** repository! This project is a comprehensive suite of tools and services designed to streamline the development and deployment of Go-based backend applications. The project is divided into three main components:

1. **go-sail**: A powerful CLI tool for generating Go backend project templates with built-in logging and caching features.
2. **go-sail-ML**: A machine learning-based code analysis tool that evaluates code quality, security, and provides detailed descriptions.
3. **go-sail-backend**: A backend service that handles user authentication, API key management, and plan upgrades.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Example](#example)
- [Contributing](#contributing)
- [References](#references)

## Features

### go-sail
- Generates templates for popular Go frameworks: Fiber, Echo, Gin.
- Integrated logging and caching setups.
- Reduces project setup time.
- Eliminates repetitive code for initializing Go projects.

### go-sail-ML
- Analyzes code quality, security, and provides detailed descriptions.
- Supports analysis of individual files, folders, and entire repositories.
- Generates comprehensive reports with scores and improvement suggestions.

### go-sail-backend
- Handles user authentication and API key management.
- Supports plan upgrades with different levels of access and features.
- Secures API endpoints with JWT-based authentication.

## Installation

### go-sail
1. **Install Go**: Make sure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).
2. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/go-sail.git
   ```
3. **Navigate to the Project Directory**:
   ```bash
   cd go-sail
   ```
4. **Install Dependencies**:
   ```bash
   go mod tidy
   ```
5. **Build and Install**:
   ```bash
   go build
   go install
   ```

### go-sail-ML
1. **Install Python**: Ensure you have Python 3.9 or later installed.
2. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/go-sail-ML.git
   ```
3. **Navigate to the Project Directory**:
   ```bash
   cd go-sail-ML
   ```
4. **Install Dependencies**:
   ```bash
   pip install -r requirements.txt
   ```

### go-sail-backend
1. **Install Go**: Ensure you have Go installed.
2. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/go-sail-backend.git
   ```
3. **Navigate to the Project Directory**:
   ```bash
   cd go-sail-backend
   ```
4. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

## Usage

### go-sail
Once installed, you can start using `go-sail` to generate templates for your Go projects.

1. **Initialize a new Go project**:
   ```bash
   go-sail create my-app
   ```

### go-sail-ML
To analyze code quality and security:

1. **Run the ML analysis**:
   ```bash
   python src/eval/code_analyser.py
   ```

### go-sail-backend
To start the backend service:

1. **Run the server**:
   ```bash
   go run main.go
   ```

## Commands

### go-sail
- **Create a new project**:
  ```bash
  go-sail create my-app
  ```
- **Evaluate a project**:
  ```bash
  go-sail evaluate --file path/to/file
  ```

### go-sail-ML
- **Analyze a repository**:
  ```bash
  python src/eval/code_analyser.py
  ```

### go-sail-backend
- **Start the server**:
  ```bash
  go run main.go
  ```

## Example

### go-sail
![Example](https://github.com/user-attachments/assets/2e363533-9637-4a52-a784-ddb0aa338911)

### go-sail-ML
![Example](https://github.com/user-attachments/assets/4dd8542f-dc5c-4fce-bddf-82c6cf483368)

### go-sail-backend
![Example](https://github.com/user-attachments/assets/1797d83a-1993-430a-80c2-207d9ad52a7d)

## Contributing

Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with a clear commit message.
4. Push your changes to your fork.
5. Submit a pull request to the `dev` branch of the repository.

### Commit Message Guidelines
We follow the Conventional Commits specification for our commit messages. Here are some of the most common types of commits:

- `feat`: A new feature for the user or a significant addition to the codebase.
- `fix`: A bug fix.
- `chore`: Regular maintenance tasks, updates to build processes, or other changes that don't modify src or test files.
- `docs`: Documentation only changes.
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc).
- `refactor`: A code change that neither fixes a bug nor adds a feature.
- `test`: Adding missing tests or correcting existing tests.

## References

1. [Cobra Library](https://github.com/spf13/cobra)
2. [Gin Framework](https://github.com/gin-gonic/gin)
3. [GORM](https://gorm.io/)
4. [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

---

Thank you for using **IETHack24-TeamNobels**! We hope this toolset makes your Go development process smoother and more efficient. If you have any questions or feedback, please feel free to open an issue or reach out to us directly. Happy coding! 🚀
