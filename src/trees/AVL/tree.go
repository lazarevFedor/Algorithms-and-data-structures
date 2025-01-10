package AVL

import (
	"fmt"
	"strconv"
)

type Node struct {
	Key    int
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

type Tree struct {
	Root *Node
	size int
}

// New creates and initializes a new empty AVL tree.
//
// Returns:
//   - *Tree: A pointer to the newly created AVL tree.
func New() *Tree {
	return &Tree{nil, 0}
}

// Size returns the number of nodes in the AVL tree.
//
// Returns:
//   - int: The total number of nodes in the tree.
func (t *Tree) Size() int {
	return t.size
}

// String returns the string representation of a Node.
//
// Returns:
//   - string: The string representation of the node's key.
func (node *Node) String() string {
	return fmt.Sprintf("%v", node.Key)
}

// comparator compares two integers and returns an integer indicating their relative order.
//
// Parameters:
//   - a: The first integer to compare.
//   - b: The second integer to compare.
//
// Returns:
//   - int: -1 if a < b, 1 if a > b, or 0 if a == b.
func comparator(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// maximumNode returns the node with the maximum key value in the subtree rooted at this node.
//
// Returns:
//   - *Node: The node with the maximum key value in the subtree, or nil if the node is nil.
func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// getHeight returns the height of the given node in the AVL tree.
//
// Parameters:
//   - node: The node whose height is to be determined.
//
// Returns:
//   - int: The height of the node. Returns -1 if the node is nil.
func (t *Tree) getHeight(node *Node) int {
	if node == nil {
		return -1
	}
	return node.Height
}

// updHeight updates the height of the given node in the AVL tree.
//
// Parameters:
//   - node: The node whose height is to be updated. If the node is nil, the function returns immediately.
func (t *Tree) updHeight(node *Node) {
	if node == nil {
		return
	}
	if t.getHeight(node.Left) >= t.getHeight(node.Right) {
		node.Height = t.getHeight(node.Left) + 1
	} else {
		node.Height = t.getHeight(node.Right) + 1
	}
}

// Put inserts a new key-value pair into the AVL tree or updates an existing key.
//
// Parameters:
//   - key: The key of the node to be inserted or updated.
//   - value: The value associated with the key.
func (t *Tree) Put(key, value int) {
	t.Root = t.insert(t.Root, key, value)
}

// insert inserts a new node with the specified key and value into the AVL tree,
// or updates the value of an existing node with the same key.
//
// Parameters:
//   - node: The current node in the tree where the insertion or update is to be performed.
//   - key: The key of the node to be inserted or updated.
//   - value: The value associated with the key to be inserted or updated.
//
// Returns:
//   - *Node: The root node of the subtree after the insertion or update, balanced to maintain AVL properties.
func (t *Tree) insert(node *Node, key, value int) *Node {
	if node == nil {
		t.size++
		return &Node{Key: key, Value: value}
	} else if comparator(key, node.Key) == -1 {
		node.Left = t.insert(node.Left, key, value)
	} else if comparator(key, node.Key) == 1 {
		node.Right = t.insert(node.Right, key, value)
	} else {
		node.Value = value
	}
	return t.balance(node)
}

// Remove deletes a node with the specified key from the AVL tree.
//
// Parameters:
//   - key: The key of the node to be removed.
func (t *Tree) Remove(key int) {
	t.Root = t.delete(t.Root, key)
}

// getBalanceFactor calculates the balance factor of a given node in the AVL tree.
//
// Parameters:
//   - node: The node for which the balance factor is to be calculated. If the node is nil, the function returns 0.
//
// Returns:
//   - int: The balance factor of the node, which is the difference between the heights of the left and right subtrees.
func (t *Tree) getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return t.getHeight(node.Left) - t.getHeight(node.Right)
}

// rotateRight performs a right rotation on the given node in the AVL tree.
//
// Parameters:
//   - node: The node on which the right rotation is to be performed. It must not be nil.
//
// Returns:
//   - *Node: The new root of the subtree after the right rotation.
func (t *Tree) rotateRight(node *Node) *Node {
	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node
	t.updHeight(node)
	t.updHeight(newRoot)
	return newRoot
}

// rotateLeft performs a left rotation on the given node in the AVL tree.
//
// Parameters:
//   - node: The node on which the left rotation is to be performed. It must not be nil.
//
// Returns:
//   - *Node: The new root of the subtree after the left rotation.
func (t *Tree) rotateLeft(node *Node) *Node {
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node
	t.updHeight(node)
	t.updHeight(newRoot)
	return newRoot
}

// balance ensures the given subtree is balanced according to AVL tree properties.
//
// Parameters:
//   - node: The root of the subtree to be balanced.
//
// Returns:
//   - *Node: The new root of the balanced subtree.
func (t *Tree) balance(node *Node) *Node {
	t.updHeight(node)
	balanceFactor := t.getBalanceFactor(node)
	if balanceFactor > 1 {
		if t.getBalanceFactor(node.Left) >= 0 {
			return t.rotateRight(node)
		}
		node.Left = t.rotateLeft(node.Left)
		return t.rotateRight(node)
	}
	if balanceFactor < -1 {
		if t.getBalanceFactor(node.Right) <= 0 {
			return t.rotateLeft(node)
		}
		node.Right = t.rotateRight(node.Right)
		return t.rotateLeft(node)
	}
	return node
}

// delete removes a node with the specified key from the subtree.
//
// Parameters:
//   - node: The root of the subtree where the node is to be deleted.
//   - key: The key of the node to be removed.
//
// Returns:
//   - *Node: The root of the subtree after the deletion, balanced to maintain AVL properties.
func (t *Tree) delete(node *Node, key int) *Node {
	if node == nil {
		return nil
	} else if comparator(key, node.Key) == -1 {
		node.Left = t.delete(node.Left, key)
	} else if comparator(key, node.Key) == 1 {
		node.Right = t.delete(node.Right, key)
	} else {
		t.size--
		if node.Left == nil {
			return t.balance(node.Right)
		} else if node.Right == nil {
			return t.balance(node.Left)
		} else {
			maxNode := node.maximumNode()
			node.Key = maxNode.Key
			node.Value = maxNode.Value
			node.Left = t.delete(node.Left, maxNode.Key)
		}
	}
	t.updHeight(node)
	return node
}

// PreOrderTravers performs a tree traversal in pre-order,
// where nodes are visited in the following order: root, left subtree, right subtree.
//
// Parameters:
//   - node: A pointer to the current node where traversal starts.
//
// Returns:
//   - string: A string representation of the node keys, separated by spaces.
func (t *Tree) PreOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += strconv.Itoa(node.Key) + " "
		str += t.PreOrderTravers(node.Left) + " "
		str += t.PreOrderTravers(node.Right) + " "
	}
	return str
}

// InOrderTravers performs a tree traversal in in-order,
// where nodes are visited in the following order: left subtree, root, right subtree.
//
// Parameters:
//   - node: A pointer to the current node where traversal starts.
//
// Returns:
//   - string: A string representation of the node keys, separated by spaces.
func (t *Tree) InOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += t.InOrderTravers(node.Left) + " "
		str += strconv.Itoa(node.Key) + " "
		str += t.InOrderTravers(node.Right) + " "
	}
	return str
}

// PostOrderTravers performs a tree traversal in post-order,
// where nodes are visited in the following order: left subtree, right subtree, root.
//
// Parameters:
//   - node: A pointer to the current node where traversal starts.
//
// Returns:
//   - string: A string representation of the node keys, separated by spaces.
func (t *Tree) PostOrderTravers(node *Node) string {
	var str string
	if node != nil {
		str += t.PostOrderTravers(node.Left) + " "
		str += t.PostOrderTravers(node.Right) + " "
		str += strconv.Itoa(node.Key) + " "
	}
	return str
}

// LevelOrderTravers performs a tree traversal in level-order (breadth-first search),
// where nodes are visited level by level, starting from the root.
//
// Parameters:
//   - root: A pointer to the root node of the tree.
//
// Returns:
//   - string: A string representation of the node keys, separated by spaces.
func (t *Tree) LevelOrderTravers(root *Node) string {
	var str string
	if root == nil {
		return str
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		str += strconv.Itoa(node.Key) + " "
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return str
}

// Output formats the tree as a string, visually representing its structure using prefixes
// to denote hierarchy.
//
// Parameters:
//   - node: A pointer to the current node.
//   - prefix: A string representing the current prefix (used for level alignment).
//   - isTail: A boolean indicating whether the current node is the last child at this level.
//   - str: A pointer to a string where the result will be appended.
//
// Returns:
//   - None. The result is appended to the provided string pointer.
func Output(node *Node, prefix string, isTail bool, str *string) {
	if str == nil || node == nil {
		return
	}
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "    "
		}
		Output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "    "
		}
		Output(node.Left, newPrefix, true, str)
	}
}
