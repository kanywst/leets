def shortestToChar(s,c):
    index_nums = []
    for i, letter in enumerate(s):
        if letter == c:
            index_nums.append(i)
    result = []
    for i in range(len(s)):
        d = []
        for j in index_nums:
            d.append(abs(i-j))
        else:
            result.append(min(d))
    return result




def main():
    s = input()
    c = input()
    print(shortestToChar(s,c))

if __name__ == "__main__":
    main()