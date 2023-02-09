import copy
def luckyNumbers(matrix):
    result = []
    for i in range(len(matrix)):
        tmp = []
        for j in range(len(matrix[0])):
            tmp.append(matrix[j][i])
        result.append(tmp)
    max_row,max_cow = [],[]
    for i in matrix:
        max_row.append(min(i))
    for i in result:
        max_cow.append(max(i))
    return list(set(set(max_cow)&set(max_row)))
def main():
    matrix = [[3,7,8],[9,11,13],[15,16,17]]
    print(luckyNumbers(matrix))
if __name__ == "__main__":
    main()