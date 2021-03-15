def maxPower(s):
    cnt = 1
    ans = 0
    for i in range(len(s)-1):
        if s[i] == s[i+1]:
            cnt += 1
        else:
            ans = max(cnt,ans)
            cnt = 1
    else:
        ans = max(cnt,ans)
    return ans

def main():
    s = input()
    print(maxPower(s))

if __name__ == "__main__":
    main()