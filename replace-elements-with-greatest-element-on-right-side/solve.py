def replaceElements(arr):
    result = []
    for i in range(len(arr)):
        tmp = arr[i+1:]
        if tmp == []:
            result.append(-1)
        else:
            result.append(max(tmp))
    return result

def main():
    arr = list(map(int,input().split(' ')))
    print(replaceElements(arr))

if __name__ == "__main__":
    main()