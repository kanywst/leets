import re

def minOperations(logs):
    cnt = 0
    for i in logs:
        pattern_current = r'^\./$'
        pattern_back = r'^\.\./$'
        if re.search(pattern_current, i):
            continue
        elif re.search(pattern_back, i):
            if cnt <= 0:
                cnt = 0
            else:
                cnt -= 1
        else:
            cnt += 1
    return cnt

def main():
    logs = ["d1/","d2/","./","d3/","../","d31/"]
    print(minOperations(logs))

if __name__ == "__main__":
    main()