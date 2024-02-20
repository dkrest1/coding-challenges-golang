# Write your own wc tool

This challenge corresponds to the first part of the Coding Challenges series by John Crickett:  [coding challenges](https://codingchallenges.fyi/challenges/challenge-wc)

## Description

A command-line tool inspired by the classic wc utility with a coding challenges twist. This Go program, `wc-tool`, allows you to count words, lines, characters, and bytes in a file or from standard input. Perfect for enhancing your coding skills through practical challenges!

## Getting Started

Follow the steps in the challenge description to set up your development environment and start coding! 

## Usage

**PS:** You can run the the program directly using `go run main.go --options[l, m, w, c]` or you build the program running `make build` to create the executable, it will be saved in the `bin/` directory.

1. Count Bytes

<!-- Directly -->

```bash
go run main.go -c test.txt 
```

<!-- Executable -->

```bash
./bin/ccwc -c test.txt
```

- Output

```bash
    335039 test.txt
```

2. Count Lines

<!-- Directly -->

```bash
go run main.go -l test.txt 
```

<!-- Executable -->

```bash
./bin/ccwc -l test.txt
```

- Output

```bash
    7143 test.txt
```

3. Count Words

<!-- Directly -->

```bash
go run main.go -w test.txt 
```

<!-- Executable -->

```bash
./bin/ccwc -w test.txt
```

- Output

```bash
    58164 test.txt
```

4. Count Characters

<!-- Directly -->

```bash
go run main.go -m test.txt 
```

<!-- Executable -->

```bash
./bin/ccwc -m test.txt
```

- Output (result may vary based on locale):

```bash
    327897 test.txt

```

5. Default Count:

<!-- Directly -->

```bash
go run main.go test.txt 
```

<!-- Executable -->

```bash
./bin/ccwc test.txt
```

- Output

```bash
    335039 test.txt
    7143 test.txt
    58164 test.txt
```

6. Read from Standard Input:

```bash
 cat test.txt | go run main.go -l
```
- Output

```bash
    7143
```

## Run Test

To test the `wc` tool, follow these steps:

1. **Build the ccwc executable:**

    ```bash
    make build
    ```

2. **Run Tests:**

    ```bash
    make test
    ```

    This will execute the tests and provide verbose output.

3. **Clean Up:**

    Optionally, you can clean up the generated binary:

    ```bash
    make clean
    ```