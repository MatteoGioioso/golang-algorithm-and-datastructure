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

