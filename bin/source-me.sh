
# set up GOPATH and GOBIN for this project
export GOPATH=$(pwd -P)
export GOBIN=$GOPATH/bin
echo "GOPATH IS " $GOPATH
echo "GOBIN IS " $GOBIN
export PATH=$GOBIN:$PATH

