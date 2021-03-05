def toGoatLatin(S):
    s = S.split(' ')
    result = []
    for i in range(len(s)):
        tmp = ''
        c = s[i]
        if c[0] == 'a' or c[0] == 'i' or c[0] == 'u' or c[0] == 'e' or c[0] == 'o' or c[0] == 'A' or c[0] == 'I' or c[0] == 'U' or c[0] == 'E' or c[0] == 'O':
            tmp += c + 'ma'
        else:
            tmp += c[1:] + c[0] + 'ma'
        tmp += 'a' * (i+1)
        result.append(tmp)

    return ' '.join(result)

def main():
    S = input()
    print(toGoatLatin(S))

if __name__ == "__main__":
    main()