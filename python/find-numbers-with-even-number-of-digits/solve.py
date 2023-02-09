def findNumbers(nums):
    result = 0
    for i in nums:
        if len(str(i))%2==0:
            result+=1
    else:
        return result

def main():
    nums = list(map(int,input().split(' ')))
    print(findNumbers(nums))

if __name__ == "__main__":
    main()