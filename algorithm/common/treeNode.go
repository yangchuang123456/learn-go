package common


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func creatCompleteTree(nodeNumber,index int) *TreeNode{
	var tree TreeNode
	if index < nodeNumber{
		tree.Left = creatCompleteTree(nodeNumber,2*index+1)
		tree.Right = creatCompleteTree(nodeNumber,2*index+2)
	}else{
		return nil
	}
	return &tree
}

func GetTestTree(slice []*int) *TreeNode{
	//先构造完全二叉树
	tree := creatCompleteTree(len(slice),0)

	q := NewQueue()
	q.Enqueue(tree)
	//层次遍历填入数据,并删除空节点
	i := 0
	for !q.IsEmpty(){
		tmpNode := q.Dequeue()
		if tmpNode != nil{
			node := tmpNode.(*TreeNode)
			node.Val = *slice[i]

			//处理左子树
			if 2*i+1 < len(slice){
				if slice[2*i+1] == nil{
					node.Left = nil
					q.Enqueue(nil)
				}else{
					q.Enqueue(node.Left)
				}
			}

			//处理右子树
			if 2*i+2 < len(slice){
				if slice[2*i+2] == nil{
					node.Right = nil
					q.Enqueue(nil)
				}else{
					q.Enqueue(node.Right)
				}
			}
		}
		i++
	}

	return tree
}