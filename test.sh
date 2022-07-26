#!/bin/bash
assert() {
    expected="$1"
    input="$2"

    go run *.go "$input" > tmp.s
    cc -o tmp tmp.s
    ./tmp
    actual="$?"

    if [ "$actual" = "$expected" ]; then
        echo "$input => $actual"
    else
        echo "$input => $expected expected, but got $actual"
        exit 1
    fi
}

# case with integer input
assert 0 '0;'
assert 42 '42;'
# case with addtion/subtraction only expression
assert 21 "5+20-4;"
assert 41 " 12 + 34- 5;"
# case with alithmetic operation with prantheses
assert 47 '5+6*7;'
assert 15 '5*(9-6);'
assert 4 '(3+5)/2;'
# case with unary opeartor
assert 10 '-10+20;'
assert 10 '20+-10;'
# case with equality operator
assert 1 '1==1;'
assert 0 '1==0;'
# case with combination of equality and alithmetic operation with prantheses
assert 2 '1+(1==1);'
assert 0 '1==1+1;'
# case with variable
assert 1 'a=b=1;'
assert 0 'a=2+2==1*2;'
assert 0 'abc=2+2==1*2;'

echo OK
