#!/bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

# the script must have one argument
if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

day=$(printf "%02d" "$1") # %02d = performs decimal integer conversion 'd', formatted with zero padding '0', with width '2'

mkdir "AoC23/$day"
cd "AoC23/$day"
touch "input.txt"
cp "../../tmp/main.go" "main.go"
cp "../../tmp/puzzle1.go" "puzzle1.go"
cp "../../tmp/puzzle2.go" "puzzle2.go"
touch "README.md"
echo "# Day $day" >> "README.md"
go mod init "$day"
