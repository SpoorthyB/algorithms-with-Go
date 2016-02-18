package main

// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3

type CacheNode struct {
    key int
    value int
    prev *CacheNode 
    next *CacheNode 
}
func NewCacheNode(key int, value int) *CacheNode{
	obj := new(CacheNode)
	obj.key = key
	obj.value = value
	obj.prev = nil
	obj.next = nil
	return obj
}

type LRUCache struct {
   capacity int
   head *CacheNode 
   tail *CacheNode
   nodeNum int 
   m map[int]*CacheNode
   //m = make(map[int]int)
}
 func NewLRUCache(cap int) *LRUCache {
 	cach := new(LRUCache)
    cach.capacity = cap;
    cach.m = make(map[int]*CacheNode)
    cach.head = NewCacheNode(-1,-1)
    cach.tail = NewCacheNode(-1, -1)
    cach.head.next = cach.tail
    cach.tail.prev = cach.head
    cach.nodeNum = 0
    return cach
}
  func (lru *LRUCache) removeNode(node *CacheNode) {
 	
 	//fmt.Println(*(*node.prev))
    node.prev.next = node.next
    node.next.prev = node.prev
    lru.nodeNum--
}

  func (lru *LRUCache)  appendNode(node *CacheNode ) {
    node.next = lru.tail
    lru.tail.prev.next = node
    node.prev = lru.tail.prev
    lru.tail.prev = node
    lru.nodeNum++
}

  func (lru *LRUCache) updateAccessTime(node *CacheNode ) {
    lru.removeNode(node)
    lru.appendNode(node)
}

func (lru *LRUCache) get( key int) int {
    node := lru.m[key];
    if node == nil {
      return -1;
    }    
    lru.updateAccessTime(node)
    return node.value
}

  func (lru *LRUCache) set(key int,value int) {
    node := lru.m[key]
    if node != nil {
    	//fmt.Println("Should not enter here ")
      node.value = value;
      lru.updateAccessTime(node);
    } else {
      if lru.nodeNum == lru.capacity {
        delete(lru.m,lru.head.next.key)  // remove entry from map as well.
        lru.removeNode(lru.head.next)
      }
      node := NewCacheNode(key, value)
      lru.appendNode(node)
      lru.m[key] = node
    }
 }

 	//Create a new cache
var c *LRUCache


func Set(key int, value int) {
	// TODO: add your code here!
	if c == nil {

		c = NewLRUCache(CACHE_SIZE)
	}

	c.set(key,value)
}

func Get(key int) int {
	// TODO: add your code here!
	return(c.get(key))
}
