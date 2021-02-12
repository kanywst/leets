def repeatedNTimes(A):
    n = len(A)//2
    x = [i for i in set(A) if A.count(i)==n]
    return x[0]

def main():
    A = list(map(int,input().split(' ')))
    print(repeatedNTimes(A))

if __name__ == "__main__":
    main()