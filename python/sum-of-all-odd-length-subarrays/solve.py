def sumOddLengthSubarrays(arr):
    result = 0
    for i in range(len(arr)):
        for j in range(i,len(arr),2):
            tmp = arr[i:j+1]
            result += sum(tmp)
    return result

def main():
    arr = list(map(int,input().split(' ')))
    print(sumOddLengthSubarrays(arr))

if __name__ == "__main__":
    main()