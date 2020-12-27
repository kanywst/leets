def subtractProductAndSum(n):
    p = eval('*'.join(list(str(n))))
    s = eval('+'.join(list(str(n))))
    return p-s

def main():
    n = int(input())
    print(subtractProductAndSum(n))

if __name__ == "__main__":
    main()