def judgeCircle(moves):
    x,y = 0,0
    result = False
    for i in moves:
        if i == 'U':
            y+=1
        elif i == 'D':
            y-=1
        elif i == 'R':
            x+=1
        elif i == 'L':
            x-=1
    if x==0 and y==0:
        return True
    else:
        return False
        

def main():
    moves = input()
    print(judgeCircle(moves))

if __name__ == "__main__":
    main()