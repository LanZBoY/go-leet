# Analysis: Reverse Linked List Approaches

This document compares three common ways to reverse a singly linked list: Iterative, Recursive, and Stack-based.

## 1. Iterative Approach (`sln_v1`)

The iterative approach uses three pointers (`prev`, `curr`, `next`) to reverse the directions of the links in-place.

- **Pros**: 
  - Most efficient in terms of space.
  - No recursive stack overhead.
- **Cons**: 
  - Slightly more complex to reason about if you're not comfortable with pointer manipulations.
- **Complexity**:
  - Time: $O(n)$ - one pass through the list.
  - Space: $O(1)$ - only constant extra space used for pointers.

## 2. Recursive Approach (`sln_v2`)

The recursive approach reverses the "rest" of the list first, then attaches the current head to the end of that reversed list.

- **Pros**: 
  - Elegant and concise code.
  - Very readable once you understand the base case and recurrence.
- **Cons**: 
  - Uses $O(n)$ space for the call stack.
  - Can lead to `Stack Overflow` on very long lists.
- **Complexity**:
  - Time: $O(n)$ - each node is visited once.
  - Space: $O(n)$ - depth of the call stack.

## 3. Stack-based Approach (`sln_v3`)

The stack-based approach pushes all node pointers into a stack, then pops them one by one to rebuild the list in reverse order.

- **Pros**: 
  - Conceptually very simple. Reverses the order naturally by using a LIFO structure.
- **Cons**: 
  - Uses the most explicit memory (storing $n$ pointers in a slice/stack).
  - Two passes: one to push, one to pop (though still $O(n)$).
- **Complexity**:
  - Time: $O(n)$ - two passes.
  - Space: $O(n)$ - storing all nodes in a stack.

## Summary Table

| Approach | Time | Space | Implementation |
| :--- | :--- | :--- | :--- |
| **Iterative** | $O(n)$ | $O(1)$ | Best for production code due to space efficiency. |
| **Recursive** | $O(n)$ | $O(n)$ | Great for interviews/conceptual understanding. |
| **Stack-based** | $O(n)$ | $O(n)$ | Simple to think about but memory-heavy. |
