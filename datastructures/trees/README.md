## Tree traversal
All in Depth First Search, this should be implemented using stacks or simply
use recursion which is a natural stack

- Preorder: print before the two recursive calls
- Inorder: prints in between the two recursive calls
- Postorder: prints after the two recursive calls


### Level order traversal
All in Breadth First Search, this should be implemented using iteration and queues.

**Note**: In embedded environment where only the stack (along with any read-only memory space) is available,
it's actually quite efficient to write software without using the heap memory at all if you know exactly
what your program is going to do.

---

## Real world usage

- **Binary Search Tree** - Used in many search applications where data is constantly entering/leaving,
  such as the map and set objects in many languages' libraries.

- **Binary Space Partition** - Used in almost every 3D video game to determine what objects need to be rendered.

- **Binary Tries** - Used in almost every high-bandwidth router for storing router-tables.

- **Hash Trees** - used in p2p programs and specialized image-signatures in which a hash needs to be verified, 
  but the whole file is not available.

- **Heaps** - Used in implementing efficient priority-queues, 
  which in turn are used for scheduling processes in many operating systems, 
  Quality-of-Service in routers, and A* (path-finding algorithm used in AI applications, 
  including robotics and video games). Also used in heap-sort.
  
- **Huffman Coding Tree** (Chip Uni) - used in compression algorithms, 
  such as those used by the .jpeg and .mp3 file-formats.
  
- **GGM Trees** - Used in cryptographic applications to generate a tree of pseudo-random numbers.
  
- **Syntax Tree** AST - Constructed by compilers and (implicitly) calculators to parse expressions.
  
- **Treap** - Randomized data structure used in wireless networking and memory allocation.
  
- **T-tree and B-tree** - Though most databases use some form of B-tree to store data on the drive, 
  databases which keep all (most) their data in memory often use T-trees to do so.
  Postgres uses B-tree for the index

### Why databases do not use Hash table instead of B-Tree?
In hash table (or associative array) can only access elements by their primary key in a hashtable.
This is faster than with a tree algorithm (O(1) instead of log(n)), but **you cannot select ranges**. 
Tree algorithms support this in Log(n) whereas hash indexes can result in a full table scan O(n). 
Also, the **constant overhead of hash indexes is usually bigger**. Also tree algorithms are usually easier to maintain, 
grow with data, scale.

### Why using a B-tree over a Binary search tree?
- B-trees are self-balanced. 
- B-Tree can fit a variable numbers of keys in one node, this will work nicely with CPU cache where 
we can utilize a block of memory more efficiently as we read data from a disk and put it into the CPU cache.
  So we can simply pick a B-tree order that will exactly fit in a block of memory, to **minimize disk to memory transfer**

## Common tree problems
- Find the height of a binary tree
- Find kth maximum value in a binary search tree
- Find nodes at “k” distance from the root
- Find ancestors of a given node in a binary tree

