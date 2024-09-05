#### Design Hashmap

Design a HashMap data structure that supports the following operations:

- **Constructor()**: Initializes the hash map object with an empty map to
  store the key-value pairs.
- **Put(key, value)**: Inserts a key-value pair into the hash map. If the
  specified key is already present in the hash map, the associated value is
  updated. Otherwise, the key-value pair is added to the hash map.
- **Get(key)**: Returns the value associated with the specified key if the key
  exists in the hash map. Otherwise, it returns −1, indicating the absence of a
  mapping for the key.
- **Remove(key)**: Removes the entry for the specified key from the hash map,
  effectively removing both the key and its associated value. This elimination
  only takes place when the key exists in the hash map.

> **Note**: Your implementation should not utilize any built-in hash table
> libraries

**Constraints**:

- <code>0 ≤ key ≤ 10<sup>6</sup></code>
- <code>0 <=value <= 10<sup>6</sup></code>
- At most, <code>10<sup>6</sup></code> calls can be made to **Put**, **Get**,
  **Remove** functions.
