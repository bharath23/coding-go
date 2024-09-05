package educative0007

// Bucket represents a bucket for storing key-value pairs.
type Bucket struct {
	bucket []Pair // Slice to store key-value pairs
}

// Pair represents a key-value pair.
type Pair struct {
	key   int
	value int
}

// NewBucket creates a new instance of Bucket.
func NewBucket() *Bucket {
	return &Bucket{
		bucket: make([]Pair, 0), // Initialize the slice
	}
}

// Get return the value from the Bucket
func (b *Bucket) Get(hashKey int) int {
	for _, p := range b.bucket {
		if p.key == hashKey {
			return p.value
		}
	}
	return -1
}

// Update function put value in the Bucket
func (b *Bucket) Update(hashKey int, value int) {
	found := false
	for i, p := range b.bucket {
		if p.key == hashKey {
			found = true
			b.bucket[i].value = value
		}
	}

	if !found {
		b.bucket = append(b.bucket, Pair{key: hashKey, value: value})
	}
}

// Remove function delete value from the Bucket
func (b *Bucket) Remove(hashKey int) {
	idx := -1
	for i, p := range b.bucket {
		if p.key == hashKey {
			idx = i
			break
		}
	}

	if idx != -1 {
		b.bucket = append(b.bucket[:idx], b.bucket[idx+1:]...)
	}
}

type DesignHashMap struct {
	hashKey int
	hashMap []*Bucket
}

// Use the constructor below to initialize the hash map based on the keyspace
func Constructor() DesignHashMap {
	h := DesignHashMap{
		hashKey: 4001,
		hashMap: make([]*Bucket, 4001),
	}

	for i := range h.hashKey {
		h.hashMap[i] = NewBucket()
	}

	return h
}

func (this *DesignHashMap) Put(key int, value int) {
	hashKey := key % this.hashKey
	b := this.hashMap[hashKey]
	b.Update(key, value)
}

func (this *DesignHashMap) Get(key int) int {
	hashkey := key % this.hashKey
	b := this.hashMap[hashkey]
	return b.Get(key)
}

func (this *DesignHashMap) Remove(key int) {
	hashkey := key % this.hashKey
	b := this.hashMap[hashkey]
	b.Remove(key)
}
