from collections import Counter

def findLucky(arr):
    a = Counter(arr)
    result = -1
    for i,j in a.items():
        if i == j:
            result = max(result,i)
    return result

def main():
    arr = list(map(int,input().split(' ')))
    print(findLucky(arr))

if __name__ == "__main__":
    main()