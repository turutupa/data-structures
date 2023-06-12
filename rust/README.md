# Data Structures in Rust

## Doubly Linked Deque

Upon my failed attempt in implementing a LRU (see down below) I followed [this tutorial](https://rust-unofficial.github.io/too-many-lists/fourth.html) to learn how to create a doubly linked list which is used in a LRU. The main problem, no surprise, were the many borrow errors. 

The solution. This doubly linked deque uses Rc (Reference Counting)
- allows multiple ownership of data
- when an Rc smart pointer is cloned, it doesn't create a new data object; instead, it increases the count of references to the data object

and RefCell (Reference Cell)
- provides interior mutability 
- allows you to mutate data even when there are immutable references to that data
- enforces the borrowing rules at runtime (dynamic checking), instead of compile time
- it keeps track of borrowing with a borrowing counter

##  Least Recently Used implementation in Rust

This project is based on ThePrimeagen's course `The Last Algorithm's Course You'll Need`, more specifically, the section on Maps and LRU Cache. The aim of this project is to implement LRU in Rust instead of TypeScript.

https://frontendmasters.com/courses/algorithms/implementing-an-lru-cache/

