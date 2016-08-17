// Package minimax implements the minimax algorithm
// Minimax (sometimes MinMax or MM[1]) is a decision rule used in decision theory,
// game theory, statistics and philosophy for minimizing the possible loss for
// a worst case (maximum loss) scenario
// See for more details: https://en.wikipedia.org/wiki/Minimax
package minimax

import (
	"fmt"
	"strconv"
)

// Node represents an element in the decision tree
type Node struct {
	// Score is available when supplied by an evaluation function or when calculated
	Score      *int
	parent     *Node
	children   []*Node
	isOpponent bool

	// Data field can be used to store additional information by the consumer of the
	// algorithm
	Data interface{}
}

// New returns a new minimax structure
func New() Node {
	n := Node{isOpponent: false}
	return n
}

// GetBestChildNode returns the first child node with the matching score
func (node *Node) GetBestChildNode() *Node {
	for _, cn := range node.children {
		if cn.Score == node.Score {
			return cn
		}
	}

	return nil
}

// Evaluate runs through the tree and caculates the score from the terminal nodes
// all the the way up to the root node
func (node *Node) Evaluate() {
	for _, cn := range node.children {
		if !cn.isTerminal() {
			cn.Evaluate()
		}

		if cn.parent.Score == nil {
			cn.parent.Score = cn.Score
		} else if cn.isOpponent && *cn.Score > *cn.parent.Score {
			cn.parent.Score = cn.Score
		} else if !cn.isOpponent && *cn.Score < *cn.parent.Score {
			cn.parent.Score = cn.Score
		}
	}
}

// Print the node for debugging purposes
func (node *Node) Print(level int) {
	var padding = ""
	for j := 0; j < level; j++ {
		padding += " "
	}

	var s = ""
	if node.Score != nil {
		s = strconv.Itoa(*node.Score)
	}

	fmt.Println(padding, node.isOpponent, node.Data, "["+s+"]")

	for _, cn := range node.children {
		level += 2
		cn.Print(level)
		level -= 2
	}
}

// AddTerminal adds a terminal node (or leave node).  These nodes
// should contain a score and no children
func (node *Node) AddTerminal(score int, data interface{}) *Node {
	return node.add(&score, data)
}

// Add a new node to structure, this node should have children and
// an unknown score
func (node *Node) Add(data interface{}) *Node {
	return node.add(nil, data)
}

func (node *Node) add(score *int, data interface{}) *Node {
	childNode := Node{parent: node, Score: score, Data: data}

	childNode.isOpponent = !node.isOpponent
	node.children = append(node.children, &childNode)
	return &childNode
}

func (node *Node) isTerminal() bool {
	return len(node.children) == 0
}
