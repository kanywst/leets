def maximumUnits(boxTypes, truckSize):
    boxTypes.sort(key=lambda x:x[-1],reverse=True)
    print(boxTypes)
    ans = 0
    for i,j in boxTypes:
        if truckSize > i:
            ans += i*j
            truckSize -= i
        else:
            ans += truckSize*j
            break
    return ans


def main():
    boxTypes = [[1,3],[2,2],[3,1]]
    truckSize = 4
    print(maximumUnits(boxTypes,truckSize))

if __name__ == "__main__":
    main()