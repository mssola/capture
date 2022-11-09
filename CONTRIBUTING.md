# Contributing

This package is written in Go. Thus, you should follow the standards for this
programming language. In short, if you want to test your changes, just:

```bash
$ go test -v ./...
$ golint ./...
$ go vet ./...
```

## Issue reporting

I'm using Github in order to host the code. Thus, in order to report issues you
can do it on its [issue tracker](https://github.com/mssola/capture/issues). A
couple of notes on reports:

- Check that the issue has not already been reported or fixed in the `main` branch.
- Try to be concise and precise in your description of the problem.
- Provide a step by step guide on how to reproduce this problem.
- Provide the version you are using (git commit SHA).

## Pull requests

- Write a [good commit message](https://chris.beams.io/posts/git-commit/).
- Run the tests.
- Update the [changelog](./CHANGELOG.md) (if relevant).
- Open a pull request with *only* one subject and a clear title and description.
  Refrain from submitting pull requests with tons of different unrelated
  commits.
