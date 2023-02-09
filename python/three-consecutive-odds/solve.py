def threeConsecutiveOdds(arr):
    cnt = 0
    odd_flag = False
    result = False
    for i in arr:
        if i%2==1:
            odd_flag = True
            cnt += 1
        else:
            odd_flag = False
            cnt = 0
        if cnt == 3:
            result = True
            break
    return result

def main():
    arr = list(map(int,input().split(' ')))
    print(threeConsecutiveOdds(arr))

if __name__ == "__main__":
    main()