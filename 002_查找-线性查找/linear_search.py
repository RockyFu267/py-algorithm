def linear_search(dataSet,valueSet):
    for i in dataSet:
        if dataSet[i] == valueSet:
            print(i)
            return[i]
    
    return

linear_search([1,2,3,4,5],4)
