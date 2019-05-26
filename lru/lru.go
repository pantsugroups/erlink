package lru

type Node struct {
	Key   uint32
	Value int
	pre   *Node
	next  *Node
}

type LRUCache struct {
	limit   int
	HashMap map[uint32]*Node
	head    *Node
	end     *Node
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{limit: capacity}
	lruCache.HashMap = make(map[uint32]*Node, capacity)
	return lruCache
}

func (l *LRUCache) Get(key uint32) int {
	if v, ok := l.HashMap[key]; ok {
		l.refreshNode(v)
		return v.Value
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key uint32, value int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(l.head)
			delete(l.HashMap, oldKey)
		}
		node := Node{Key: key, Value: value}
		l.addNode(&node)
		l.HashMap[key] = &node
	} else {
		v.Value = value
		l.refreshNode(v)
	}
}

func (l *LRUCache) refreshNode(node *Node) {
	if node == l.end {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

func (l *LRUCache) removeNode(node *Node) uint32 {
	if node == l.end {
		l.end = l.end.pre
	} else if node == l.head {
		l.head = l.head.next
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return node.Key
}

func (l *LRUCache) addNode(node *Node) {
	if l.end != nil {
		l.end.next = node
		node.pre = l.end
		node.next = nil
	}
	l.end = node
	if l.head == nil {
		l.head = node
	}
}
