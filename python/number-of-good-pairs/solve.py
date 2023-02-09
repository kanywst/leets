def numIdenticalPairs(nums):
    cnt = 0
    for i in range(len(nums)):
        for j in range(i,len(nums)):
            if i != j and nums[i] == nums[j]:
                cnt+=1
    else:
        return cnt 

def main():
    print(numIdenticalPairs([1,2,3]))

if __name__ == "__main__":
    main()