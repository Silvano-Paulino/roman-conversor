# Roman Numeral Converter

This project provides a command-line interface (CLI) tool for converting between Roman numerals and decimal numbers. It includes two main functionalities: converting integers to Roman numerals and converting Roman numerals to integers. The project is implemented in Go and is designed to be simple, efficient, and easy to use.

---

## Features

- Convert integers to Roman numerals.
- Convert Roman numerals to integers.
- Error handling for invalid inputs.
- Cross-platform support (Linux and Windows).

---

## Installation

### Prerequisites

- [Go](https://golang.org/dl/)

### Steps

1. Clone the repository:
    ```bash
    git clone https://github.com/silvano-paulino/conversor-roman.git
    cd conversor-roman
    ```

2. Build the CLI tool for your platform:

    - **Linux**:
      ```bash
      make build-linux
      ```

    - **Windows**:
      ```bash
      make build-windows
      ```

3. The executable file (`cli` for Linux or `cli.exe` for Windows) will be generated in the project root directory.

---

## Usage

The CLI tool supports two subcommands: `int` and `roman`.

### Convert Integer to Roman Numeral

```bash
./cli int --value=<integer>
```

**Example**:
```bash
./cli int --value=2025
```

**Output**:
```
Decimal to Roman: 2025 -> MMXXV
```

### Convert Roman Numeral to Integer

```bash
./cli roman --value=<roman_numeral>
```

**Example**:
```bash
./cli roman --value=MMXXV
```

**Output**:
```
Roman to Decimal: MMXXV -> 2025
```

---

## Running Tests

To ensure the correctness of the implementation, you can run the test suite:

```bash
make tests
```

---

## Project Structure

- `internal/converter/`: Contains the core logic for converting between Roman numerals and integers.
- `internal/main.go`: Entry point for the CLI tool.

---

## Error Handling

- If the input value is invalid or missing, the tool will display an appropriate error message and exit.

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
