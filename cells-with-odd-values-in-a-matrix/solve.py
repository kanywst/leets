def oddCells(n,m,indices):
    result = [[0]*m for i in range(n)]
    cnt=0
    for i in indices:
        x,y = i[0],i[1]
        for j in range(m):
            result[x][j] += 1
        for j in range(n):
            result[j][y] += 1
    for i in range(len(result)):
        for j in range(len(result[0])):
            if result[i][j]%2 != 0:
                cnt+=1
    return cnt

def main():
    n = int(input())
    m = int(input())
    indices = [list(map(int,input().split(' '))) for _ in range(n)]
    print(oddCells(n,m,indices))

if __name__ == "__main__":
    main()