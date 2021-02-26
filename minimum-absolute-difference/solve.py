def minimumAbsDifference(arr):
    arr.sort()
    m = 10000000000
    result = []
    for i in range(len(arr)-1):
        if m > arr[i+1] - arr[i]:
            m = arr[i+1] - arr[i]
    for i in range(len(arr)-1):
        if m == arr[i+1] - arr[i]:
            result.append([arr[i],arr[i+1]])
    return result

def main():
    arr = list(map(int,input().split(' ')))
    print(minimumAbsDifference(arr))

if __name__ == "__main__":
    main()