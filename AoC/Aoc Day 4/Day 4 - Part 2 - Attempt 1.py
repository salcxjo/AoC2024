cnt = 0
line = 0
array=[]
while True:
    inp = input()
    if inp == '': break
    else:
       array.insert(line, list(inp))
       
for i in range(len(array)-2):
    for j in range(len(array[i])-2):
            if array[i][j] == 'S':
                if array[i+1][j+1] == 'A':
                    if array[i+2][j+2] == "M":
                        if array[i][j+2] == 'S':
                            if array[i+2][j] == 'M':
                                cnt += 1
            if array[i][j] == 'M':
                if array[i+1][j+1] == 'A':
                    if array[i+2][j+2] == "S":
                        if array[i][j+2] == 'S':
                            if array[i+2][j] == 'M':
                                cnt += 1
            if array[i][j] == 'S':
                if array[i+1][j+1] == 'A':
                    if array[i+2][j+2] == "M":
                        if array[i][j+2] == 'M':
                            if array[i+2][j] == 'S':
                                cnt += 1
            if array[i][j] == 'M':
                if array[i+1][j+1] == 'A':
                  if array[i+2][j+2] == "S":
                     if array[i][j+2] == 'M':
                          if array[i+2][j] == 'S':
                                cnt += 1
print(cnt)