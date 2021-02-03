def hammingDistance(x,y):
    return bin(x^y).count("1")

def main():
    x,y = map(int,input().split(' '))
    print(hammingDistance(x,y))

if __name__ == "__main__":
    main()