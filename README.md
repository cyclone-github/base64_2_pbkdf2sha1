[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=base64_2_pbkdf2sha1&theme=gruvbox)](https://github.com/cyclone-github/)

## Cyclone's base64 to PBKDF2-HMAC-SHA1 hash converter
Program to convert base64 encoded strings to PBKDF2-HMAC-SHA1 (hashcat -m 12000).

### Example Usage:
- ./base64_2_pbkdf2sha1.bin -i base64.txt

### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/base64_2_pbkdf2sha1.git`
  - `cd base64_2_pbkdf2sha1`
  - `go mod init base64_2_pbkdf2sha1`
  - `go mod tidy`
  - `go build -ldflags="-s -w" .`
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt
