def twoSum(nums, target):
    hashmap = {}
    for i, num in enumerate(nums):
        search_key = target - num
        if search_key not in hashmap:
            hashmap[num] = i
        else:
            return [i, hashmap[search_key]]

def main():
    lst = [3,3]
    print(twoSum(lst, 6))
if __name__ == "__main__":
    main()