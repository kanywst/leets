def numJewelsInStones(jewels,stones):
    cnt = 0
    for i in jewels:
        cnt += stones.count(i)
    return cnt

def main():
    jewels, stones = input().split()
    print(numJewelsInStones(jewels, stones))

if __name__ == "__main__":
    main()