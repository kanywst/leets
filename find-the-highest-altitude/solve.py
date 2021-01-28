def largestAltitude(gain):
    result = [0]
    for i in range(len(gain)):
        result.append(gain[i]+result[-1])
    return max(result)

def main():
    gain = list(map(int,input().split(' ')))
    print(largestAltitude(gain))

if __name__ == "__main__":
    main()