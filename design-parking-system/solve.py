def __init__(self,big,medium,small):
    self.slots = {1: big, 2: medium, 3: small}

def addCar(self,carType):
    if self.slots[carType] == 0: return False
    self.slots[carType] -= 1
    return True