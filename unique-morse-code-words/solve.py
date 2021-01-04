def uniqueMorseRepresentations(words):
    morse_list = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
    morse_dic = {}
    result = []
    for i in range(len(morse_list)):
        morse_dic[chr(ord('a')+i)] = morse_list[i]
    cnt = 0
    for i in words:
        tmp = ''
        for j in i:
            tmp += morse_dic[j]
        result.append(tmp)
    return len(set(result))

def main():
    words = input().split(' ')
    print(uniqueMorseRepresentations(words))

if __name__ == "__main__":
    main()