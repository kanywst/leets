from collections import Counter

def commonChars(A):
    compare = Counter(A[0])

    for i in range(len(A)):
        cnt = Counter(A[i])
        compare = compare & cnt
    answer = list(compare.elements())

    return answer

def main():
    A = input().split(',')
    print(commonChars(A))

if __name__ == "__main__":
    main()