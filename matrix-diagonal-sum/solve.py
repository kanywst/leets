def diagonalSum(mat):
    result = 0
    for i in range(len(mat)):
        result += mat[i][i]
    for i in range(len(mat)):
        j = len(mat)-i-1
        result += mat[i][j]
    if len(mat)%2!=0:
        tmp = len(mat)//2
        result -= mat[tmp][tmp]
    return result
def main():
    mat = [[1,2,3],
           [4,5,6],
           [7,8,9]]
    print(diagonalSum(mat))

if __name__ == "__main__":
    main()