def countMatches(items, ruleKey, ruleValue):
    cnt = 0
    if ruleKey == "type":
        indx = 0
    elif ruleKey == "color":
        indx = 1
    elif ruleKey == "name":
        indx = 2
    for i in items:
        if i[indx] == ruleValue:
            cnt+=1
    return cnt

def main():
    items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]]
    ruleKey = "color"
    ruleValue = "silver"
    print(countMatches(items,ruleKey,ruleValue))

if __name__ == "__main__":
    main()