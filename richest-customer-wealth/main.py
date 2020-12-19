def maximumWealth(accounts):
    tmp = 0
    for i in accounts:
        s = sum(i)
        if tmp < s:
            tmp = s
    else:
        return tmp

def main():
    accounts = [[1,2,3],[3,2,1]]
    print(maximumWealth(accounts))

if __name__ == "__main__":
    main()