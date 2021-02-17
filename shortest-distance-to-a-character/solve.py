def shortestToChar(s,c):
    index_nums = []
    for i, letter in enumerate(s):
        if letter == c:
            index_nums.append(i)
    ans = []
    l  = len(s)
    l2 = len(index_nums)
    d=[None]*l2
    for j in range(l):
        for k in range(l2):
            d[k] = abs(j-index_nums[k])
        ans.append(min(d))
    return ans


def main():
    s = input()
    c = input()
    print(shortestToChar(s,c))

if __name__ == "__main__":
    main()