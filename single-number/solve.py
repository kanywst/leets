from collections import Counter

def singleNumber(nums):
    c = Counter(nums)
    return c.most_common()[-1][0]

def main():
    nums = list(map(int,input().split(' ')))
    print(singleNumber(nums))

if __name__ == "__main__":
    main()