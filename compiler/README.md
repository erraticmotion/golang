# Compiler

Source code from the [Writing a compiler in Go][book] book.

## GO path

Create folder for go source code e.g. lexer

```console
mkdir lexer
cd lexer
echo export GOPATH=$(pwd) > .envrc
direnv allow 
mkdir src
cd src
-- add code project source code files
-- then Go will create /bin and /package directories when required.
```

[book]: https://www.amazon.co.uk/Writing-Compiler-Go-Thorsten-Ball/dp/398201610X/ref=sxts_sxwds-bia-wc1_0