def sortString(s):
    s = list(s)
    ans = []
    while s:
        ss = list(set(s))
        ss.sort(key=lambda c: ord(c))
        for i in ss:
            ans.append(i)
            s.remove(i)
        ss = list(set(s))
        ss.sort(key=lambda c: ord(c),reverse=True)
        for i in ss:
            ans.append(i)
            s.remove(i)
    return ''.join(ans)

def main():
    s = input()
    print(sortString(s))

if __name__ == "__main__":
    main()