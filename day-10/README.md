# Day 10
## Part 1
This went ok, I went for a strategy that follows a pipe and keeps track of the direction I came from to know where to
go next which got a bit confusing knowing how each pipe transformed the direction I'm going, that code could definitely
be simplified because I'm applying a movement direction e.g. go north, then inverting it for the next move because I've
just come from the south.

Overall I think it's a pretty efficient solve.

## Part 2
This also was ok, the squeezing through gaps was the tricky bit. I solved it by having an expanded version of the map
which made each pipe segment a 3x3 grid so that gaps between pipes were easily identifiable. Then for every space I 
check if I'm able to reach the border; I initially implemented this check using recursion but recursion is really slow in
rust, so I moved to a while loop that goes through an increasing job list which solved it pretty fast.

The main optimisation I can think to make is that I'm search from each point without the knowledge of what I've done before
which probably leads to rechecking points I already know the outcome of, so having some kind of cache would speed it up.

Out of curiosity I implemented it with a cache and it solved it in 45ms compared to 10s so a big improvement.
