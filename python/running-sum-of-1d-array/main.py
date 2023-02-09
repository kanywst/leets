def runningSum(nums):
    result = []
    for i in range(1,len(nums)+1):
        result.append(sum(nums[:i]))
    return result
def main():
    print(runningSum([1,2,3,4]))

if __name__ == "__main__":
    main()