# 1.1

- dev: using `staticcheck`. This change forced us to change a bit the string on `panic` handling (thus the version change).
- dev: not using the deprecated `io/ioutil` module anymore.

# 1.0

- Implemented the `All` public function that users can call in order to safely
  capture all output from a given block of code.
