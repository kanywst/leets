from collections import Counter

def findLucky(arr):
    a = Counter(arr)
    result = []
    for i,j in a.items():
        if i == j:
            result.append(i)
    if result == []:
        ans = -1
    else:
        ans = max(result)
    return ans

def main():
    arr = list(map(int,input().split(' ')))
    print(findLucky(arr))

if __name__ == "__main__":
    main()