def destCity(paths):
    path_dic = {}
    tmp = paths[0][0]
    for i in paths:
        path_dic[i[0]] = i[1]
    while True:
        if tmp in path_dic:
            tmp = path_dic[tmp]
        else:
            break
    return tmp

def main():
    paths = [["B","C"],["D","B"],["C","A"]]
    print(destCity(paths))

if __name__ == "__main__":
    main()