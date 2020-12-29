def decompressRLElist(nums):
    result = []
    for i in range(0,len(nums),2):
        print(i)
        result += nums[i] * [str(nums[i+1])]
    else:
        return result

def main():
    nums = list(map(int,input().split()))
    print(nums)
    print(decompressRLElist(nums))

if __name__ == "__main__":
    main()