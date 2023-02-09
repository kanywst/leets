from collections import deque

def diStringMatch(S):
    d = deque(list(range(len(S)+1)))
    result = []
    for i in S:
        if i == 'I':
            result.append(d.popleft())
        elif i == 'D':
            result.append(d.pop())
    else:
        result.append(d.pop())
    return result

def main():
    S = input()
    print(diStringMatch(S))

if __name__ == "__main__":
    main()