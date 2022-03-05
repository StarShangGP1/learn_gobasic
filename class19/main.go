package main

import (
	"fmt"
)

func main() {
	//doubleLinkedList()
	//tree()
	SortBubble()
}

// doubleLinkedList -------------------------------
// 双链表 内存中非连续的、链式结构的数据空间

type LinkNode struct {
	data int
	prev *LinkNode
	next *LinkNode
}

func doubleLinkedList() {
	n1 := &LinkNode{data: 1}
	n5 := &LinkNode{data: 5}
	n10 := &LinkNode{data: 10}

	n1.next = n5
	n5.prev = n1
	n5.next = n10
	n10.prev = n5
	fmt.Println("输出 doubleLinkedList")
	rangeLink(n1)

	fmt.Println("插入节点")
	root := insertNode(n1, &LinkNode{data: 3})
	root = insertNode(root, &LinkNode{data: 7})
	root = insertNode(root, &LinkNode{data: 11})
	root = insertNode(root, &LinkNode{data: 0})
	root = insertNode(root, &LinkNode{data: -1})
	rangeLink(root)

	fmt.Println("删除节点")
	root = deleteNode(root, -1)
	root = deleteNode(root, 3)
	root = deleteNode(root, 0)
	root = deleteNode(root, 11)
	rangeLink(root)
}

func deleteNode(root *LinkNode, data int) *LinkNode {
	//空链表
	if root == nil {
		return nil
	}
	if root.data == data {
		// 删除第一个节点
		left := root
		root = root.next

		left.next = nil
		root.prev = nil

		// 只有一个节点的情况
		//if root.next == nil && root.prev == nil {
		//	return nil
		//}

		return root
	}

	tmpNode := root
	for {
		if tmpNode.next == nil {
			// 走到链表的尾部，仍然没有找到要删除的数据，直接返回原root
			return root
		} else {
			if tmpNode.next.data == data {
				// 找到节点，开始删除，删除完成后返回原root
				right := tmpNode.next
				tmpNode.next = right.next
				right.next.prev = tmpNode

				// 清理掉右手上的link，保证GC正常回收
				right.next = nil
				right.prev = nil

				return root
			}
		}
		tmpNode = tmpNode.next
	}
}

func insertNode(root, newNode *LinkNode) *LinkNode {
	tmpNode := root
	// 整个链表是空的情况，新增
	if root == nil {
		return newNode
	}

	// 在链表的头，添加节点
	if root.data >= newNode.data {
		newNode.next = tmpNode
		tmpNode.prev = newNode
		return newNode
	}

	for {
		if tmpNode.next == nil {
			// 已经到头，在链表的尾追加节点即可
			tmpNode.next = newNode
			newNode.prev = tmpNode
			return root
		} else {
			if tmpNode.next.data >= newNode.data {
				// 找到位置，在中间插入新节点
				newNode.prev = tmpNode
				newNode.next = tmpNode.next

				tmpNode.next = newNode
				tmpNode.next.prev = newNode

				return root
			}
		}
		tmpNode = tmpNode.next
	}
}

func rangeLink(root *LinkNode) {
	fmt.Println("从头到尾")
	tmpNode := root
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}

	fmt.Println("从尾到头")
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.prev == nil {
			break
		}
		tmpNode = tmpNode.prev
	}
}

// tree -------------------------------
// 树 由节点、关联关系形成的有序的、分层的、树状结构的结构

type TreeNode struct {
	data  int
	root  *TreeNode
	left  *TreeNode
	right *TreeNode
}

func tree() {
	n1 := &TreeNode{data: 51}
	n2 := &TreeNode{data: 35}
	n3 := &TreeNode{data: 65}

	n1.left = n2
	n1.right = n3

	n2.root = n1
	n3.root = n1

	fmt.Println("插入Node")
	root := insertTreeNode(n1, &TreeNode{data: 43})
	root = insertTreeNode(root, &TreeNode{data: 28})

	fmt.Println("删除Node")
	deleteTreeNode(root, 43)
	deleteTreeNode(root, 35)

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

// 中间用于帮助理解删除的是最后叶子节点的情况
func deleteNodeLeaf(root *TreeNode, data int) *TreeNode {
	leftRoot := root
	if leftRoot.data == data && leftRoot.left == nil && leftRoot.root == nil {
		leftRoot = leftRoot.root
		right := root
		if leftRoot.left == right {
			// 删除左边叶子
			leftRoot.left = nil
			right.right = nil
			return leftRoot
		} else {
			// 删除右边叶子
			leftRoot.right = nil
			right.right = nil
			return leftRoot
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
