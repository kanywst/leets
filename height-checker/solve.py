def heightChecker(heights):
    h = sorted(heights)
    cnt = 0
    for i in range(len(heights)):
        if heights[i] != h[i]:
            cnt += 1
    return cnt

def main():
    heights = list(map(int,input().split(' ')))
    print(heightChecker(heights))

if __name__ == "__main__":
    main()