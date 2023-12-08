# Notes for day 1

Guess Regexp.ReplaceAllStrings() is a goroutine.

Originally I had an implementation where I'd first attempt to replace every "worded" number with the number, and then implement the solution from part 1. This didn't work because of a race condition; "eightwo" would sometimes become "8wo" and other times "eigh2". Turns out that Replace All Strings is async. Every time I ran the code, a different result would come out, and I was baffled to no end because the few cases I looked at - the 10-ish at the start and end of the inputs - didn't have this condition.

There had to be a way to make it consistent, though. The problem was well-defined. I suppose it would've been possible to do a greedy algorithm and figure out a way to find the "first" and "last" digits, but that's an optimisation for a later time.
