package utils

//直接插入排序 稳定 额外空间 1
// 1.设置监视哨 r[i] = tmp
// 2.从r[1]开始比较
// 3.r[i] 和 r[i-1] 比较; r[i-1] >= tmp r[i]=tmp 即r[i-1]后移； r[i-1] < tmp r[i] = tmp
func InsertionSort(r []int) {
	length := len(r)
	for i := 1; i < length; i++ {
		tmp := r[i]
		for j := i - 1; j >= -1; j-- {
			if j == -1 {
				r[0] = tmp
				break
			}
			if tmp <= r[j] {
				r[j+1] = r[j]
			} else {
				r[j+1] = tmp
				break
			}
		}
	}
}

//没有监视哨
func InsertSort(r []int) {
	length := len(r)
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if r[j] >= r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
}

//希尔排序  不稳定
func ShellSort(r []int) {
	length := len(r)
	//区间
	var gap int = 1
	for gap < length {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			var tmp = r[i]
			var j = i - gap
			//跨区间排序
			for j >= 0 && r[j] > tmp {
				r[j+gap] = r[j]
				j -= gap
			}
			r[j+gap] = tmp
		}
		gap = int(gap / 3)
	}
}

//冒泡排序  稳定
func BubbleSort(r []int) {
	length := len(r)
	for i := 0; i < length-1; i++ {
		isChange := false
		for j := 0; j < length-1-i; j++ {
			if r[j] < r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}

//快速排序  不稳定 选定r[0]
func QuickSort(r []int) {
	length := len(r)
	if length <= 1 {
		return
	}
	mid, i := r[0], 1
	head, tail := 0, length-1
	for head < tail {
		if r[i] > mid {
			r[i], r[tail] = r[tail], r[i]
			tail--
		} else {
			r[i], r[head] = r[head], r[i]
			head++
			i++
		}
	}
	r[head] = mid
	QuickSort(r[:head])
	QuickSort(r[head+1:])
}

// 对牌值从小到大排序,采用快速排序算法
func Qsort(arr []byte) {
	len := len(arr)
	if len <= 1 {
		return
	}
	mid, i := arr[0], 1
	head, tail := 0, len-1
	if head < tail {
		if arr[i] > mid {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}
	arr[head] = mid
	Qsort(arr[:head])
	Qsort(arr[head+1:])
}

// 对牌值从小到大排序，采用快速排序算法
func QuSort(arr []byte, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			QuSort(arr, start, j)
		}
		if end > i {
			QuSort(arr, i, end)
		}
	}
}

//选择排序  不稳定
func SelectionSort(r []int) {
	length := len(r)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if r[j] < r[i] {
				r[j], r[i] = r[i], r[j]
			}
		}
	}
}

//归并排序  稳定
func MergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := MergeSort(r[:num])
	right := MergeSort(r[num:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

//基数排序 稳定
// MSD法:从优先级最高的关键字开始
// LSD法:从优先级最低的关键字开始,速度快,但对数据有类型要求
// 时间复杂度:O(d n) n为关键字数量
func RadixSort(sz []int) {
	// 链表的数据类型
	type ele struct {
		data int
		next *ele
	}

	// 将切片转化为链表
	n := len(sz)
	r := make([]ele, n+10) // 保管要排序的数据
	t := make([]*ele, 10)  // 记录各分段的段尾
	s := r[n:]             // 作为各分段的段头
	p := &r[0]             // 记录链表的头部
	for i := 0; i < n; i++ {
		r[i].data = sz[i]
		r[i].next = &r[i+1]
	}
	r[n-1].next = nil

	// 预制分成十段的链表，每段都有头尾（目前二者相同）
	initialize := func() {
		for i := 0; i < 9; i++ {
			t[i] = &s[i]
			s[i].next = &s[i+1]
		}
		t[9], s[9].next = &s[9], nil
	}

	// 将段中的预置的段头剔除出来，必须从后向前（请思考）
	separate := func() *ele {
		for i := 9; i > 0; i-- {
			t[i-1].next = s[i].next
		}
		return s[0].next
	}

	// 依次取出链表中元素顺序插入各段中
	insert := func(fac func(int) int) {
		for p != nil {
			q := p.next
			j := fac(p.data)
			p.next = t[j].next
			t[j].next = p
			t[j], p = p, q
		}
	}

	// 按照个位排序第一次
	initialize()
	insert(func(k int) int { return k % 10 })
	p = separate()

	// 按照十位排序第二次
	initialize()
	insert(func(k int) int { return k / 10 % 10 })
	p = separate()

	// 按照百位排序第三次
	initialize()
	insert(func(k int) int { return k / 100 % 10 })
	p = separate()

	// 从链表返回给数组
	for i := 0; p != nil; i, p = i+1, p.next {
		sz[i] = p.data
	}
}
