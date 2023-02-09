def removeOuterParentheses(S):
    start = 0
    check = 0
    result = ''
    for i in range(len(S)):
        if '(' == S[i]:
            check += 1
        else:
            check -= 1
        if check == 0:
            result += S[start+1:i]
            start = i+1
    return result

def main():
    S = input()
    print(removeOuterParentheses(S))

if __name__ == "__main__":
    main()