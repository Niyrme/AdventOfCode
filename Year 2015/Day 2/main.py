def calculatePaper(dimensions):
	side1 = int(dimensions[0]) * int(dimensions[1])
	side2 = int(dimensions[1]) * int(dimensions[2])
	side3 = int(dimensions[0]) * int(dimensions[2])
	smallest = min(side1, side2, side3)
	return (2 * side1) + (2 * side2) + (2 * side3) + smallest

def calculateRibbon(dimensions):
	d = [int(dimensions[0]), int(dimensions[1]), int(dimensions[2])]
	sm1 = int(sorted(d)[0]) # Smallest number
	sm2 = int(sorted(d)[1]) # Second smallest number
	return (2 * sm1) + (2 * sm2) + (d[0] * d[1] * d[2])

def main():
	with open("./input.txt", 'r') as f:
		totalPaper = 0
		totalRibbon = 0
		for line in f.readlines():
			dimensions = line.split("x")
			totalPaper += calculatePaper(dimensions)
			totalRibbon += calculateRibbon(dimensions)

		print(f"Total required amount of wrapping paper: {totalPaper}")
		print(f"Total required amount of ribbon: {totalRibbon}")

if __name__ == "__main__":
	main()
