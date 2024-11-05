# !/bin/bash

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

if [ $? != 0 ]; then
  echo "golangci-lint installation failed!"
  exit -1
fi

~/go/bin/golangci-lint > /dev/null

if [ $? != 0 ]; then
  echo "golangci-lint installation failed!"
  exit -1
fi

for i in *-service
do

  echo "Testing $i..."
  cd $i

  go test

  if [ $? != 0 ]; then
    echo "Failed!"
    exit -1
  fi

  for j in $(find . -type f -name *.go)
  do
    echo "Linting $j..."
    ~/go/bin/golangci-lint run $j

    if [ $? != 0 ]; then
      echo "Failed!"
      exit -1
    fi
  done

  cd ..
done
