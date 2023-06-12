use std::cell::{RefCell, RefMut};
use std::cmp::{Eq, PartialEq};
use std::collections::HashMap;
use std::hash::Hash;
use std::rc::Rc;

#[derive(Debug)]
struct LRU<K, V: Clone, const N: usize> {
    capacity: usize,
    length: usize,
    nodes: [Node<V>; N],
    lookup: HashMap<K, usize>,
    reverse_lookup: HashMap<usize, K>,
}

#[derive(Debug)]
struct Node<V: Clone> {
    value: V,
}

// impl<K: Clone, V: Clone, const N: usize> LRU<K, V, N>
// where
//     K: Eq + Hash,
// {
//     fn detach(&mut self, node: &mut RefMut<InternalNode<V>>) {
//         let prev_node = node.prev.take();
//         let next_node = node.next.take();
//         if let Some(prev) = prev_node {
//             prev.borrow_mut().next = next_node.clone();
//         }
//         if let Some(next) = next_node.clone() {
//             next.borrow_mut().prev = next_node;
//         }

//         node.next = None;
//         node.prev = None;

//         // update head/tail if necessary
//     }

//     fn prepend(&self, node: &mut RefMut<InternalNode<V>>) {
//         // update head (and tail if necessary)
//         // update links to prev head
//     }

//     fn trim_cache(&self) {
//         // remove least recently used elements if length greater than capacity
//     }
// }

#[cfg(test)]
mod test {
    // use super::*;
    //
    // fn lru_new() -> LRU<String, u32> {
    //     let mut lru: LRU<String, u32> = LRU::new(10);
    //     for i in 0..11 {
    //         lru.update(i.to_string(), i);
    //     }
    //     return lru;
    // }

    // #[test]
    // fn test_example() {
    //     let mut lru = lru_new();
    //     let val = lru.get(&"0".to_string());
    //     match val {
    //         Some(_) => assert!(false),
    //         None => assert!(true),
    //     }
    // }
}
