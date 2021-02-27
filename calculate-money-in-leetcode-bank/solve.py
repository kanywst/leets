def totalMoney(n):
    t1 = n//7
    t2 = n%7
    result = 0
    m = 1
    for i in range(1,t1+1):
        result += sum(range(i,i+7))
        m = i+1
    result += sum(range(m,m+t2))
    return result

def main():
    n = int(input())
    print(totalMoney(n))

if __name__ == "__main__":
    main()