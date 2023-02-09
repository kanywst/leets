def sumZero(n):
    result = []
    if n%2!=0:
        result.append(0)
    for i in range(1,n//2+1):
        result.append(-i)
        result.append(i)
    return result


def main():
    n = int(input())
    print(sumZero(n))

if __name__ == "__main__":
    main()