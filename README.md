# github-activity

A command-line interface (CLI) to fetch and display recent public activity for a given GitHub user.

## Features

- Fetches a user's public GitHub events.
- Displays a summary of each event, indicating the type of activity and the repository involved.

## Installation

To install the `github-activity` CLI, make sure you have Go installed (Go 1.22 or higher is recommended).

```bash
go install ./cmd
```

This will install the `github-activity` executable in your `$GOPATH/bin` directory, which should be in your system's PATH.

## Usage

To use the CLI, simply run the `github-activity` command followed by the GitHub username you want to query:

```bash
github-activity [username]
```

### Arguments

- `username`: The GitHub username for which to fetch activity. This is a required argument.

### Options

- `--help`: Displays the usage message and exits.

### Examples

1. **Get activity for a specific user:**

   ```bash
   github-activity dacostaaboagye
   ```

2. **Display help message:**

   ```bash
   github-activity --help
   ```

## Building from Source

If you prefer to build the executable from the source code, follow these steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Dacostasolo/github-activity.git
   cd github-activity
   ```

2. **Build the project:**

   ```bash
   make build
   ```

   This will create an executable named `github-activity` in the `bin/` directory.

3. **Project Source**
   ```bash
   roadmap.sh golang project https://roadmap.sh/projects/github-user-activity
   ```

## Running the Executable (after building)

After building the executable, you can run it directly:

```bash
./bin/github-activity [username]
```

## Cleaning the Build

To remove the built executable and other build artifacts:

```bash
make clean
```
