def finalPrices(prices):
    for i in range(len(prices)):
        for j in range(i+1,len(prices)):
            if prices[i] >= prices[j]:
                prices[i] -= prices[j]
                break
    return prices

def main():
    prices = list(map(int,input().split(' ')))
    print(finalPrices(prices))

if __name__ == "__main__":
    main()