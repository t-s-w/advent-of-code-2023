# Notes on Day 6

The problem of finding the number of ways to win is basically finding the times `x` for which `x(time - x) > distance`, equivalent to finding integer values of x for which `x^2 - time * x + distance < 0`. This is the count of all integers between the two roots (exclusive) of this equation.
