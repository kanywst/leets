def freqAlphabets(s):
    result = ''
    i = 0
    while i < len(s):
        if i+2 < len(s) and s[i+2] == '#':
            result += chr(int(s[i:i+2])+96)
            i+=3
        else:
            result += chr(int(s[i])+96)
            i+=1
    return result

def main():
    s = input()
    print(freqAlphabets(s))

if __name__ == "__main__":
    main()