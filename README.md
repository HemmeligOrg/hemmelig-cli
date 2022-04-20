# Hemmelig CLI

With this CLI you can use the API from [https://hemmelig.app](https://hemmelig.app) to create secret URLs on the fly.

## Features

- Pipe data to Hemmelig
- Set a password for the secret URL
- Adjust the time to live (TTL) for the secret
- Override the URL if you host Hemmelig yourself

## Usage

Install the binary (windows, osx, linux) from <https://github.com/HemmeligOrg/hemmelig-cli/releases>, or go directly to the build the binary manually step.

```bash
# Example
cat your_secret_file.txt | hemmelig --password=cantguessthislol
# The secret URL: https://hemmelig.app/secret/0-ii79E5tViCv6OBPEmzC/talented_RU4NfXNvxTLJAf1R_QFtp
```

```bash
NAME:
   [he`m:(ə)li] - Create a secret URL directly from your CLI.

USAGE:
   cat your_secret_file.txt | hemmelig --password=cantguessthislol
   Or just pass it as the first argument: hemmelig "This is my secret" --password=secret

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --password value, -p value  Set a password to protect the secret
   --ttl value, -t value       Secret expiration time in seconds. 0 - 605800 seconds. (default: "14400")
   --url value, -u value       Override the Hemmelig app URL if you host it yourself (default: "https://hemmelig.app/")
   --help, -h                  show help (default: false)

COPYRIGHT:
   (c) 2022 Hemmelig.app
```

## Get it up and running [DEV]

```bash
# Install dependencies
go install

# By using the go binary directly
go run main.go
```

## Build the binary manually

To use the bleeding edge codebase, this is the way.

```bash
# Build binary
go build -o hemmelig

chmod +x hemmelig

mv hemmelig /usr/local/bin

hemmelig "my secret text"
```

## Disclaimer

Use this tool at your own risk. The owner of this repository is not responsible for its usage.

## LICENSE

See [LICENSE](./LICENSE)
