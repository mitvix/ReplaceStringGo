# Replace String in Go
This is a port in Golang of old replacestring.sh written originally in shell.

In this Go version we have the same funcionality but without dependencies of GNU Tools like sed, xargs and findutils and with the flexibility to run in another platforms, like Win or Mac.

With this code it's possible replace content in files recursively, you just need pass the text that you are looking for and the text used to exchange.

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

This is a Open Source software, please share! 

that's all folks
