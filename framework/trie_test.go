package framework

import (
	"testing"
)

func Test_filterChildNodes(t *testing.T) {
	root := &node{
		isLast:  false,
		segment: "",
		handler: func(*Context) error { return nil },
		childs: []*node{
			{
				isLast:  true,
				segment: "FOO",
				handler: func(*Context) error { return nil },
				childs:  nil,
			},
			{
				isLast:        false,
				isWildSegment: true,
				segment:       ":id",
				handler:       nil,
				childs:        nil,
			},
		},
	}

	{
		nodes := root.filterChild("FOO")
		if len(nodes) != 2 {
			t.Error("foo error")
		}
	}

	{
		nodes := root.filterChild(":FOO")
		if len(nodes) != 2 {
			t.Error(":foo error")
		}
	}

}

func Test_matchNode(t *testing.T) {
	root := &node{
		isLast:  false,
		segment: "",
		handler: func(*Context) error { return nil },
		childs: []*node{
			{
				isLast:  true,
				segment: "foo",
				handler: nil,
				childs: []*node{
					&node{
						isLast:  true,
						segment: "bar",
						handler: func(*Context) error { panic("not implemented") },
						childs:  []*node{},
					},
				},
			},
			{
				isLast:        true,
				segment:       ":id",
				isWildSegment: true,
				handler:       nil,
				childs:        nil,
			},
		},
	}

	{
		node := root.matchNode("foo/bar")
		if node == nil {
			t.Error("match normal node error")
		}
	}

	{
		node := root.matchNode("test")
		if node == nil {
			t.Error("match test")
		}
	}

}
