def countConsistentStrings(allowed, words):
    cnt = 0
    for i in words:
        tmp = 0
        for j in list(allowed):
            tmp += i.count(j)
        else:
            if len(i) == tmp:
                cnt+=1
    return cnt

def main():
    allowed = input()
    words = input().split(' ')
    print(countConsistentStrings(allowed, words))

if __name__ == "__main__":
    main()