# !/bin/sh

for i in *-service
do

  echo "Testing $i..."
  cd $i

  go test

  cd ..
done
