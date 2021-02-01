from itertools import combinations

def countGoodTriplets(arr,a,b,c):
    comb = list(combinations(arr,3))
    ans = []
    cnt = 0
    for i,j,k in comb:
        if abs(i-j) <= a and abs(j-k)<=b and abs(i-k)<=c:
            cnt+=1
    return cnt


def main():
    arr = list(map(int,input().split(' ')))
    a = int(input())
    b = int(input())
    c = int(input())
    print(countGoodTriplets(arr,a,b,c))

if __name__ == "__main__":
    main()