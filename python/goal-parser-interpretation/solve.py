def interpret(command):
    return command.replace('(al)','al').replace('()','o')

def main():
    print(interpret(input()))

if __name__ == "__main__":
    main()