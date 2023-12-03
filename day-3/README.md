# Day 3
### Part 1
Part 1 was pretty simple, the main insight was collecting 
all the numbers first, then for each one checking if it's adjacent
to any symbols.

This could be improved to run in a single pass, but the code got 
a bit complicated and I assumed I would need the all the numbers
for part 2.

Also assuming part 2 would be more tedious, I added some helpers
like `Coord::adjacent_coords` because advent of code loves that kind of thing.

### Part 2
Again not too hard, I ran into some problems with the rust borrow checker at the end
which took a few attempts to figure out (I'd forgotten to put `&` where they were needed).

Again this could be more performant, it actually takes a few seconds to run the full
puzzle. I could be done in 1 pass, look for '*' symbols, check adjacent cells for numbers,
if it's an adjacent cell then resolve the full number (this is the tricky bit). I settled 
for the code as it was because it reused a lot of what I'd already written for part 1.