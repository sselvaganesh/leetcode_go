package tree

/*
297. Serialize and Deserialize Binary Tree
Hard
8.8K
317
Companies
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

Example 1:
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]

Example 2:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 104].
-1000 <= Node.val <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	s := []string{}
	this.serializeHelper(root, &s)
	return strings.Join(s, ",")
}
func (this *Codec) serializeHelper(node *TreeNode, s *[]string) {
	if node == nil {
		*s = append(*s, "NIL")
		return
	}
	*s = append(*s, strconv.Itoa(node.Val))
	this.serializeHelper(node.Left, s)
	this.serializeHelper(node.Right, s)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(s string) *TreeNode {
	if s == "NIL" {
		return nil
	}

	strs := strings.Split(s, ",")
	nodes := make([]*TreeNode, len(strs))
	for i := 0; i < len(strs); i++ {
		if strs[i] != "NIL" {
			n, _ := strconv.Atoi(strs[i])
			nodes[i] = &TreeNode{
				Val: n,
			}
		}
	}

	return this.deserializeHelper(&nodes)
}

func (this *Codec) deserializeHelper(nodes *[]*TreeNode) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	node := (*nodes)[0]
	*nodes = (*nodes)[1:len(*nodes)]
	if node == nil {
		return node
	}
	node.Left = this.deserializeHelper(nodes)
	node.Right = this.deserializeHelper(nodes)
	return node
}
