package hash

import (
	"testing"
)

const (
	node1 = "node1"
	node2 = "node2"
	node3 = "node3"
	node4 = "node4"
)

func TestHashRing(t *testing.T) {

	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 100
	nodeWeight[node2] = 100
	nodeWeight[node3] = 100
	nodeWeight[node4] = 100
	replicas := 1

	hash := NewHashRing(replicas, nil)

	hash.AddNodes(nodeWeight)
	if hash.GetNode("1") != node4 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node4 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node4 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("3"))
	}
	if hash.GetNode("4") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("4"))
	}

	hash.RemoveNode(node4)
	if hash.GetNode("1") != node3 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node3 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node3 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("3"))
	}
	if hash.GetNode("4") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("4"))
	}

	hash.AddNode(node4, 100)
	if hash.GetNode("1") != node4 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node4 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node4 {
		t.Fatalf("expetcd %v got %v", node4, hash.GetNode("3"))
	}
	if hash.GetNode("4") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("4"))
	}
}
