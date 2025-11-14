package mapcode

import "fmt"
 
/*
Explanation
apple → Hash value: 0
banana → Hash value: 9
grape → Hash value: 5
melon → Hash value: 5 (Collision with "grape")
orange → Hash value: 3
peach → Hash value: 8
Key Observations
✅ The hash function maps keys to indices efficiently.
✅ Collision occurs when "grape" and "melon" both map to index 5.
In a real-world scenario, collision-handling strategies like chaining or open addressing are necessary to manage such cases. 

collison occurs when we try to automate hashvalue generation for keys for handling this we now here using linkedlist based hash value generation which completelu avoids collision.*/


// Node represents a key-value pair for chaining
type Node struct {
    key   string
    value int
    next  *Node
}

// HashTable structure
type HashTable struct {
    buckets []*Node
    size    int
}

// Hash function (simple character sum mod table size)
func (ht *HashTable) hash(key string) int {
    hashValue := 0
    for _, ch := range key {
        hashValue += int(ch)
    }
    return hashValue % ht.size
}

// Insert key-value pair
func (ht *HashTable) Insert(key string, value int) {
    index := ht.hash(key)
    newNode := &Node{key: key, value: value}

    // Insert at the beginning if the bucket is empty
    if ht.buckets[index] == nil {
        ht.buckets[index] = newNode
        return
    }

    // Collision handling using chaining (linked list)
    current := ht.buckets[index]
    for current != nil {
        if current.key == key {
            current.value = value // Update if key exists
            return
        }
        if current.next == nil {
            current.next = newNode // Append at the end of the chain
            return
        }
        current = current.next
    }
}

// Get value by key
func (ht *HashTable) Get(key string) (int, bool) {
    index := ht.hash(key)
    current := ht.buckets[index]

    for current != nil {
        if current.key == key {
            return current.value, true
        }
        current = current.next
    }
    return 0, false // Key not found
}

// Display hash table contents
func (ht *HashTable) Display() {
    for i, bucket := range ht.buckets {
        fmt.Printf("Bucket %d: ", i)
        current := bucket
        for current != nil {
            fmt.Printf("[%s: %d] -> ", current.key, current.value)
            current = current.next
        }
        fmt.Println("nil")
    }
}

func hasing_linkedlist() {
    ht := &HashTable{
        buckets: make([]*Node, 5), // Small size to force collisions
        size:    5,
    }

    ht.Insert("apple", 100)
    ht.Insert("banana", 200)
    ht.Insert("grape", 300)
    ht.Insert("orange", 400)
    ht.Insert("melon", 500)
    ht.Insert("peach", 600)

    ht.Display()

    // Retrieve values
    if val, found := ht.Get("apple"); found {
        fmt.Println("Value for 'apple':", val)
    } else {
        fmt.Println("'apple' not found")
    }

    if val, found := ht.Get("melon"); found {
        fmt.Println("Value for 'melon':", val)
    } else {
        fmt.Println("'melon' not found")
    }
}
