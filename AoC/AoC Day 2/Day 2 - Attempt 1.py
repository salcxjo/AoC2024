safeCount = 0
reportCount = 0
unSafeCount = 0
while True:
    inp = input()
    if inp == '': break
    else:
        comparison = unSafeCount
        reportCount += 1
        inp = inp.split()
        order = int(inp[0]) - int(inp[1])
        if order > 0: order = 1
        elif order < 0: order = 0
        else: 
            unSafeCount +=1
        for i in range(len(inp)-1):
            if abs(int(inp[i])-int(inp[i+1])) > 3 or abs(int(inp[i])-int(inp[i+1])) == 0:
               unSafeCount += 1
               break
            if (int(inp[i])-int(inp[i+1]) > 0 and order == 0) or (int(inp[i])-int(inp[i+1]) < 0 and order == 1):
                unSafeCount += 1
                break
        if comparison == unSafeCount: safeCount += 1
print(str(safeCount))