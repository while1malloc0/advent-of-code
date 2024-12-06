from typing import Dict, List, Optional, Type


class Rule:
    @staticmethod
    def from_str(s: str):
        lhs, rhs = [int(item.strip()) for item in s.strip().split("|")]
        return Rule(lhs, rhs)

    def __init__(self, lhs: int, rhs: int) -> None:
        self.lhs = lhs
        self.rhs = rhs

    def __eq__(self, other):
        return self.lhs == other.lhs and self.rhs == other.rhs

    def check(self, report: List[int]) -> bool:
        for i, num in enumerate(report):
            if num == self.rhs:
                for j in range(i, len(report)):
                    if report[j] == self.lhs:
                        return False
            if num == self.lhs:
                for j in range(i, 0, -1):
                    if report[j] == self.rhs:
                        return False
        return True


class Rules:
    @staticmethod
    def from_str(s):
        raw = [Rule.from_str(line) for line in s.strip().split("\n")]
        table = dict()
        for rule in raw:
            if not table.get(rule.lhs):
                table[rule.lhs] = []
            if not table.get(rule.rhs):
                table[rule.rhs] = []
            table[rule.lhs].append(rule)
            table[rule.rhs].append(rule)
        return Rules(table)

    def __init__(self, rule_table: Dict[int, List[Rule]]):
        self._table = rule_table

    def check(self, report: List[int]) -> bool:
        for num in report:
            rules = self._applicable_rules(num)
            for rule in rules:
                if rule and not rule.check(report):
                    return False
        return True

    def _applicable_rules(self, n: int) -> Optional[List[Rule]]:
        return self._table.get(n, None)


def p1(input: str) -> int:
    with open(input) as content:
        rules_raw, reports_raw = content.read().strip().split("\n\n")
        rules = Rules.from_str(rules_raw)
        reports = []
        for line in reports_raw.split("\n"):
            if line == "":
                continue
            reports.append([int(c) for c in line.split(",")])
        correct_reports = [r for r in reports if rules.check(r)]
        result = sum([middle_number(row) for row in correct_reports])
        return result


def p2(input: str) -> int:
    with open(input) as content:
        pass


def middle_number(haystack: List[int]) -> int:
    halfway = len(haystack) // 2
    return haystack[halfway]


if __name__ == "__main__":
    part1 = p1("input.txt")
    print("Part1: {}".format(part1))
    part2 = p2("input.txt")
    print("Part2: {}".format(part2))
