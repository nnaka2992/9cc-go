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
assert 0 'return 0;'
assert 42 'return 42;'
# case with addtion/subtraction only expression
assert 21 "return 5+20-4;"
assert 41 "return  12 + 34- 5;"
# case with alithmetic operation with prantheses
assert 47 'return 5+6*7;'
assert 15 'return 5*(9-6);'
assert 4 'return (3+5)/2;'
# case with unary opeartor
assert 10 'return -10+20;'
assert 10 'return 20+-10;'
# case with equality operator
assert 1 'return 1==1;'
assert 0 'return 1==0;'
# case with combination of equality and alithmetic operation with prantheses
assert 2 'return 1+(1==1);'
assert 0 'return 1==1+1;'
# case with variable
assert 1 'return a=b=1;'
assert 0 'return a=2+2==1*2;'
assert 0 'return abc=2+2==1*2;
abc=abc-1;
abc=abc*0;
return abc;'

echo OK
