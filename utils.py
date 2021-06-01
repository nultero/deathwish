
def unicodes(match: str) -> str:
    tmp = ""
    cases = {
        '9': "\u277e",
        'finger': "\u261E",
    }
    if match in cases.keys():
        tmp += cases[match]

    return tmp