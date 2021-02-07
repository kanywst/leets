def sumOfUnique(nums):
    result = [i for i in set(nums) if nums.count(i) == 1]
    return sum(result)
    
def main():
    nums = list(map(int,input().split(' ')))
    print(sumOfUnique(nums))

if __name__ == "__main__":
    main()