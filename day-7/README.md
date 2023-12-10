# Day 7
### Part 1
This wasn't too hard, I might have written more code than necessary. I wanted to do it by implementing rust traits 
which add a lot of extra code. It could be optimized because I'm recomputing the type of the hand every comparison,
but I could instead calculate that once and then compare.

### Part 2
Struggled a bit more with this one, mainly because I missed the part where J becomes the lowest value. It also took some
time to figure out hand calculation, but I realised it's always optimal to put all the wildcards into the already most
frequent card. 