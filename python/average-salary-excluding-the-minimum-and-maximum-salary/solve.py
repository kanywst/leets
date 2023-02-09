def average(salary):
    salary.sort()
    tmp = salary[1:-1]
    return sum(tmp)/len(tmp)

def main():
    salary = list(map(int,input().split(' ')))
    print(average(salary))

if __name__ == "__main__":
    main()