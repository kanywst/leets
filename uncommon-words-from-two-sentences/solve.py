from collections import Counter

def uncommonFromSentences(A,B):
    array = A.split(' ') + B.split(' ')
    cnt = list(Counter(array).items())
    result = [item[0] for item in cnt if item[1] == 1]

    return result
    
def main():
    A = input()
    B = input()
    print(uncommonFromSentences(A,B))

if __name__ == "__main__":
    main()