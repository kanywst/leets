def defangIPaddr(address):
    return '[.]'.join(address.split('.'))

def main():
    print(defangIPaddr("1.1.1.1"))

if __name__ == "__main__":
    main()