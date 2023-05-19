def quickSort(arr):
    if len(arr) <= 1:
        return arr
    else:
        pivot = arr[len(arr) // 2]
        left = [x for x in arr if x < pivot]
        middle = [x for x in arr if x == pivot]
        right = [x for x in arr if x > pivot]
        return quickSort(left) + middle + quickSort(right)

arr = [10, 90, 40, 60, 50, 80, 30, 20, 70]
print("Original array:", arr)

sorted_arr = quickSort(arr)
print("Sorted array:", sorted_arr)
