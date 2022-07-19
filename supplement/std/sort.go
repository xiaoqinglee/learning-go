package std

//1. type sort.Interface
//
//	type Interface interface {
//
//		Len() int
//
//		//If both Less(i, j) and Less(j, i) are false,
//		//then the elements at index i and j are considered equal.
//
//		//Less must describe a transitive ordering:
//		// - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//		// - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//		Less(i, j int) bool
//
//		// Swap swaps the elements with indexes i and j.
//		Swap(i, j int)
//	}

//2. 排序
//
//	func Float64s(x []float64)
//	func Ints(x []int)
//	func Strings(x []string)
//	func SliceStable(x any, less func(i, j int) bool) //func Slice(x any, less func(i, j int) bool)
//	func Stable(data Interface) //func Sort(data Interface)

//3. 验证顺序
//
//	func Float64sAreSorted(x []float64) bool
//	func IntsAreSorted(x []int) bool
//	func StringsAreSorted(x []string) bool
//	func SliceIsSorted(x any, less func(i, j int) bool) bool
//	func IsSorted(data Interface) bool

//4. 二分搜索
//
//	func SearchFloat64s(a []float64, x float64) int
//	func SearchInts(a []int, x int) int
//	func SearchStrings(a []string, x string) int
//
//	Search uses binary search to find and return the smallest index i in [0, n) at which f(i) is true,
//	assuming that on the range [0, n), f(i) == true implies f(i+1) == true.
//	That is, Search requires that f is false for some (possibly empty) prefix of the input range [0, n)
//	and then true for the (possibly empty) remainder;
//	Search returns the first true index.
//	If there is no such index, Search returns n.
//
//	To complete the example above,
//	the following code tries to find the value x
//	in an integer slice data sorted in ascending order:
//
//		x := 23
//		i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
//		if i < len(data) && data[i] == x {
//		// x is present at data[i]
//		} else {
//		// x is not present in data,
//		// but i is the index where it would be inserted.
//		}
//
//	func Search(n int, f func(int) bool) int

//5. 逆序函数(不排序, 只是重新定义Less(), 返回新Interface)
//
//	Reverse returns the reverse order for data.
//	func Reverse(data Interface) Interface
//
// 注意reverse Less()的定义,
// return r.Interface.Less(j, i) 和 return ! r.Interface.Less(i, j) 是不一样的,
// 后者会让逆序后的稳定排序算法在业务逻辑上失去稳定性, 前者不会.
//
// // Less returns the opposite of the embedded implementation's Less method.
// func (r reverse) Less(i, j int) bool {
// 	return r.Interface.Less(j, i)
// }

//6. 将slice转换为拥有Interface默认实现的slice
//
//	type Float64Slice
//	type IntSlice
//	type StringSlice

func TestSortInts() {

}
func TestSortSliceStable() {

}

func TestSortStable() {

}

func TestSortSearch() {

}

func TestReverse() {

}

//https://learnku.com/articles/38269
