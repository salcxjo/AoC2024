array = []
left = []
right = []
while True:
    inp = input()
    if inp == '': break
    else:
        inp = inp.replace('   ','n')
        list1 = inp.split('n')
        list1 = ' '.join(list1)
        array += list1.split()
for i in range(0,len(array),2):
        left.append(array[i])
for i in range(1,len(array),2):
        right.append(array[i])
right.sort()
left.sort()
running = 0
for i in range(len(left)):
    running += abs(((int(left[i])) - (int(right[i]))))
print(running)