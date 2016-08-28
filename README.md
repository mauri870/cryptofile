# Cryptolock

Encrypt or decrypt files using AES-256 or AES-128

> Caution when using with the -delete flag, if you miss the key used for encrypt your file probably you never decrypt it

## Installation
```
 go get -v github.com/mauri870/cryptolock
 go install github.com/mauri870/cryptolock
```

# Usage
Use `cryptolock -h` for a complete list of flags

```
echo Hello World! > filetoencrypt.txt
cryptolock -key Your32or16bytesAlphanumericKey -in filetoencrypt.txt
cryptolock -decrypt -key Your32or16bytesAlphanumericKey -in filetoencrypt.txt.enc
```
You can also use `-delete` to remove the input file
