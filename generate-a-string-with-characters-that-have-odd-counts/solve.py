def generateTheString(n):
    result = ''
    if n%2==0:
        result = 'a'*(n-1) + 'b'
    else:
        result = 'a'*n
    return result

def main():
    n = int(input())
    print(generateTheString(n))

if __name__ == "__main__":
    main()