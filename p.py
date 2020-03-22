from sys import argv
from typing import List

if len(argv) == 1:
    print("Enter something. E.g: 23 34 45 56")
else:
    inp: List[str] = argv[1:]
    output: str = ", ".join(inp)
    print(output)
