#!/bin/bash
# Create the directory for a new day's code.
#

usage () {
    echo "Usage: $0 <day>"
    echo "Create the directory for the given day."
    exit 0
}

error () {
    echo "Error: $*"
    exit 1
}

test -n "$1" || usage
DAY=$1

dir="day${DAY}"
prog1="day${DAY}.go"
prog2="day${DAY}_test.go"

mkdir -p $dir                                          || error "Unable to create $dir"
test -f "$dir/$prog1" || cp dayN.go "$dir/$prog1"      || error "Unable to install $prog1"
test -f "$dir/$prog2" || cp dayN_test.go "$dir/$prog2" || error "Unable to install $prog2"

echo "Created $dir"

