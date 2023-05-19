package main
import "fmt"
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortRecursive(arr, 0, len(arr)-1)
}
func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		quickSortRecursive(arr, low, pivot-1)
		quickSortRecursive(arr, pivot+1, high)
	}
}
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
func main() {
	arr := []int{10, 90, 40, 60, 50, 80, 30, 20, 70}
	fmt.Println("Original array:", arr)
	quickSort(arr)
	fmt.Println("Sorted array:", arr)
}
