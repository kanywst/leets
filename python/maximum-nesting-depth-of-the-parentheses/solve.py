def maxDepth(s):
    result=0
    cnt=0
    for i in s:
        if i=='(':
            cnt+=1
            result=max(cnt,result)
        elif i==')':
            cnt-=1
            result=max(cnt,result)
    return result

def main():
    s = input()
    print(maxDepth(s))

if __name__ == "__main__":
    main()