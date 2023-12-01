#!/bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

# the script must have one argument
if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

day=$(printf "%d" "$1") # %02d = performs decimal integer conversion 'd', formatted with zero padding '0', with width '2'

mkdir "AoC23/$day"
cd "AoC23/$day"
touch "input.txt"
touch "main.go"
touch "puzzle1.go"
touch "puzzle2.go"
go mod init "$day"
