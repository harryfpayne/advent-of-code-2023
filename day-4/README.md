# Day 4
### Part 1
This was nice and simple, nothing too hard here

### Part 2
I really struggled with this one, I did the naive solution using for loops
to get the answer, but I can tell there's some underlying pattern that I'm
not spotting.

Part 1 making you shift bits was an obvious hint that powers of 2 are involved,
and looking in my `repetitions` vector there's clearly some pattern, for the
test case the repetitions go `1, 2, 4, 8, 14, 1`, it counts up in powers of 2
until it gets to the 5th element which is the first 'not reached' by a card
(Card 2 doesn't increment it). At the moment I'm having to do this nested for
loop which is incrementing each subsequent card, it feels like there's some
trick that can do it in a single pass, probably with bit shifting.