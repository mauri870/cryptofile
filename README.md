# Cryptofile

Encrypt or decrypt files using AES-256 or AES-128

> Caution when using with the -delete flag, if you miss the key used for encrypt your file probably you never decrypt it

## Installation
```
 go get -v github.com/mauri870/cryptofile
 go install github.com/mauri870/cryptofile
```

# Usage
Use `cryptofile -h` for a complete list of flags

```
echo Hello World! > filetoencrypt.txt
cryptofile -key Your32or16bytesAlphanumericKey -in filetoencrypt.txt
cryptofile -decrypt -key Your32or16bytesAlphanumericKey -in filetoencrypt.txt.encrypted
```
You can also use `-delete` to remove the input file
