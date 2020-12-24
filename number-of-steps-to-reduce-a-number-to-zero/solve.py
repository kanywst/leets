def numberOfSteps(num):
    cnt = 0
    while True:
        if num==0:
            break
        elif num%2==0:
            num=num//2
            cnt+=1
        else:
            num-=1
            cnt+=1
    return cnt

def main():
    num = int(input())
    print(numberOfSteps(num))

if __name__ == "__main__":
    main()