#! /bin/bash

if [[ -z $1 ]]; then
  echo "no input"
  exit 1
fi

DAY_NUMBER=$1

if [ -d "./day_${DAY_NUMBER}" ]; then
  echo "folder already exists"
  exit 1
fi

function make_solution_folder() {
  folder_name="solution_$1"
  mkdir $folder_name
  cd $folder_name
  cat "../../solution-template.go" >> main.go
  go mod init main.go
  cd ..
}

mkdir "day_${DAY_NUMBER}"
cd "day_${DAY_NUMBER}"
curl -H "Cookie: session=53616c7465645f5f7ec2582e1ae2ebbb147f62e80cf197cb8e2830001003ce38634f80997a9dee82ff4f9675b2e9841aeff5998ff9268a7c5d365d1802a5880f" \
      -o "input.txt" \
      "https://adventofcode.com/2024/day/${DAY_NUMBER}/input"

make_solution_folder "0"
make_solution_folder "1"

