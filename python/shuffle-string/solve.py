def restoreString(s,indices):
    result = ['A']*len(s)
    for i in range(len(indices)):
        result[indices[i]] = s[i]
    else:
        return ''.join(result)

def main():
    s = input()
    indices = list(map(int,input().split(' ')))

    print(restoreString(s,indices))

if __name__ == "__main__":
    main()