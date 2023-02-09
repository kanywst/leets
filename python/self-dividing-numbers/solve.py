def selfDividingNumbers(left,right):
    result = []
    for i in range(left,right+1):
        for j in str(i):
            if int(j)==0 or i%int(j)!=0:
                break
        else:
            result.append(i)
    return result

def main():
    left = int(input())
    right = int(input())
    print(selfDividingNumbers(left,right))

if __name__ == "__main__":
    main()