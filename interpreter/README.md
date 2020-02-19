# Interpreter

Source code from the [Writing an interpreter in Go][book] book.

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

[book]: https://www.amazon.co.uk/Writing-Interpreter-Go-Thorsten-Ball/dp/3982016118/ref=sr_1_1