total1 = 0
total2 = 0

File.read_lines("./input.txt").each do |val|
	total1 += (val.to_i // 3) - 2
	total2 += calcFuel(val.to_i)
end

puts "part 1: #{total1}"
puts "part 2: #{total2}"

def calcFuel(fuel)
	c = (fuel.to_i // 3) - 2
	if c <= 0
		return 0
	end
	return c + calcFuel(c)
end
