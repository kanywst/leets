def decode(encoded,first):
    tmp = first
    result = []
    for i in range(len(encoded)):
        result.append(tmp)
        tmp = encoded[i]^tmp
    else:
        result.append(tmp)
    return result

def main():
    decoded = list(map(int,input().split(' ')))
    first = int(input())
    print(decode(decoded,first))

if __name__ == "__main__":
    main()