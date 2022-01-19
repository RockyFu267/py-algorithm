import random


def bubble_sort(listSet):
    for i in range(len(listSet)-1):
        exchangeStatus = False
        for j in range(len(listSet)-i-1):
            if listSet[j] > listSet[j+1]:
                listSet[j], listSet[j+1] = listSet[j+1], listSet[j]
                exchangeStatus= True
        print(listSet)
        if not exchangeStatus:
            return
        


#listSet = [random.randint(0, 10000) for i in range(0, 10)]
listSet = [9,8,7,1,2,3,4,5,6]
print(listSet)
print("------------")
bubble_sort(listSet)
print("------------")
print(listSet)
