def arrayPairSum(nums):
    B = sorted(nums)
    result = 0
    for i in range(0,len(nums)-1,2):
        result += min(B[i],B[i+1])
    return result

def main():
    nums = list(map(int,input().split(' ')))
    print(arrayPairSum(nums))

if __name__ == "__main__":
    main()