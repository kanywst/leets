def arrayStringsAreEqual(word1,word2):
    if ''.join(word1) == ''.join(word2):
        return True
    else:
        return False

def main():
    word1 = input().split(' ')
    word2 = input().split(' ')
    print(arrayStringsAreEqual(word1,word2))

if __name__ == "__main__":
    main()