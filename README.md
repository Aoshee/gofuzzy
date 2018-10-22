# GoFuzzy

GoFuzzy is a web directory/file finder and a HTTP request fuzzer. GoFuzzy is inspired by `wfuzz` which is one of my favorite tools.
An occurence of the `FUZZ` keyword (anywhere in the request) will be replaced by an entry from the wordlist.

## Sample run

```bash
$ gofuzzy -u example.com -w wordlists/wl.txt

   _ __________ + _________     *
_ _ __ /  ________/ ____________________ ___   *
    __/  /  / _  /  __/  / /__ /___  /  /  / -+
   +  \____/\___/__/  \___/_____/____\_   /
           *          - -+          /____/        *

---------------------------------------------------------------------------------
Chars(-hh)    Words(-hw)   Lines(-hl)   Header(-hr)  Code(-hc)    Result
---------------------------------------------------------------------------------
185           22           7            140          301          Admin
185           22           7            140          301          Login
185           22           7            140          301          login
0             0            0            198          200          passwords
185           22           7            119          301          test
```

## Build

Make sure Go is [installed](https://golang.org/doc/install) and the `$GOPATH` is set correctly:

```bash
go get github.com/shellrausch/gofuzzy
cd $GOPATH/src/github.com/shellrausch/gofuzzy
go build
```

## Install

Make sure you have followed the _Build_ step above:

```bash
cd $GOPATH/src/github.com/shellrausch/gofuzzy
go install
```

After the installation the `gofuzzy` binary will be in `$PATH`. That means you can call `gofuzzy` from everywhere.

### Kali 2018.3

Install Go:

```bash
apt-get update
apt-get install golang-1.10 -y
```

Configure pathes:

```bash
mkdir $HOME/go
export GOROOT=/usr/lib/go-1.10
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
```

Install GoFuzzy:

```bash
go get github.com/shellrausch/gofuzzy
cd $GOPATH/src/github.com/shellrausch/gofuzzy
go install
gofuzzy -h
```

## Run/Usage

Find hidden files or directories:

```bash
gofuzzy -u example.com -w wl.txt
gofuzzy -u example.com/subdir/FUZZ/config.bak -w wl.txt
```

Brute force a header field:

```bash
gofuzzy -u example.com -w wl.txt -H "User-Agent: FUZZ"
```

Brute force a file extension:

```bash
gofuzzy -u example.com/file.FUZZ -w ext.txt
```

Brute force a password send via a form with POST:

```bash
gofuzzy -u example.com/login.php -w wl.txt -m POST \
    -d "user=admin&passwd=FUZZ&submit=s" \
    -H "Content-Type: application/x-www-form-urlencoded"
```

Brute force HTTP methods:

```bash
gofuzzy -u example.com -w wl.txt -m FUZZ
```

## Docker

Build the image:

```bash
cd $GOPATH/src/github.com/shellrausch/gofuzzy
docker build -t gofuzzy .
```

Run GoFuzzy in a container:

```bash
docker run -v $(pwd)/wordlists:/wordlists gofuzzy -u example.com -w /wordlists/wl.txt
```

## Wordlists

See [SecLists](https://github.com/danielmiessler/SecLists/tree/master/Discovery/Web-Content) for recommended wordlists.