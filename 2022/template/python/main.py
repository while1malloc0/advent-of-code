import argparse
from os import path


def part_one(input):
    return input


def part_two(input):
    return input


parser = argparse.ArgumentParser(prog="aoc")
parser.add_argument("part", type=int)
parser.add_argument("--example", default=False)


def main():
    args = parser.parse_args()
    if args.part not in (1, 2):
        raise TypeError("part argument must be one of: 1 or 2")

    in_file_name = "example.txt" if args.example else "input.txt"
    in_file_path = path.abspath(__file__ + f"/../../data/{in_file_name}")
    input = ""
    with open(in_file_path) as f:
        input = f.read()

    result = ""
    if args.part == 1:
        result = part_one(input)
    elif args.part == 2:
        result = part_two(input)
    else:
        raise Exception("unreachable")

    print(result)


if __name__ == "__main__":
    main()
