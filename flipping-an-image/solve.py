def flipAndInvertImage(A):
    result = []
    for i in A:
        t = []
        for j in reversed(i):
            tmp = j^1
            t.append(tmp)
        result.append(t)
    return result

def main():
    A = [[1,1,0],
         [1,0,1],
         [0,0,0]]
    print(flipAndInvertImage(A))

if __name__ == "__main__":
    main()