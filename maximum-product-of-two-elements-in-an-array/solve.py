from itertools import combinations

def maxProduct(nums):
    result = 0
    for i in combinations(nums,2):
        tmp = (i[0]-1) * (i[1]-1)
        result = max(result, tmp)
    return result

def main():
    nums = list(map(int, input().split(' ')))
    print(maxProduct(nums))

if __name__ == "__main__":
    main()