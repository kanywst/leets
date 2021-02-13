def minDeletionSize(strs):
    cnt=0
    for i in zip(strs):
        if list(i) != sorted(i):
            cnt+=1
    return cnt


def main():
    strs = input().split(' ')
    print(minDeletionSize(strs))

if __name__ == "__main__":
    main()