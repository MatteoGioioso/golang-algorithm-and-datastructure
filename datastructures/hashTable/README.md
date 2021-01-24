# Hash table

Hash table are also called associative arrays for a valid reason, 
they use a hash function on the key of the element, 
and the remainder of its result to get an index. 
So simply speaking Hash table are just an implementation of arrays.


### Collisions
There are two types of solutions for collisions:

1. **Chaining**: Linked lists are used to avoid collisions 
   between keys; the better the hash function the less 
   collision we are going to have.
   

2. **Open addressing**: Open addressing means that, once a
   value is mapped to a key that's already occupied, 
   you move along the keys of the hash table until you 
   find one that's empty
   
### Rehashing

The **load factor** is a measure of how full the hash table is 
allowed to get before its capacity is automatically increased
As a general rule, the default load factor (.75) offers a good 
tradeoff between time and space costs.

If `lambda`:
```
lambda = initial_capacity * load_factor
```

will be reached then the capacity (array size) 
of HashMap will be doubled!
This will cause rehashing size the array size has changed, 
this is where `O(n)` can come from


### Time complexity
In general, we can assume that Hash table have a time complexity
of `O(1)` for Insert, Update, Get and Delete an element.
Worst case scenario is `O(n)`
