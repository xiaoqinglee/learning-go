package ch6

import "fmt"

//如果方法的receiver有nil作为零值的情景, 那么设计这个方法的时候要考虑receiver实参等于nil的情景.

type IntList struct {
	nodeValue int
	next      *IntList
}

func (listHead *IntList) Sum() int {
	//如果链表node个数大于零, 那么最后一个node的next字段为值为nil的指针
	//如果链表的node个数等于零, 那么listHead == nil
	if listHead == nil {
		return 0
	} else {
		return listHead.nodeValue + listHead.next.Sum()
	}
}

func ReceiverNil() {

	////实例化了一个node
	//linkedList := new(IntList)
	//fmt.Printf("%#v\n", linkedList) //&ch6.IntList{nodeValue:0, next:(*ch6.IntList)(nil)}
	////实例化了一个node
	//linkedList := &IntList{}
	//fmt.Printf("%#v\n", linkedList) //&ch6.IntList{nodeValue:0, next:(*ch6.IntList)(nil)}

	//不会实例化一个node
	var linkedList *IntList
	fmt.Printf("%#v\n", linkedList) //(*ch6.IntList)(nil)

	//linkedList: []
	fmt.Printf("sum: %d\n", linkedList.Sum()) //10
	for i := 4; i >= 1; i-- {
		linkedList = &IntList{
			nodeValue: i,
			next:      linkedList,
		}
	}
	//linkedList: [1,2,3,4]
	fmt.Printf("sum: %d\n", linkedList.Sum()) //10
}
