def peakIndexInMountainArray(arr):
    result = 0
    result_index = 0
    for i in range(len(arr)):
        if arr[i] > result:
            result = arr[i]
            result_index = i
    return result_index 
def main():
    arr = list(map(int,input().split(' ')))
    print(peakIndexInMountainArray(arr))

if __name__ == "__main__":
    main()