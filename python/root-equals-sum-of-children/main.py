def main():
    roots = list(map(int,input().split()))
    print(checkTree(roots))

def checkTree(root):
    child = root[1] + root[2]
    if child == 10:
        ans = True
    else:
        ans = False
    return ans

if __name__ == "__main__":
    main()