def sortArrayByParity(A):
    result = []
    for i in A:
        if i%2==0:
            result.append(i)
    for i in A:
        if i%2!=0:
            result.append(i)
    return result

def main():
    A = list(map(int,input().split(' ')))
    print(sortArrayByParity(A))

if __name__ == "__main__":
    main()