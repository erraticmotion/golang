# Go In Practice

Source code from the [Go In Practice][book] book.

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

[book]: https://www.amazon.co.uk/Go-Practice-Matt-Butcher/dp/1633430073/ref=sr_1_1