def createTargetArray(nums,index):
    result = []
    for a in range(len(index)):
        i = index[a]
        result = result[:i] + [nums[a]] + result[i:]
    return result

def main():
    nums = list(map(int, input().split(' ')))
    index = list(map(int,input().split(' ')))
    print(createTargetArray(nums,index))

if __name__ == "__main__":
    main()