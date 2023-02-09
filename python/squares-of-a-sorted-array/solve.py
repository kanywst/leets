def ss(num):
    return num**2
def sortedSquares(nums):
    return sorted(list(map(ss,nums)))
def main():
    nums = list(map(int,input().split(' ')))
    print(sortedSquares(nums))
if __name__ == "__main__":
    main()