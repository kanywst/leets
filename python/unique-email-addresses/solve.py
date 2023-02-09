def numUniqueEmails(emails):
    result = []
    for email in emails:
        email_list = email.split('@')
        email_list[0] = email_list[0].replace('.','')
        hostname = list(email_list[0])
        if '+' in hostname:
            plus_index = hostname.index('+')
        else:
            plus_index = len(hostname)
        hostname = hostname[:plus_index]
        result.append('{}@{}'.format(''.join(hostname),email_list[1]))
    return len(list(set(result)))

def main():
    emails = ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
    print(numUniqueEmails(emails))

if __name__ == "__main__":
    main()