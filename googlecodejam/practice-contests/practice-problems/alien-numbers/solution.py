#!/usr/bin/python3
'''
  Problem A. Alien Numbers (https://code.google.com/codejam/contest/32003/dashboard)
  RUN: solution.py < input_file > output_file
'''
import math

def convert_to_base10(number, language):
    '''Convert a number in a given language to base 10'''
    base10_number = 0
    digits = list(language)
    base = len(language)

    for i, val in enumerate(number[::-1]):
        base10_number += digits.index(val) * base ** i

    return base10_number


def convert_to_language(number, language):
    '''Convert a number un base 10 to a given language'''
    target_number = ""
    digits = list(language)
    base = len(language)

    while number > 0:
        target_number += digits[number % base]
        number = math.floor(number / base)

    target_number = target_number[::-1]
    return target_number


def main():
    cases = int(input())  # read a line with a single integer
    for i in range(1, cases + 1): # for each data line
        alien_number, source_language, target_language = [s for s in input().split(" ")]
        base10_number = convert_to_base10(alien_number, source_language)
        target_number = convert_to_language(base10_number, target_language)
        print("Case #{}: {}".format(i, target_number))


main()
