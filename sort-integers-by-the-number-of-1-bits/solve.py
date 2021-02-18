def sortByBits(arr):
    arr.sort()
    result = []
    i = 0
    while True:
        tmp = [x for x in arr if bin(x).count("1") == i]
        if len(result) == len(arr):
            break
        else:
            i+=1
            result += tmp
    return result

def main():
    arr = list(map(int,input().split(',')))
    print(sortByBits(arr))

if __name__ == "__main__":
    main()