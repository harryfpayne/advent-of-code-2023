# Day 1
### Part 1
Part 1 was easy enough, just loop through all characters in a line, keep track of the numeric ones, and then use first and last on the vector.

This could be more space efficient, instead of using a vector to store all the numbers 
instead use two vars `first`, `last`, overriding last with each letter.

I don't think it can be more time efficient, it is necessarily O(n) because you need to look
through every character in each line, I doubt the first and last methods on a vector have any
meaningful complexity, I'm assuming it's an O(1) lookup but might be wrong.

### Part 2
Part 2 was harder, I think most of the complexity was fighting against rust strings. Rust strings are built to handle 
UTF-8 so getting individual characters and indexed substrings doesn't really make sense.  

I make a dictionary of all the words, for efficiency I then make a vector of all the  word lengths I need to look for i.e. `[3,4,5]`.
Then looking through each character, this time using an index (`idx`), get a character if it's numeric do the same as part 1,
otherwise go through all the word lengths, getting a substring of the line at `idx` of the current word length and see if
it's in the map, adding the corresponding value if it is.

There's probably some time complexity saving I could do here, I'm brute forcing looking for words.