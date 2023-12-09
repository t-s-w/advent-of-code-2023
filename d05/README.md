# Notes for day 5

Okay, when I start a new folder/package/module for day 5, how am I supposed to even `go mod init` it? `go mod init main`? `go mod init d05`? If I do the latter, go complains if i try to `go run script.go`. If I do the former, I'll have one million packages with the name `main`. For all the nice things Go offers the module management is pretty annoying. Like why isn't there even any way to reference modules by relative paths???

The [official documentation](https://go.dev/doc/modules/layout) doesn't even help:

> A module can consist of multiple importable packages; each package has its own directory, and can be structured hierarchically. Here’s a sample project structure:

```
project-root-directory/
 go.mod
 modname.go
 modname_test.go
 auth/
   auth.go
   auth_test.go
   token/
     token.go
     token_test.go
 hash/
   hash.go
 internal/
   trace/
     trace.go
```

> Sub-packages can be imported by users as follows:

```
import "github.com/someuser/modname/auth"
import "github.com/someuser/modname/auth/token"
import "github.com/someuser/modname/hash"
```

Except that all that was a scam. I tried `go get github.com/t-s-w/advent-of-code-2023/utils` and it complained:

```
go: github.com/t-s-w/advent-of-code-2023/utils@upgrade (v0.0.0-20231209070949-8cb09bc822e6) requires github.com/t-s-w/advent-of-code-2023/utils@v0.0.0-20231209070949-8cb09bc822e6: parsing go.mod:
        module declares its path as: utils
                but was required as: github.com/t-s-w/advent-of-code-2023/utils
```

???? Is my only counterplay copying the `replace` statement every `go.mod` or what

## Part 2 brute force method

...I could hear my CPU humming as I turned it on. I'm sure there's a better way that manages interval logic and making a method of composing maps together. But that's a problem for another time, and my CPU working hard for 2 minutes certainly did it faster than I would've been able to code up that logic (and I have to go out for dinner in 10).
