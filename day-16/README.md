# Day 16
## Part 1
Pretty simple, the hardest part I found was getting which direction the beam goes to when it reaches a mirror. I also 
wasn't initially checking if a beam had already followed the path it was on because I assumed they would all go off the
sides, so had to spend a bit of time figuring out why it ran forever.

## Part 2
Also, not too bad. There was a weird memory thing which I didn't realise would happen, but I guess makes sense now. I'm 
calling `GetCoverage` on grid which contains a map of all the points visited. I'm reusing the grid each time but I thought 
it would be ok because I'm not using pointer to grid, so any mutations would only be scoped to the function but maps are
referencing a point in memory so even though grid is being copied, the map referenced is still the same, so data was 
leaking between runs.

