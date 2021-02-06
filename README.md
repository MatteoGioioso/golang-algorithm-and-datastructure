# Time complexity

### Calculation rules

- Functions do more work for more input
- Constants can be dropped: `5 * n -> O(n)`, `80 * n => O(n)`
- Ignore lower order terms: `n^3 + n^2 + n + 5 -> O(n^3)`
- Ignore the base of logs


### Time complexity for various classes of problems

- **Linear algorithms (looping)**: If you are looping through the array and looking at all the elements one by one then it is O(n) because you spend constant time looking at a single element and there are n total elements.

- **Logarithmic algorithms (Trees)**: For binary search each time you look at an element you are essentially **discarding half of the array**. So after the first lookup you only care about n/2 elements in the array. After the second lookup, you only care about n/4 elements of the array. You do the pivot lookup operation {n/2, n/4, n/8 ... n/n} times and each time you do it you spend constant time since you are doing a simple comparison operation or equality check. So basically you are counting the number of elements in the set {n/2, n/4, n/8 ... n/n} which is log base 2 of n or O(logn)

- **nLogn algorithms (sorting)**: Another class of problems are problems like quicksort or mergesort. What is quicksort for example doing? You choose a pivot for an array/subarray then look at every single element of that array/subarray and arrange them in a way that smaller elements go to the left of the pivot and larger elements go to the right side of the pivot. What this means is that every time you choose a pivot, you do O(n) work since you are processing every single element of the array/subarray. First time you choose the pivot, you do O(n) work. Then you divide the array into two parts (which should be approximately equal to n/2). Then you do quicksort for the two halves that is O(n/2) + O(n/2) work. Then you divide those subarrays into halfs and choose pivots for the 4 smaller subarrays repeat the process. You continue until the array is fully sorted. What this means is every time you choose a pivot, you do 'n' units of work. And similar to binary search the number of times we choose the pivot = number of elements in the set {n/2, n/4, n/8 ... n/n} -> log(n)

- **Exponential complexities (compare all with all)**: Next let's take a look at exponential algorithms. This is basically what you studied in permutations and combinations during math class in school/college. Let's consider the recursive implementation of n queens problem. You have 8 rows to place the queens and the queen can be placed in one of 8 columns. So for the first column you have 8 possibilities, for the second column you have 8 possibilities and so on. Complexity is nn Let us consider the problem where we are asked to find all the permutations of the string with unique characters. Let's assume the string has 5 characters - there are 5 positions we are trying to fill. So in the first position you have 5 possibilities, for the second position you have 4 possibilities (since you placed one character in the first position), for the third position you have 3 possibilities (since 2 characters have been placed in positions 1, 2) and so on. The complexity = n*n-1*n-2...*1 and this is O(n!). Next suppose you are asked to find the number of possible values in a byte - one byte = 8 bits. You have 8 positions that can take 2 values (0 or 1). For position 1 you have 2 values, for position 2 you have 2 values and so on. The value of this is 28 hence the complexity of such problems is O(2n ).


### Notation

- **Big O**: same or faster
- **Big Theta**: same rate
- **Big Omega**: same or slower


# Contiguous and non-contiguous data structures

### Contiguous
In contiguous structures, terms of data are kept together in memory (either RAM or in a file).
An **array&& is an example of a contiguous structure.

Contiguous structures can be broken drawn further into two kinds:

- those that **contain data items of all the same size** (array), in an array, each element is of the same type.
- structures: in a struct, elements may be of different data types and thus may have different sizes.

### Non-contiguous
Non-contiguous structures are implemented as a collection of data-items, called nodes,
where each node can point to one or more other nodes in the collection. The simplest kind of
non-contiguous structure is **linked list**.
Tree and Graphs are non-contiguous as well.

### Hybrid
If two basic types of structures are mixed then it is a hybrid form. Then one part
contiguous and another part non-contiguous.

# Abstract Data Type (ADT)
An abstract data type is defined by its behavior from the point of view of a user, of the data, 
specifically in terms of possible values, possible operations on data of this type, and the behavior of these operations.

The design of a data structure involves more than just its organization. You also need to
plan for the way the data will be accessed and processed.
**Non-contiguous structures can be implemented either contiguously or non-contiguously like wise, 
the structures that are normally treated as contiguously can also be implemented non-contiguously**.

**The abstract notion of a data structure is defined in terms of the operations we plan to perform on the data**.
An abstract data type in a theoretical construct that consists of data as well as the operations to be performed on
the data while hiding implementation.

---

## Resources:

- https://www.youtube.com/watch?v=RBSGKlAvoiM
- http://www.jbiet.edu.in/coursefiles/cse/HO/cse1/DS1.pdf
