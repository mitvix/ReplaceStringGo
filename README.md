# Replace String in Go
This is a port in Go of an old replacestring.sh script, originally written in shell. 
The Go version has the same functionality but without dependencies on GNU tools like sed, xargs, and findutils. 
This provides the flexibility to run on other platforms, such as Windows or macOS. 

## What I can do
With this code, you can recursively replace content in files by simply passing the text you're looking for and the replacement text

## Configure
```
git clone https://github.com/mitvix/replacestringgo.git
cd replacestringgo
go build -o replacestring main.go
sudo cp replacestring /usr/local/bin/bash
```
## Usage: 
```
Usage: replacestring <path> <old_string> <new_string>
```

## Dependences

- a computer
- golang 1.22 or higher
- no external packages is needed
- and coffee

This is a Open Source software, please share! 

that's all folks
