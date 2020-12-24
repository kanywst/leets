def smallerNumbersThanCurrent(nums):
    result = []
    for i in nums:
        tmp = 0
        for j in nums:
            if i > j:
                tmp+=1
        else:
            result.append(tmp)
    else:
        return result

def main():
    print(smallerNumbersThanCurrent(input().split()))

if __name__ == "__main__":
    main()