def countNegatives(grid):
    cnt = 0
    for i in grid:
        for j in i:
            if j < 0:
                cnt += 1
    return cnt

def main():
    grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
    print(countNegatives(grid))

if __name__ == "__main__":
    main()