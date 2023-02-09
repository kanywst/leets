
def shuffle(nums,n):
    x = nums[:n]
    y = nums[n:]
    result = []
    for i in range(n):
        result.append(x[i])
        result.append(y[i])
    return result

def main():
    nums = list(map(int,input().split(' ')))
    n = int(input())
    print(shuffle(nums,n))

if __name__ == "__main__":
    main()