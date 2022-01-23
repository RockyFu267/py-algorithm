def select_sort(listSet):
    for i in range(len(listSet)-1):
        minNum = listSet[i]
        for j in range(i+1,len(listSet)):
            if minNum > listSet[j]:
                minNum = listSet[j]
                listSet[i] , listSet[j] = listSet[j], listSet[i]
        print(listSet)

listSet = [9,8,7,1,2,3,4,5,6]
print(listSet)
print("---")
select_sort(listSet)
print("---")
print(listSet)