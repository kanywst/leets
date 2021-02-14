from itertools import combinations
def canMakeArithmeticProgression(arr):
    arr.sort()
    tmp = arr[1] - arr[0]
    for i in range(1,len(arr)):
        if arr[i] - arr[i-1] != tmp:
            return False
    else:
        return True


def main():
    arr=list(map(int,input().split(' ')))
    print(canMakeArithmeticProgression(arr))

if __name__=="__main__":
    main()