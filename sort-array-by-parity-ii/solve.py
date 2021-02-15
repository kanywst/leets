def sortArrayByParityII(A):
    oddList = []
    evenList = []
    result = []
    for i in A:
        if i%2==0:
            evenList.append(i)
        else:
            oddList.append(i)
    for i in range(len(A)):
        if i%2==0:
            result.append(evenList.pop())
        else:
            result.append(oddList.pop())
    return result        

def main():
    A = list(map(int,input().split(' ')))
    print(sortArrayByParityII(A))

if __name__ == "__main__":
    main()