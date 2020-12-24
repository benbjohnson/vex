vex
===

Vex is a variable-length, lexicographically-sortable hex format for uint64 values.
It can be used instead of `fmt.Sprintf("%016x", v)` for shorter strings.


## Usage

Values can be formatted to strings using the `Format()` function and formatted
strings can be parsed using the `Parse()` function. Easy peasy.

```go
Format(v uint64) string
Parse(s string) (uint64, error)
```


## How it works

Vex simply formats a value as an unpadded hex-encoded number and then prepends
the length of the resulting string (minus 1) as a hex value.

Here are some examples of encodings:

```
DECIMAL   HEX   VEX    ZERO-PADDED
0         0     00     0000000000000000
1         1     01     0000000000000001
2         2     02     0000000000000002
10        a     0a     000000000000000a
15        f     0f     000000000000000f
16        10    110    0000000000000010
288       120   2120   0000000000000120
```
