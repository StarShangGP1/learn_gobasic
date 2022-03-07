package main

import (
	"fmt"
	"time"
)

func treeSearch() {
	// tree code
	totalCompare = 0
	var root *TreeNode
	startTime := time.Now()
	for _, v := range sampleData {
		root = insertTreeNode(root, &TreeNode{data: int(v)})
	}
	buildFinishTime := time.Since(startTime)
	fmt.Println("构建结束", buildFinishTime)
	for i := 0; i < 100*10000; i++ {

		//beforeOrderTraversal(root, 501)
		//beforeOrderTraversal(root, 888)
		//beforeOrderTraversal(root, 900)
		//beforeOrderTraversal(root, 3)

		//inSequenceTraversal(root, 501)
		//inSequenceTraversal(root, 888)
		//inSequenceTraversal(root, 900)
		//inSequenceTraversal(root, 3)

		afterOrderTraversal(root, 501)
		afterOrderTraversal(root, 888)
		afterOrderTraversal(root, 900)
		afterOrderTraversal(root, 3)
	}
	finishTime := time.Since(startTime)
	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime)
}

// 前序遍历
func beforeOrderTraversal(root *TreeNode, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if root.data == targetNum {
		return true
	}
	if beforeOrderTraversal(root.left, targetNum) {
		return true
	}
	if beforeOrderTraversal(root.right, targetNum) {
		return true
	}
	return false
}

// 中序遍历
func inSequenceTraversal(root *TreeNode, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if inSequenceTraversal(root.left, targetNum) {
		return true
	}
	if root.data == targetNum {
		return true
	}

	if inSequenceTraversal(root.right, targetNum) {
		return true
	}
	return false
}

// 后序遍历
func afterOrderTraversal(root *TreeNode, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if afterOrderTraversal(root.left, targetNum) {
		return true
	}
	if afterOrderTraversal(root.right, targetNum) {
		return true
	}
	if root.data == targetNum {
		return true
	}
	return false
}

// tree -------------------------------
// 树 由节点、关联关系形成的有序的、分层的、树状结构的结构
// 左边最小，右边最大

type TreeNode struct {
	data  int
	root  *TreeNode
	left  *TreeNode
	right *TreeNode
}

func insertTreeNode(root, newNode *TreeNode) *TreeNode {
	// 整个树是空的情况，新增
	if root == nil {
		return newNode
	}
	if newNode.data == root.data {
		return root
	}
	if newNode.data < root.data {
		if root.left == nil {
			root.left = newNode
			newNode.root = root
		} else {
			insertTreeNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
			newNode.root = root
		} else {
			insertTreeNode(root.right, newNode)
		}
	}
	return root
}

func deleteTreeNode(root *TreeNode, data int) *TreeNode {

	// 左边
	if data < root.data {
		deleteTreeNode(root.left, data)
		// 右边
	} else if data > root.data {
		deleteTreeNode(root.right, data)
	} else {
		// 现在root指向要删除的节点
		leftNext := findNextGenFromLeft(root.left)
		rightNext := findNextGenFromRight(root.right)
		if leftNext == nil && rightNext == nil {
			// 现在要删除的是叶子结点，即最底部的节点
			top := root.root
			down := root
			if top.left == down {
				// 表示是左子树
				top.left = nil
				down.root = nil
				return nil
			} else {
				// 表示是右子树
				top.right = nil
				down.root = nil
				return nil
			}
		} else if leftNext != nil {
			root.data = leftNext.data
			deleteTreeNode(leftNext, leftNext.data)
		} else if rightNext != nil {
			root.data = rightNext.data
			deleteTreeNode(rightNext, rightNext.data)
		}

	}

	return root
}

func findNextGenFromLeft(root *TreeNode) *TreeNode {
	tmpNode := root
	if root == nil {
		return nil
	}
	for {
		if tmpNode.right != nil {
			tmpNode = tmpNode.right
		} else {
			break
		}
	}
	return tmpNode
}

func findNextGenFromRight(root *TreeNode) *TreeNode {
	tmpNode := root
	if root == nil {
		return nil
	}
	for {
		if tmpNode.left != nil {
			tmpNode = tmpNode.left
		} else {
			break
		}
	}
	return tmpNode
}
