package code_lru

type LRUCache struct {
	Data  map[int]int
	Temp  []int
	Count int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Data:  make(map[int]int),
		Temp:  make([]int, 0),
		Count: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Data[key]; !ok {
		return -1
	}
	this.DealTemp(key)
	return this.Data[key]
}

func (this *LRUCache) DealTemp(key int) {
	if len(this.Temp) > 1 {
		for k, v := range this.Temp {
			if v == key {
				if k != len(this.Temp)-1 {
					_temp := this.Temp[k]
					j := k
					for j < len(this.Temp)-1 {
						this.Temp[j], this.Temp[j+1] = this.Temp[j+1], this.Temp[j]
						j++
					}
					this.Temp[len(this.Temp)-1] = _temp
				}
			}
		}
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Data[key]; !ok {
		if this.Count == len(this.Temp) {
			delete(this.Data, this.Temp[0])
			this.Temp = append(this.Temp[1:])
		}
		this.Temp = append(this.Temp, key)
	} else {
		this.DealTemp(key)
	}
	this.Data[key] = value
}
