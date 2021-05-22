## Real world usage

A real-world example would be a **FIFO queue**. 
A simple array-based list is pretty bad for that because you need to add at one end and remove at the other end, 
and one of those operations will be O(n).


### Linked lists are preferable over arrays when:

- you need constant-time insertions/deletions from the list 
  (such as in real-time computing where time predictability is absolutely critical)

- you don't know how many items will be in the list. With arrays, 
  you may need to re-declare and copy memory if the array grows too big

- you don't need random access to any elements

- you want to be able to insert items in the middle of the list (such as a priority queue)

### Cons:

- Filled arrays take up less memory than linked lists.
  Each element in the array is just the data. 
  Each linked list node requires the **data as well as one (or more) pointers to the other elements** in the linked list.
