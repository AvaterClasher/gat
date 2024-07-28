# Cat Command in Go

A Go implementation of the `cat` command used to concatenate file(s) and print the result to the terminal. This project replicates the functionality of the Unix `cat` command with additional features.

## Features

- Concatenate and display files.
- Read from standard input if no files are provided or `-` is used.
- Number all lines or non-blank lines.
- Suppress repeated empty lines.
- Display end-of-line characters.
- Show tab characters.

## Installation

### Prerequisites

- Go 1.18 or higher

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/AvaterClasher/gat
   cd gat
   ```

2. Build the project:
   ```sh
   go build -o cat.exe
   ```

3. Add the directory containing `cat.exe` to your PATH:

   #### Windows

   - Open the Start Search, type in "env", and select "Edit the system environment variables"
   - In the System Properties window, click on the "Environment Variables" button
   - In the Environment Variables window, highlight the "Path" variable in the "System variables" section and click the "Edit" button
   - Click "New" and add the path to the directory containing `cat.exe` (e.g., `C:\path\to\cat-go`)
   - Click "OK" to close all windows

4. Verify the installation:
   ```sh
   cat --version
   ```

## Usage

```sh
cat [OPTIONS] [FILE]...
```

### Options

- `-n, --number` : Number all output lines
- `-b, --number-nonblank` : Number non-blank output lines
- `-s, --squeeze-blank` : Suppress repeated empty output lines
- `-E, --show-ends` : Display `$` at end of each line
- `-T, --show-tabs` : Display TAB characters as `^I`

### Examples

1. **Basic Concatenation**:
   ```sh
   cat file1.txt file2.txt
   ```

3. **Number All Lines**:
   ```sh
   cat --number file1.txt
   cat -n file1.txt
   ```

4. **Number Non-Blank Lines**:
   ```sh
   cat --number-nonblank file1.txt
   cat -b file1.txt
   ```

5. **Suppress Repeated Empty Lines**:
   ```sh
   cat --squeeze-blank file1.txt
   cat -s file1.txt
   ```

6. **Display End-of-Line Characters**:
   ```sh
   cat --show-ends file1.txt
   cat -E file1.txt
   ```

7. **Show Tab Characters**:
   ```sh
   cat --show-tabs file1.txt
   cat -T file1.txt
   ```

8. **Combined Options**:
   ```sh
   cat --number --show-ends file1.txt
   cat -n -E file1.txt
   ```

9. **Reading from Multiple Files**:
   ```sh
   cat file1.txt file2.txt
   ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.