def binary_search(dataSet, valueSet):
    left = 0
    right = len(dataSet)
    while left <= right:
        mid = (left + right) // 2
        if dataSet[mid] == valueSet:
            print(mid)
            return mid
        elif dataSet[mid] > valueSet:
            right = mid - 1
        elif dataSet[mid] < valueSet:
            left = mid + 1
    return

binary_search([1,2,3,4,5,6,7,8,9,10],9)
