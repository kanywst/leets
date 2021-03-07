def reformatNumber(number):
    result = []
    number = number.replace('-','')
    number = number.replace(' ','')
    while True:
        if len(number) == 4:
            result.append(number[:2])
            number = number[2:]
        else:
            result.append(number[:3])
            number = number[3:]
        if number == "":
            break
    return '-'.join(result)

def main():
    number = input()
    print(reformatNumber(number))

if __name__ == "__main__":
    main()