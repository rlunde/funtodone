
# set up GOPATH and GOBIN for this project
export GOPATH=$(pwd -P)
export GOBIN=$GOPATH/bin
echo "GOPATH set to current directory: " $GOPATH
echo "GOBIN set to \$GOPATH/bin: " $GOBIN
export PATH=$GOBIN:$PATH
export JAYCONFIG=/Users/ron/Dropbox/personal/projects/funtodone/funtodone/newsrc/blueprint/env.json
