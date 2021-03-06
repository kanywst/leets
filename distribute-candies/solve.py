def distributeCandies(candyType):
    eatcnt = len(candyType)//2
    l = len(set(candyType))
    if eatcnt <= l:
        return eatcnt
    else:
        return l

def main():
    candyType = list(map(int,input().split(' ')))
    print(distributeCandies(candyType))

if __name__ == "__main__":
    main()