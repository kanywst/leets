def lastStoneWeight(stones):
    stones = list(reversed(sorted(stones)))
    while len(stones) > 1:
        coll = stones[0:1]
        diff = abs(stones[0]-stones[1])
        if diff != 0:
            stones = list(reversed(sorted(stones[2:] + [diff])))
        else:
            stones = list(reversed(sorted(stones[2:])))
    if stones != []:
        return stones[0]
    else:
        return 0

def main():
    stones = [2,2]
    print(lastStoneWeight(stones))

if __name__ == "__main__":
    main()