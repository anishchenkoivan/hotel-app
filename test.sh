# !/bin/sh

for i in *-service
do

  echo "Testing $i..."
  cd $i

  go test

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
