def numberOfMatches(n):
    matches = 0
    while True:
        if n==1:
            break
        elif n%2==0:
            matches += n//2
            n = n//2
        else:
            matches += (n-1)//2
            n = (n-1)//2+1
    return matches


def main():
    n = int(input())
    print(numberOfMatches(n))

if __name__ == "__main__":
    main()