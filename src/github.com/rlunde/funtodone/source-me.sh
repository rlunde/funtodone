
export CURPATH=$(pwd -P)
cd ../../../..
export GOPATH=$(pwd -P)
cd $CURPATH
echo "GOPATH IS " $GOPATH
echo "CWD IS " $CURPATH
export PATH=$GOPATH:$PATH

