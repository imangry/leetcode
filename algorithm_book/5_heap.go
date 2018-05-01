package algorithm_book

/*
1.二叉堆：披着二叉树羊皮的数组，完全二叉树
2.1parent = (son-1)/2
2.2left = parent*2+1
2.3right = parent*2+2
3.堆的操作：
创建堆
SiftUp
SiftDown
*/




func maxHeapify(list []int, i int) {
	tmp := list[i]
	for k := 2*i + 1; k < len(list); k = 2*k + 1 {
		//找到当前parent的两个子节点值最大的那个
		if k+1 < len(list) && list[k+1] > list[k] {
			k++
		}
		//值都比parent小，说明堆结构是对的，跳出循环
		if list[k] <= tmp {
			break
		}
		list[i] = list[k]
		i = k
	}
	list[i] = tmp

}
