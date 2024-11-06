# !/bin/bash

check() {
  if [ $1 != 0 ]; then
    echo "Pipeline failed: $2"
    exit -1
  fi
}

setup() {
  echo "Setting up..."
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
  check $? "failed to install golangci-lint"
}

checktests() {
  echo "Checking tests for $1..."
  cd tests

  if [ $? == 0 ]; then
    go test
    check $? "go test failed!"
    cd ..
  fi
}

setup

for i in *-service
do
  echo "Entering $i..."
  cd $i
  golangci-lint run
  check $? "golangci-lint $i failed!"
  checktests
  cd ..
done
