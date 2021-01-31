# Binary Tree
A tree with at most 2 children and nodes do not have cyclic relations.
Trees are acyclic connected graph.

## Binary Search Tree
- The left subtree has smaller elements
- The right subtree has larger elements
- The left and right subtree each must also be a binary search tree
- Node without children are called `Leaf Node`
- The first node is called `Root Node`

### Operations
Worst case is when the tree is not balanced, and it ends up being almost like a linked list

- Insert `O(log(n))` Average, `O(n)` Worst. Item must be comparable with the other elements inside the tree
- Delete `O(log(n))`: find and remove
- Remove `O(log(n))`: find and replace
- Search `O(log(n))`


## Tree traversal
All in Depth First Search, this should be implemented using stacks

- Preorder: print before the two recursive calls
- Inorder: prints in between the two recursive calls
- Postorder: prints after the two recursive calls


### Level order traversal
All in Breadth First Search, this should be implemented using iteration and queues.

**Note**: In embedded environment where only the stack (along with any read-only memory space) is available, 
it's actually quite efficient to write software without using the heap memory at all if you know exactly 
what your program is going to do.


