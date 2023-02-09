def countGoodRectangles(rectangles):
    result = []
    for i in rectangles:
        result.append(min(i))
    return result.count(max(result))

def main():
    rectangles = [[5,8],[3,9],[5,12],[16,5]]
    print(countGoodRectangles(rectangles))

if __name__ == "__main__":
    main()