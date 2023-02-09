def countBalls(lowLimit, highLimit):
    ball_counts = {}
    for i in range(lowLimit,highLimit+1):
        tmp = sum(list(map(int,list(str(i)))))
        if tmp in ball_counts:
            ball_counts[tmp] += 1
        else:
            ball_counts[tmp] = 1
    return max(ball_counts.values())
    

def main():
    lowLimit = int(input())
    highLimit = int(input())
    print(countBalls(lowLimit,highLimit))

if __name__ == "__main__":
    main()