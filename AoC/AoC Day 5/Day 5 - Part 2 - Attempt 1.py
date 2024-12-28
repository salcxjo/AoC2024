#Take input for Rules
rules = {}
while True:
    inp = input()
    if inp == '': break
    else:
        inp = inp.split('|')
        nums = [ int(x) for x in inp ]
        if nums[1] in rules:
             rules[nums[1]].append(nums[0])
        else:
             rules[nums[1]] = [nums[0]]

#Take input of reports
keyRules = list(rules.keys())
rightMiddle = []
while True:
    inp = input()
    if inp == '': break
    else:
        wrong = 0
        inp = inp.split(',')
        nums = [int(x) for x in inp]
        for i in range(len(nums)):
            if keyRules.count(nums[i]) != 0:
                keyRuleTemp = keyRules.index(nums[i])
                for j in range(i+1, len(nums)):
                    x = keyRules[keyRuleTemp]
                    thing = rules.get(x)
                    if thing.count(nums[j]) != 0:
                        wrong= 1
                        keyruleTemp2 = thing.index(nums[j])
                        y = thing[keyruleTemp2]
                        temp = nums[j]
                        nums[j] = y
                        nums[y] = temp
                        
                        
                        # and then ** we are only adding the middles of the newly corrected reports (middles of the new sorted arrays)
                        # so maybe we add all and subtract what we found in part 1?
                        continue

        else: rightMiddle.append(int(inp[int((len(inp)-1)/2)]))

#Find middle of correct reports and return the sum of all middles
print(sum(rightMiddle))