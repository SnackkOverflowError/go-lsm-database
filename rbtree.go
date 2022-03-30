package main

type DataType string

const (
	INT    DataType = "int"
	STRING DataType = "string"
	FLOAT  DataType = "float"
)

type Node struct {
	key        string
	data       []byte
	dataType   DataType
	isRed      bool
	leftChild  *Node
	rightChild *Node
}

func (node *Node) Insert(parent *Node, child *Node) bool {

	// if the node exists, simply update the node
	if node.key == child.key {
		node.data = child.data
		node.dataType = child.dataType
		return true
	}

	if node.key > child.key {
		if node.leftChild != nil {
			return node.leftChild.Insert(node, child)
		}
		// do some fixing and inserting
	}

	if node.key < child.key {
		if node.rightChild != nil {
			return node.rightChild.Insert(node, child)
		}
		// do some fixing and inserting
		return node.insertRight(parent, child)

	}

	return false
}

func (node *Node) Delete(parent *Node, key string) {

}

func (node *Node) Get(key string) *Node {
	if key == node.key {
		return node
	}

	if key > node.key {
		if node.rightChild == nil {
			return nil
		}
		return node.rightChild.Get(key)
	}

	if node.leftChild == nil {
		return nil
	}
	return node.leftChild.Get(key)
}

// this inserts right and also handles
func (parent *Node) insertRight(grandParent *Node, child *Node) bool {
	return false
}
