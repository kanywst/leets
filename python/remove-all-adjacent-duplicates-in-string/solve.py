def removeDuplicates(S):
    result = []
    for c in S:
        if result and result[-1] == c:
            result.pop()
        else:
            result.append(c)
    return ''.join(result)

def main():
    S = input()
    print(removeDuplicates(S))

if __name__ == "__main__":
    main()