def minTimeToVisitAllPoints(points):
    result = 0
    for i in range(len(points)-1):
        x_tmp = abs(points[i][0] - points[i+1][0])
        y_tmp = abs(points[i][1] - points[i+1][1])
        result += max(x_tmp,y_tmp)
    return result

def main():
    points = [[1,1],[3,4],[-1,0]]
    print(minTimeToVisitAllPoints(points))

if __name__ == "__main__":
    main()