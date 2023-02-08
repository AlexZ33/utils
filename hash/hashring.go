//带权重的一致性hash

package hash

import (
	"hash/crc32"
	"math"
	"sort"
	"strconv"
	"sync"
)

const (
	// 默认虚拟节点200
	DefaultReplicas = 200
)

type Node struct {
	Key   string
	Value uint32
}

type Nodes []Node

//可自定义hash函数
type Hash func(data []byte) uint32

type HashRing struct {
	HashFunc Hash
	// 原始数据，所有的节点
	Nodes Nodes
	Mutex sync.RWMutex
	// 一致性hash 虚拟node，以及关联的节点
	NodeMap map[uint32]string
	// 原始数据，所有的节点权重
	Weights  map[string]int
	Replicas int
}

// 排序支持
func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Less(i, j int) bool { return n[i].Value < n[j].Value }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nodes) Sort()              { sort.Sort(n) }

func NewHashRing(replicas int, hashFunc Hash) *HashRing {
	if replicas == 0 {
		replicas = DefaultReplicas
	}
	//https://cloud.tencent.com/developer/section/1142447
	//默认使用crc32.ChecksumIEEE
	if hashFunc == nil {
		hashFunc = crc32.ChecksumIEEE
	}
	// 返回声明的一个对象
	return &HashRing{
		HashFunc: hashFunc,
		NodeMap:  make(map[uint32]string),
		Weights:  make(map[string]int),
		Replicas: replicas,
	}
}

//核心逻辑 生成虚拟node
func (h *HashRing) generate() {
	totalWeight := 0
	for _, weight := range h.Weights {
		totalWeight += weight
	}
	totalReplicas := h.Replicas * len(h.Weights)
	h.Nodes = make(Nodes, 0)

	for nodeKey, weight := range h.Weights {
		//使用权重计算虚拟节点数
		replicas := int(math.Floor(float64(weight) / float64(totalWeight) * float64(totalReplicas)))
		for i := 1; i <= replicas; i++ {
			hashValue := h.HashFunc([]byte(nodeKey + strconv.Itoa(i)))
			node := Node{
				Key:   nodeKey,
				Value: hashValue,
			}
			h.Nodes = append(h.Nodes, node)
			h.NodeMap[hashValue] = nodeKey
		}
	}
	h.Nodes.Sort()
}

// 添加节点并设置节点权重
func (h *HashRing) AddNode(nodeKey string, weight int) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	h.Weights[nodeKey] = weight
	h.generate()
}

// 添加多个节点与权重
func (h *HashRing) AddNodes(nodeWeight map[string]int) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	for nodeKey, weight := range nodeWeight {
		h.Weights[nodeKey] = weight
	}

	h.generate()
}

// 删除节点
func (h *HashRing) RemoveNode(nodeKey string) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	delete(h.Weights, nodeKey)
	h.generate()
}

// 删除多个节点
func (h *HashRing) RemoveNodes(nodeKeys []string) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	for _, nodeKey := range nodeKeys {
		delete(h.Weights, nodeKey)
	}

	h.generate()
}

// 获取指定字符串对应的节点名称
func (h *HashRing) GetNode(str string) string {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	if len(h.Nodes) == 0 {
		return ""
	}

	hash := h.HashFunc([]byte(str))
	idx := sort.Search(len(h.Nodes), func(i int) bool { return h.Nodes[i].Value >= hash })
	if idx == len(h.Nodes) {
		idx = 0
	}

	return h.NodeMap[h.Nodes[idx].Value]
}
