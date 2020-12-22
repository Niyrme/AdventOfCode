def main():
	floor = 0
	basementIdx = -1
	with open("./input.txt", 'r') as f:
		alreadyBasement = False
		i=0
		for char in f.read():
			i+=1
			if str(char) == '(':
				floor+=1
			elif str(char)==')':
				floor-=1
			if floor == -1 and not alreadyBasement:
				alreadyBasement = True
				basementIdx = i

	print(f"The instructions take Santa to floor {floor}")
	print(f"The first time Santa enters the basement is on idx {basementIdx}")

if __name__ == "__main__":
	main()
