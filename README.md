# 🌀 GoLang-Algorithms-And-Data-Structures

Welcome to the GoLang Algorithms and Data Structures repository! Here you'll find implementations of various algorithms and data structures in GoLang.

## Algorithms

### Linear Search
The Linear Search algorithm sequentially checks each element in a slice until it finds the target element or reaches the end of the slice.
This algorithm is simple and straightforward, making it suitable for small slices or unsorted data.
One advantage is its ease of implementation and understanding, but it is less efficient than other search algorithms for larger datasets due to its linear time complexity.
The main idea behind linear search is to iterate through each element of the slice, comparing it with the target element until a match is found or the end of the slice is reached.
It is commonly used in situations where the dataset is small or unsorted, or when a quick and easy-to-understand solution is needed.

### Iterative Binary Search
The Iterative Binary Search algorithm, designed to find the index of a target value in a sorted slice. It operates by repeatedly dividing the search interval in half until the target value is found or the search interval is empty.
This algorithm boasts a logarithmic time complexity, making it highly efficient for large datasets compared to linear search.
Its main advantage lies in its ability to quickly locate target values in sorted slices, making it ideal for use cases where data is sorted and memory consumption is a concern.
The key idea behind binary search is to continuously halve the search interval, narrowing down the range where the target value may be found with each iteration.
This algorithm is widely used in scenarios where quick and efficient searching is required, such as in computer science algorithms and data structures.

### Recursive Binary Search
The Recursive Binary Search algorithm, searches for a target element in a sorted slice. It operates by continuously dividing the search interval in half until the target value is found or the search interval is empty.
This recursive approach offers the same logarithmic time complexity as the iterative binary search, providing efficient searching for large datasets.
One advantage of recursive binary search lies in its concise and elegant implementation, making it easy to understand and maintain.
The key idea behind recursive binary search is to divide the search interval in half recursively until the target value is found or deemed not present in the slice.
This algorithm is particularly useful in scenarios where a sorted dataset needs to be searched efficiently, such as in search algorithms and sorting algorithms.

### Recursive Binary Search Tree
A tree is a kind of data structure that is used to represent the data in hierarchical form. It can be defined as a collection of objects or entities called as nodes that are linked together to simulate a hierarchy. Tree is a non-linear data structure as the data in a tree is not stored linearly or sequentially.
A Binary Search Tree is a data structure used in computer science for organizing and storing data in a sorted manner. Each node in a Binary Search Tree has at most two children, a left child and a right child, with the left child containing values less than the parent node and the right child containing values greater than the parent node. This hierarchical structure allows for efficient searching, insertion, and deletion operations on the data stored in the tree.

### Exponential Search
Exponential Search is a searching algorithm designed to find an element in a sorted array. It’s named “exponential” because of the way it goes about searching for an element. This method has a time complexity of O(log n), which means it’s very efficient. The Exponential Search uses the Binary Search algorithm but in a different way. It starts with an element of the array and compares it with the target. If the target is greater, it doubles the number of elements taken from the array. Typically this number is called bound. If the bound is greater than the target, we search the array from bound /2 to bound. Now that we have a range, we can perform a binary search within that range.

### Selection Sort
The Selection Sort algorithm, sorts a slice in ascending order. This algorithm works by repeatedly finding the minimum element from the unsorted part of the slice and placing it at the beginning of the sorted portion.
It iterates through the slice, selecting the smallest element in each iteration and swapping it with the element at the current position.
One advantage of Selection Sort is its simplicity and ease of implementation, making it suitable for small datasets or situations where code simplicity is preferred over performance optimization.
The main idea behind Selection Sort is to iteratively select the smallest element from the unsorted portion of the slice and swap it with the element at the current position until the entire slice is sorted.
This algorithm's time complexity is O(n^2), making it less efficient than other sorting algorithms for large datasets.
However, Selection Sort's simplicity and ease of understanding make it a valuable teaching tool and a practical choice for sorting small slices.

### Bubble Sort
The Bubble Sort algorithm, sorts a slice in ascending order. This algorithm works by repeatedly iterating through the slice, comparing adjacent elements, and swapping them if they are in the wrong order.
It continues this process until the entire slice is sorted. One advantage of Bubble Sort is its simplicity and ease of implementation, making it suitable for educational purposes and small datasets.
However, Bubble Sort's time complexity is O(n^2), which can make it inefficient for large datasets.
The main idea behind Bubble Sort is to "bubble up" the largest elements to their correct positions by repeatedly comparing and swapping adjacent elements.
Despite its inefficiency for large datasets, Bubble Sort's simplicity makes it a useful introductory algorithm for understanding sorting concepts and algorithms.

### Quick Sort
The Quick Sort algorithm, sorts a slice in ascending order. Quick Sort is a divide-and-conquer algorithm that works by selecting a pivot element from the slice and partitioning the other elements into two sub-slices according to whether they are less than or greater than the pivot.
It then recursively sorts the sub-slices. One of the key advantages of Quick Sort is its efficient average-case time complexity of O(n log n), making it suitable for sorting large datasets.
However, Quick Sort's worst-case time complexity is O(n^2) when the pivot selection is poorly chosen, although this scenario is rare with good pivot selection strategies.
The main idea behind Quick Sort is to divide the slice into smaller sub-slices, sort the sub-slices recursively, and then combine them to produce the final sorted slice.

### Iterative Tree Sum
The Iterative Tree Sum algorithm, calculates the sum of all node values in a tree by performing an iterative depth-first traversal. It starts by initializing a sum variable to 0 and a stack containing the root node of the tree.
While there are nodes in the stack, it pops the top node, adds its value to the sum, and if the node has children, it pushes them onto the stack. This process continues until all nodes in the tree have been visited.
This algorithm offers an efficient way to compute the sum of tree nodes without using recursion, making it suitable for large trees and scenarios where recursion depth is a concern.

### Recursive Tree Sum
The Recursive Tree Sum algorithm, recursively calculates the sum of all node values in a tree by performing a depth-first traversal. It starts with an initial sum of 0 and iterates through each node in the tree.
For each node, it adds the value of the current node to the sum and, if the node has children, recursively calls the function on each child, adding their values to the sum.
This process continues until all nodes in the tree have been visited. While recursion offers a straightforward approach to tree traversal and sum computation, it may lead to stack overflow errors for very large trees due to excessive recursion depth.

### Cache
The cacheFunction higher-order function is designed to cache the results of another function, enhancing performance by avoiding redundant computations.
It takes a function fn as input and returns a cached version of that function. Internally, it maintains a cache object to store computed results for different inputs.
The factorial function calculates the factorial of a given number. When called with an argument n, it iteratively computes the factorial by multiplying n with each decreasing integer until reaching 1.
This implementation is straightforward but can be inefficient for repetitive calculations, which is where caching with cacheFunction becomes beneficial.

> [!TIP]
> When **cacheFunction** is called, it creates a new **lexical environment** (or scope) where the **cache object** is declared. The **inner function** returned by **cacheFunction** forms a **closure**, which means it **retains a reference** to the **lexical environment** where it was defined, including any **variables** in that environment (like **cache**).

Now, let's break down why the cache object is not removed from memory:

> [!IMPORTANT]
> 1. **Closure Retains References**: Even after **cacheFunction** finishes executing and its local variables (including **cache**) would typically be destroyed, the **inner function maintains a reference to the cache object**. This is because **the inner function is a closure and retains access to the lexical environment where cache is declared**.

> [!IMPORTANT]
> 2. **Garbage Collection**: GoLang engines use a mechanism called **"garbage collection"** to manage memory. Garbage collection automatically identifies and frees up memory that is no longer needed by the program. However, since the **inner function maintains a reference to the cache object**, the GoLang engine recognizes that **the object is still in use and therefore does not remove it from memory**.

> [!IMPORTANT]
> 3. **Repeated Access**: Each time you call a **certain instance, created from the cacheFunction**, you're essentially invoking **the same inner function that retains access to the same cache object**. This means that **the same cache object** is used across multiple invocations of a **certain instance, created from the cacheFunction**, leading to consistent behavior where cached results are retained and reused.

## Graph Algorithms

### Breadth First Search (BFS)
The Breadth-First Search (BFS) on a graph, starts from a specified vertex. It utilizes a queue data structure to explore vertices layer by layer, ensuring that vertices closer to the starting vertex are visited first.
This algorithm is particularly useful for finding the shortest path between two vertices in an unweighted graph, as it systematically explores all vertices at a given distance from the start vertex before moving on to vertices at the next distance.
The BFS traversal order returned by this function represents the sequence in which vertices are visited during the BFS traversal, providing insight into the structure and connectivity of the graph.

### Depth First Search (DFS)
The graph Depth-First Search (DFS) on a graph, starts from a specified vertex. It explores the graph's vertices in a depth-first manner, prioritizing traversal along a branch until it reaches a leaf node before backtracking to explore other adjacent vertices.
Utilizing recursion, it systematically visits each vertex, marking it as visited and adding it to the traversal order. Through the recursive dfs function, it traverses the graph, ensuring that each adjacent vertex is visited only once.
Finally, it returns the DFS traversal order, providing insight into the order in which vertices were visited during the traversal.

### Maximal Independent Set
The Maximal Independent Set algorithm aims to find an independent set within a graph. An independent set is a subset of vertices such that no two vertices in the set are adjacent. The algorithm operates iteratively, selecting a vertex of minimum degree in each iteration to ensure the independence of the set. By repeatedly selecting and removing vertices and their adjacent neighbors from the graph, it constructs a maximal independent set. This approach effectively identifies and includes vertices that can coexist without direct connections, ultimately returning a set representing the maximal independent set.

### Dijkstra's Algorithm
Dijkstra's algorithm, efficiently computes the shortest paths from a specified starting node to all other nodes within a weighted graph. Initially, it sets all node distances to infinity except for the starting node, which is set to zero.
Using a priority queue, it iteratively explores neighboring nodes, updating their distances if shorter paths are discovered. By repeatedly selecting and updating the shortest path candidates, the algorithm gradually converges, ultimately yielding an object containing the shortest distances from the start node to all other nodes in the graph.
This algorithmic approach is invaluable for finding optimal routes in various applications such as network routing and transportation logistics.

### Prim's Algorithm (Minimum Spanning Tree)
> [!IMPORTANT]
> *In graph theory, the problem of constructing a minimum spanning tree is one of the fundamental problems, which aims to find a subset of graph edges that connects all its vertices without forming cycles and has the minimum total cost. To solve this problem, there are several algorithms, one of which is the Prim method. This algorithm uses a greedy approach, starting with choosing a random vertex and sequentially adding edges with the least weight to the spanning tree, connecting already visited vertices with unvisited ones.*

Prim's algorithm, efficiently constructs the Minimum Spanning Tree (MST) of a graph by iteratively selecting edges with the minimum weight. Beginning from an arbitrary vertex, it progressively grows the MST by greedily choosing the edge that connects the existing MST to the nearest unvisited vertex.
This process continues until all vertices are visited, resulting in an MST that spans all vertices with minimal total edge weight. The calculatePrimMSTCost function complements this by computing the total cost of the MST.
Prim's algorithm finds extensive application in network design, where it helps minimize the cost of connecting various locations or components while ensuring connectivity.

### Kruskal's Algorithm (Minimum Spanning Tree)
Kruskal's algorithm, constructs the Minimum Spanning Tree (MST) of a weighted, connected, undirected graph by iteratively selecting edges with the minimum weight.
It begins by sorting all edges in increasing order of their weight and then iterates over these edges, adding them to the MST if they do not form a cycle.
This is achieved by maintaining a disjoint set data structure to track the connectivity of vertices. Kruskal's algorithm ensures that the resulting MST spans all vertices with the lowest possible total edge weight.
The calculateMSTCost function complements this by computing the total cost of the MST. Kruskal's algorithm finds applications in various fields, including network design, where it helps minimize the cost of connecting disparate components while maintaining connectivity.
