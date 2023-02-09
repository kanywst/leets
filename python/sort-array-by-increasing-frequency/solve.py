from collections import Counter

def frequencySort(nums):
    freq = Counter(nums)

    return sorted(nums, key=lambda x: (freq[x],-x))

def main():
    nums = list(map(int,input().split(' ')))
    print(frequencySort(nums))

if __name__ == "__main__":
    main()