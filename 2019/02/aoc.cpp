#include "aoc.hpp"

// Input hard coded because make somehow breaks reading from file
const std::vector<int> input = {
	1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 6, 19, 23, 2, 6, 23, 27, 1, 5, 27, 31, 2, 31, 9, 35, 1, 35, 5, 39, 1, 39, 5, 43, 1, 43, 10, 47, 2, 6, 47, 51, 1, 51, 5, 55, 2, 55, 6, 59, 1, 5, 59, 63, 2, 63, 6, 67, 1, 5, 67, 71, 1, 71, 6, 75, 2, 75, 10, 79, 1, 79, 5, 83, 2, 83, 6, 87, 1, 87, 5, 91, 2, 9, 91, 95, 1, 95, 6, 99, 2, 9, 99, 103, 2, 9, 103, 107, 1, 5, 107, 111, 1, 111, 5, 115, 1, 115, 13, 119, 1, 13, 119, 123, 2, 6, 123, 127, 1, 5, 127, 131, 1, 9, 131, 135, 1, 135, 9, 139, 2, 139, 6, 143, 1, 143, 5, 147, 2, 147, 6, 151, 1, 5, 151, 155, 2, 6, 155, 159, 1, 159, 2, 163, 1, 9, 163, 0, 99, 2, 0, 14, 0
};
const int inpSize = input.size();

struct value {
	int pos;
	int val;
}in1, in2, out;

int main() {
	// Part 1
	std::vector<int> in = input;
	in.at(1) = 12;
	in.at(2) = 2;
	std::cout << "[Part 1] Value at position 0 is: " << loopInput(in) << std::endl;

	// Part 2
	int target = 19690720;
	for (int noun = 0; noun < 100; noun++) {
		for (int verb = 0; verb < 100; verb++) {
			in.at(1) = noun;
			in.at(2) = verb;
			if (loopInput(in) == target) {
				std::cout << "[Part 2] Inputs resulting in " << target << " are:" << std::endl <<
					"\t\tnoun = (" << noun << ")" << std::endl <<
					"\t\tverb = (" << verb << ")" << std::endl;
				std::cout << "[Part 2] 100 * noun + verb = " << 100 * noun + verb << std::endl;
				return 0;
			}
		}
	}

	return 0;
}

int loopInput(std::vector<int> i) {
	int pos = 0;

	while (true) {
		/* opcodes
			1  = add
			2  = multiply
			99 = exit
			_  = error
		*/
		int code = i.at(pos);

		in1.pos = i.at(pos + 1);
		in1.val = i.at(in1.pos);

		in2.pos = i.at(pos + 2);
		in2.val = i.at(in2.pos);

		out.pos = i.at(pos + 3);

		switch (code) {
			case 1:
				out.val = in1.val + in2.val;
				pos += 4;
				break;
			case 2:
				out.val = in1.val * in2.val;
				pos += 4;
				break;
			case 99:
				return i.at(0);

			default:
				std::cout << "unsupported opcode at position [" << pos << "]: " << i.at(pos) << "!" << std::endl;
				out.val = -1;
				break;
		}

		if (out.val == -1)
			break;
		else
			i.at(out.pos) = out.val;
	}

	return -1;
}
