import hashlib

if __name__ == "__main__":
	i = 0
	input = "bgvyzdsv"
	out5 = out6 = 0

	while True:
		if out5 != 0 and out6 != 0:
			break
		hash = hashlib.md5(f"{input}{i}".encode())
		hashOut = hash.hexdigest()
		if hashOut[:5] == "00000" and out5 == 0:
			out5 = i
		if hashOut[:6] == "000000" and out6 == 0:
			out6 = i
		i += 1

	print(f"hash with 5 '0's: {out5}")
	print(f"hash with 6 '0's: {out6}")
