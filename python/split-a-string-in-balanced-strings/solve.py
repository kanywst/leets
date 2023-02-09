def balancedStringSplit(s):
    cnt = 0
    check = {'R': 0, 'L': 0}
    for i in s:
        check[i] += 1
        if check['R'] == check['L']:
            cnt+=1
            check['R'] = 0
            check['L'] = 0
    return cnt

def main():
    s = input()
    print(balancedStringSplit(s))

if __name__ == "__main__":
    main()