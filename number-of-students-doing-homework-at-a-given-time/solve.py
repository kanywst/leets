def busyStudent(startTime, endTime, queryTime):
    cnt = 0
    for i in range(len(startTime)):
        if startTime[i] <= queryTime <= endTime[i]:
            cnt += 1
    return cnt

def main():
    startTime = list(map(int,input().split(' ')))
    endTime = list(map(int,input().split(' ')))
    queryTime = int(input())
    print(busyStudent(startTime, endTime, queryTime))

if __name__ == "__main__":
    main()