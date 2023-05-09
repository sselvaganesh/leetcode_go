package tree

/*
199. Binary Tree Right Side View
Medium
9.9K
597
Companies
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:
Input: root = [1,null,3]
Output: [1,3]

Example 3:
Input: root = []
Output: []
 

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func RightSideView(root *TreeNode) []int {
    
    if root==nil {
        return nil
    }

    var res []int

    nodes:=[]*TreeNode{root}
    for len(nodes)>0 {
        res=append(res, nodes[len(nodes)-1].Val)
        nodes=getNextLevel(nodes)
    }

    return res
}

func getNextLevel(nodes []*TreeNode) []*TreeNode {

    var res []*TreeNode
    for i:=0; i<len(nodes); i++ {
        if nodes[i].Left!=nil {
            res=append(res, nodes[i].Left)
        }

        if nodes[i].Right!=nil {
            res=append(res, nodes[i].Right)
        }
    }

    return res
}
