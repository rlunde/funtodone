
# set up GOPATH and GOBIN for this project
export GOPATH=$(pwd -P)
export GOROOT=/Users/ron/Projects/go/go
export GOBIN=$GOROOT/bin
echo "GOPATH set to current directory: " $GOPATH
echo "GOROOT set to: " $GOROOT
echo "GOBIN set to \$GOROOT/bin: " $GOBIN
export PATH=$GOBIN:$PATH
