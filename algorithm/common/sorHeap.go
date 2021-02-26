package common

type SortHeap []int

func (s SortHeap) Len() int {
	return len(s)
}

func (s SortHeap) Less(i, j int) bool{
	return false
}

func (s SortHeap) Swap(i, j int){
	s[i],s[j] = s[j],s[i]
}

func (s *SortHeap) Push(x interface{}){
	*s=append(*s,x.(int))
}

func (s *SortHeap) Pop() interface{}{
	res := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]

	return res
}
