def xorOperation(n, start):
    result = start
    for i in range(1,n):
        tmp = start + 2*i
        result ^= tmp
    return result

def main():
    n = int(input())
    start = int(input())
    print(xorOperation(n, start))

if __name__ == "__main__":
    main()