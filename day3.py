class Claim:
    id = ""
    coord = [0, 0]
    size = [0, 0]

    def __init__(self, claim_id, claim_coord, claim_size):
        self.id = claim_id
        self.coord = coord
        self.size = size

# Create two dimensional array (1000 x 1000) with all zeros
max_size = 1000
fabric = [0] * max_size
for i in range(0, max_size):
    fabric[i] = [0] * max_size

# Parse input file
claims = []
with open("input_day3.txt") as f:
    for line in f:
        claim_nbr = line[1: line.find(" @")]
        x = line[line.find("@")+1: line.find(",")]
        y = line[line.find(",")+1: line.find(":")]
        width = line[line.find(":")+2: line.find("x")]
        height = line[line.find("x")+1:]
        print("Claim: {}, x: {}, y: {}, w: {}, h: {}".format(claim_nbr, x, y, width, height))
        coord = [int(x), int(y)]
        size = [int(width), int(height)]
        claims.append(Claim(claim_nbr, coord, size))

# Check overlap
for c in claims:
    x = c.coord[0]
    y = c.coord[1]
    w = c.size[0]
    h = c.size[1]
    print("Claim: {}, x: {}, y: {}, w: {}, h: {}".format(c.id, x, y, w, h))
    for i in range(0, w):
        for j in range(0, h):
            fabric[x+i][y+j] += 1

# Count overlapping inches
overlap = 0
for i in range(0, max_size):
    for j in range(0, max_size):
        if fabric[i][j] > 1:
            overlap += 1

print("Overlapping: {}".format(overlap))

