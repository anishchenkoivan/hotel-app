# !/bin/bash


golangci-lint > /dev/null

if [ $? != 0 ]; then
  echo "golangci-lint is required!"
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
    golangci-lint run $j

    if [ $? != 0 ]; then
      echo "Failed!"
      exit -1
    fi
  done

  cd ..
done
