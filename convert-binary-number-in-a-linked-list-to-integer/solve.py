class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def getDecimalValue(self, head: ListNode) -> int:
    ans = []
    while head:
        ans.append(str(head.val))
        head = head.next
    return int(''.join(list(ans)),2)