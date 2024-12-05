cnt = 0
line = 0
array=[]
while True:
    inp = input()
    if inp == '': break
    else:
       cnt += inp.count('XMAS')
       cnt += inp.count('SAMX') 
       array.insert(line, list(inp))
       
for i in range(len(array)-3):
    for j in range(len(array[i])-3):
        if array[i][j] == 'X':
            if array[i+1][j+1] == 'M':
                if array[i+2][j+2] == "A":
                    if array[i+3][j+3] == 'S':
                        cnt += 1
    for j in range(3, len(array[i])):
        if array[i][j] == 'X':
            if array[i+1][j-1] == 'M':
                if array[i+2][j-2] == "A":
                    if array[i+3][j-3] == 'S':
                        cnt += 1
    for j in range(len(array[i])):
        if array[i][j] == 'X':
            if array[i+1][j] == 'M':
                if array[i+2][j] == "A":
                    if array[i+3][j] == 'S':
                        cnt += 1

for i in range(3,len(array)):
    for j in range(3,len(array[i])):
        if array[i][j] == 'X':
            if array[i-1][j-1] == 'M':
                if array[i-2][j-2] == "A":
                    if array[i-3][j-3] == 'S':
                        cnt += 1

    for j in range(len(array)-3):
        if array[i][j] == 'X':
            if array[i-1][j+1] == 'M':
                if array[i-2][j+2] == "A":
                    if array[i-3][j+3] == 'S':
                        cnt += 1
    for j in range(len(array[i])):
        if array[i][j] == 'X':
            if array[i-1][j] == 'M':
                if array[i-2][j] == "A":
                    if array[i-3][j] == 'S':
                        cnt += 1


print(cnt)