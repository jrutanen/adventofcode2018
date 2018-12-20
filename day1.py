
#Part 1
from typing import List, Any, Union

f_offset = 0

with open("input_day1.txt") as f:
    for line in f:
        f_offset += int(line)

print("part 1 result: {}".format(f_offset))

#Part 2
f_offset = 0
f_changes = []
offset_values = [f_offset]


with open("input_day1.txt") as f:
    for line in f:
        f_changes.append(int(line))

counter = 0

while True:
    for value in f_changes:
        f_offset += value
#        print("Offset: {}".format(f_offset))
        if f_offset in offset_values:
            print("part 2 result: {}".format(f_offset))
            exit()
        else:
            offset_values.append(f_offset)

    counter += 1
    print("Counter: {}".format(counter))
    if counter > 1000:
        print("While loop ran {} times".format(counter))
        exit()
