def canBeEqual(target,arr):
    if sorted(target) == sorted(arr):
        return True
    else:
        return False
def main():
    target = list(map(int,input().split(' ')))
    arr = list(map(int,input().split(' ')))
    print(canBeEqual(target,arr))
if __name__ == "__main__":
    main()