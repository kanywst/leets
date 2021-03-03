def countCharacters(words, chars):
    cnt = 0
    for c in words:
        for i in set(list(c)):
            if c.count(i) > chars.count(i):
                break
        else:
            cnt += len(c)
    return cnt

def main():
    words = ["cat","bt","hat","tree"]
    chars = "atach"
    print(countCharacters(words,chars))

if __name__ == "__main__":
    main()