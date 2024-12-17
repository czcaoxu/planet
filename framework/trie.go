package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	isLast        bool   // uri的终点，表示一个完整的uri
	segment       string // uri
	isWildSegment bool
	handler       ControllerHandler
	childs        []*node
}

func newNode() *node {
	return new(node)
}

func (n *node) filterChild(segment string) []*node {
	if isWildSegment(segment) {
		return n.childs
	}

	var childs []*node
	for _, node := range n.childs {
		if node.isWildSegment || node.segment == segment {
			childs = append(childs, node)
		}
	}

	return childs
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]
	childs := n.filterChild(segment)
	if len(childs) == 0 {
		return nil
	}

	if len(segments) == 1 {
		for _, c := range childs {
			if c.isLast {
				return c
			}
		}

		return nil
	}

	for _, c := range childs {
		node := c.matchNode(segments[1])
		if node != nil {
			return node
		}
	}

	return nil
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func NewTree() *Tree {
	return &Tree{root: newNode()}
}

func (t *Tree) AddRouter(uri string, handler ControllerHandler) error {
	root := t.root
	if root.matchNode(uri) != nil {
		return errors.New("route exits: " + uri)
	}

	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		var currentNode *node
		childs := root.filterChild(segment)
		for _, n := range childs {
			if n.segment == segment {
				currentNode = n
				break
			}
		}

		if currentNode == nil {
			currentNode = new(node)
			currentNode.segment = segment
			if index == len(segments)-1 {
				currentNode.isLast = true
				currentNode.handler = handler
			}
			if isWildSegment(segment) {
				currentNode.isWildSegment = true
			}
			root.childs = append(root.childs, currentNode)
		}
		root = currentNode
	}

	return nil
}

func (t *Tree) FindHandler(uri string) ControllerHandler {
	node := t.root.matchNode(uri)
	if node != nil {
		return node.handler
	}

	return nil
}
