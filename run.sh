#!/usr/bin/env bash

declare -i count=0

# write to file only the opportunities
test(){
  ((count++))
  if [ $count -lt 35 ]
  then
        ./spring_block >> results.out &
        sleep 2400
        echo $count
        pkill -f spring_block
        test
  fi
}


test
echo over