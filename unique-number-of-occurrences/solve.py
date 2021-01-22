def uniqueOccurrences(arr):
    cnt = []
    result = True
    for i in set(arr):
        c = arr.count(i)
        if c in cnt:
            result = False
            break
        else:
            cnt.append(c)
    return result
def main():
    arr = list(map(int,input().split(' ')))
    print(uniqueOccurrences(arr))
if __name__ == "__main__":
    main()