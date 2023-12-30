### Overview of Bloom Filters
A Bloom filter is an efficient probabilistic data structure, used to test whether an element is part of a set. It allows for false positives but guarantees no false negatives, making it suitable for applications where space efficiency is crucial, and some false positives are acceptable.

### How Bloom Filters Work
- **Initialization**: Begins with an array of `m` bits set to 0.
- **Insertions**: For each element, `k` independent hash functions map it to one of the `m` bits, which are then set to 1.
- **Query**: To check if an item is in the set, the same `k` hash functions are applied. If all `k` bits are set to 1, the item is presumed in the set, which might lead to a false positive.

### False Positive Probability
The probability of a false positive in a Bloom filter after `n` insertions is given by the expression \((1 - (1 - \frac{1}{m})^{nk})^k\).

- `m`: Number of bits in the Bloom filter.
- `n`: Number of elements inserted.
- `k`: Number of hash functions.

After `n` insertions, the probability that a specific bit is still 0 is \((1 - \frac{1}{m})^{nk}\). Consequently, the probability that a specific bit is 1 is \(1 - (1 - \frac{1}{m})^{nk}\). A false positive occurs if all `k` hash functions for a queried item point to bits that are already set to 1, and the probability of this happening is \((1 - (1 - \frac{1}{m})^{nk})^k\).

### Practical Implication
This formula is crucial in optimizing a Bloom filter's parameters (bit array size and the number of hash functions) based on the acceptable level of false positives and the anticipated number of elements to be inserted. It's a balance between space efficiency and accuracy.

### Resources
https://redis.com/blog/bloom-filter/