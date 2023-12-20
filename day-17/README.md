# Day 17
## Part 1
Probably the most difficult part 1 so far. I vaguely remember for last year that for path finding I'd need a priority
queue so made a bad version of that initially which meant it took around a minute to get the answer. I got stuck on it 
finding incorrect paths for a while because I was only letting it take 2 moves in a straight line (the prompt treats 
turns as not part of the straight line).

## Part 2
Part 2 was pretty simple once part 1 was done. I spent some time optimising it, I kind of want all the solutions to be
under 5s run time, I assumed that parallelizing the queue processing would be the biggest win but I switched the queue to
a binary heap and that sped it up to under 2s so I was happy. I was really surprised how much it sped it up.

