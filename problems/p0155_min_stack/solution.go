package p0155_min_stack

type MinStack struct {
	datas []Data
}

type Data struct {
	val    int
	minVal int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {

	var minVal = val
	if len(this.datas) > 0 {
		minVal = min(this.datas[len(this.datas)-1].minVal, minVal)
	}
	this.datas = append(this.datas, Data{
		val:    val,
		minVal: minVal,
	})
}

func (this *MinStack) Pop() {
	this.datas = this.datas[:len(this.datas)-1]
}

func (this *MinStack) Top() int {
	return this.datas[len(this.datas)-1].val
}

func (this *MinStack) GetMin() int {
	return this.datas[len(this.datas)-1].minVal
}
