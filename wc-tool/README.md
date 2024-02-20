# Write your own wc tool

This challenge corresponds to the first part of the Coding Challenges series by John Crickett:  [coding challenges](https://codingchallenges.fyi/challenges/challenge-wc)

## Description

A command-line tool inspired by the classic wc utility with a coding challenges twist. This Go program, ccwc, allows you to count words, lines, characters, and bytes in a file or from standard input. Perfect for enhancing your coding skills through practical challenges!


## Usage

1. Count Bytes

```bash
ccwc -c test.txt
```
- Output

```bash
    342190 test.txt
```

2. Count Lines

```bash
ccwc -l test.txt
```
- Output

```bash
    7145 test.txt
```

3. Count Words

```bash
ccwc -w test.txt
```
- Output

```bash
    58164 test.txt
```

4. Count Characters

```bash
ccwc -m test.txt
```
- Output (result may vary based on locale):

```bash
    339292 test.txt

```

5. Default Count:

```bash
ccwc test.txt
```
- Output

```bash
    7145   58164  342190 test.txt
```


6. Read from Standard Input:

```bash
cat test.txt | ccwc -l
```
- Output

```bash
    7145
```

## Getting Started

Follow the steps in the challenge description to set up your development environment and start coding! Test your solution with the provided test.txt file and enjoy the journey of building your own wc tool.