[![Readme Card](https://github-readme-stats-fast.vercel.app/api/pin/?username=cyclone-github&repo=base64_2_pbkdf2sha1&theme=gruvbox)](https://github.com/cyclone-github/)

## Cyclone's base64 to PBKDF2-HMAC-SHA1 hash converter
Program to convert base64 encoded strings to PBKDF2-HMAC-SHA1 (hashcat -m 12000).

### Example Usage:
- ./base64_2_pbkdf2sha1.bin -i base64.txt

### Install latest release:
```
go install github.com/cyclone-github/base64_2_pbkdf2sha1@latest
```
### Install from latest source code (bleeding edge):
```
go install github.com/cyclone-github/base64_2_pbkdf2sha1@main
```

### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/base64_2_pbkdf2sha1.git`  # clone repo
  - `cd base64_2_pbkdf2sha1`                                               # enter project directory
  - `go mod init base64_2_pbkdf2sha1`                                      # initialize Go module (skips if go.mod exists)
  - `go mod tidy`                                              # download dependencies
  - `go build -ldflags="-s -w" .`                              # compile binary in current directory
  - `go install -ldflags="-s -w" .`                            # compile binary and install to $GOPATH
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt
