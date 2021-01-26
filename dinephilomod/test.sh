#!/bin/zsh

a=0
while [ "$a" -lt 1000 ]
do
    go run dinephilomod.go
    echo "*********************run number is $a"
    a=`expr $a + 1`
done

