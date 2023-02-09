def mergeAlternately(word1,word2):
    result = ''
    if len(word1) < len(word2):
        l = len(word1)
        word = word2
    else:
        l = len(word2)
        word = word1
    for i in range(l):
        result += word1[i]
        result += word2[i]
    result += word[l:]
    return result
        
def main():
    word1 = input()
    word2 = input()
    print(mergeAlternately(word1,word2))

if __name__ == "__main__":
    main()