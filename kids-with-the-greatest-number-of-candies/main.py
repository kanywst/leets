def kidsWithCandies(candies, extraCandies):
    result = []
    for i in candies:
        i+=extraCandies
        if max(candies) <= i:
            result.append(True)
        else:
            result.append(False)
    return result

def main():
    print(kidsWithCandies([2,3,5,1,3],3))

if __name__ == "__main__":
    main()