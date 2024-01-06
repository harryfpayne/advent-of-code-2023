# Day 25
## Part 1
I knew there would be some graph theory algorithm for this, I just didn't know which one. After learning that it was a
minimal cut problem it wasn't too tricky but there were a few dead ends I went down before this. I don't love the answer
to this, it's probabilistic so often takes a long time to run. The second time I ran it, it completed in 10s, so I decided
it was ok but that seems to have been a major outlier. I think a lot of the inefficiency is from reallocating the map
every time, and I push a lot of the complexity into counting the number of remaining edges, there's going to be a better
way to store the data which enables this to run much faster, but it's getting to the end of this and I want to stop having
to dedicate so much time to it, so I'll leave it here.

