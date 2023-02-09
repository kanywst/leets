def halvesAreAlike(s):
    h = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']
    s1 = s[len(s)//2:]
    s2 = s[:len(s)//2]
    s1cnt = 0
    s2cnt = 0
    for i in h:
        s1cnt += s1.count(i)
        s2cnt += s2.count(i)
    if s1cnt==s2cnt:
        return True
    else:
        return False

def main():
    s = input()
    print(halvesAreAlike(s))

if __name__ == "__main__":
    main()