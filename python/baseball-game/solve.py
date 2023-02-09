def calPoints(ops):
    result = []
    for i in ops:
        if i == 'C':
            result = result[:-1]
        elif i == 'D':
            result.append(int(result[-1])*2)
        elif i == '+':
            result.append(int(result[-1])+int(result[-2]))
        else:
            result.append(i)
    else:
        return sum(map(int,result))

def main():
    ops = input().split(' ')
    print(calPoints(ops))

if __name__ == "__main__":
    main()