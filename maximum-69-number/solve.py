import copy
def maximum69Number(num):
    s = list(str(num))
    tmp=num
    for i in range(len(s)):
        t = copy.copy(s)
        if t[i] == '6':
            t[i] = '9'
        else:
            t[i] = '6'
        tmp=max(tmp,int(''.join(t)))
    return tmp
def main():
    num = int(input())
    print(maximum69Number(num))

if __name__ == "__main__":
    main()