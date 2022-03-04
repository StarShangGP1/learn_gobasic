package main

import (
	"fmt"
	"sync"
)

func main() {
	q := &Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println("Queue 队列 先进先出")
	fmt.Println("    _______\n ->  3 2 1 ->")
	fmt.Println(q.Pop())
	fmt.Println(q.data)

	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println("Stack 栈 先进后出")
	fmt.Println("   ——————\n <- 3 2 1｜")
	fmt.Println(s.Pop())
	fmt.Println(s.data)

	fmt.Println("Linked List 链表 内存中非连续的、链式结构的数据空间")
	LinkedList()
}

// Queue -------------------------------
// Queue 队列 先进先出
//    _______
// ->  3 2 1 ->
type Queue struct {
	// 实现一个线程安全的队列，Tips：加锁
	sync.Mutex
	data []interface{}
}

func (q *Queue) Push(data interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, data)
}

func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.data) > 0 {
		// 取第一个返回
		o := q.data[0]
		// 删除第一个元素
		q.data = q.data[1:]
		return o, true
	}
	return nil, false
}

// Stack -------------------------------
// Stack 栈 先进后出
//   ——————
//<- 3 2 1｜
type Stack struct {
	data []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.data = append([]interface{}{data}, s.data...)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) > 0 {
		o := s.data[0]
		s.data = s.data[1:]
		return o, true
	}
	return nil, false
}

// LinkNode -------------------------------
// 链表 内存中非连续的、链式结构的数据空间
// 1。。。2。。。3。。。4。。。5
type LinkNode struct {
	data int
	next *LinkNode
}

func LinkedList() {
	n1 := &LinkNode{
		data: 1,
		next: nil,
	}
	n2 := &LinkNode{
		data: 2,
		next: nil,
	}
	n3 := &LinkNode{
		data: 3,
		next: nil,
	}
	n4 := &LinkNode{
		data: 4,
		next: nil,
	}
	n6 := &LinkNode{
		data: 6,
		next: nil,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6
	fmt.Println("输出 LinkedList")
	rangeLink(n1, func(data interface{}) {})

	fmt.Println("插入5")
	n5 := &LinkNode{
		data: 5,
		next: nil,
	}
	insertNode(n1, n5)
	insertNode(n1, &LinkNode{
		data: 7,
		next: nil,
	})
	insertNode(n1, &LinkNode{
		data: 5,
		next: nil,
	})
	insertNode(n1, &LinkNode{
		data: 3,
		next: nil,
	})
	rangeLink(n1, func(data interface{}) {})

	fmt.Println("删除节点")
	n1 = deleteNode(n1, 3)
	n1 = deleteNode(n1, 5)
	n1 = deleteNode(n1, 7)
	n1 = deleteNode(n1, 1)
	n1 = deleteNode(n1, 6)
	rangeLink(n1, func(data interface{}) {})

}

func deleteNode(root *LinkNode, data int) *LinkNode {
	tmpNode := root
	if root != nil && root.data == data {
		if root.next == nil {
			return nil
		}
		right := root.next
		tmpNode.next = nil
		return right
	}
	for {
		if tmpNode.next == nil {
			break
		}
		right := tmpNode.next
		if right.data == data {
			// 找到要删除的节点，开始删除
			tmpNode.next = right.next
			right.next = nil
			return root
		}
		tmpNode = tmpNode.next
	}
	return root
}

func insertNode(root, newNode *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			if newNode.data > tmpNode.data {
				if tmpNode.next == nil {
					// 已经到结尾，直接追加
					tmpNode.next = newNode
				} else {
					if tmpNode.next.data >= newNode.data {
						// 找到合适位置，准备插入数据
						newNode.next = tmpNode.next
						tmpNode.next = newNode
						break
					}
				}
			}
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}

func rangeLink(root *LinkNode, f func(data interface{})) {
	tmpNode := root
	for {
		if tmpNode != nil {
			f(tmpNode.data)
			fmt.Println(tmpNode.data)
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}
