values = []
with open("input_day2.txt") as f:
    for line in f:
        values.append(line)

# Part 1
two_count = 0
three_count = 0

for s in values:
    local_two_count = 0
    local_three_count = 0
    for letter in s:
        if s.count(letter) is 2 and local_two_count is 0:
            local_two_count = 1
        if s.count(letter) is 3 and local_three_count is 0:
            local_three_count = 1

    two_count += local_two_count
    three_count += local_three_count

print("{} * {} = {}".format(two_count, three_count, two_count*three_count))

# Part 2
nbr_of_id = len(values)
nbr_of_letters = len(values[0])
correct_boxes = []

for s1 in range(0, nbr_of_id):
    for s2 in range (s1+1, nbr_of_id):
        diff_count = 0
        for i in range(0, nbr_of_letters):
            if values[s1][i] is not values[s2][i]:
                diff_count += 1
                if diff_count > 1:
                    break
        if diff_count is 1:
            correct_boxes.append(values.index(values[s1]))
            correct_boxes.append(values.index(values[s2]))
            break
        else:
            diff_count = 0
    if diff_count is 1:
        break

result = ""
for i in range(0, nbr_of_letters):
    if values[correct_boxes[0]][i] is values[correct_boxes[1]][i]:
        result += values[correct_boxes[0]][i]

print("ID: {} word count: {} of {}".format(result, len(result), nbr_of_letters))
