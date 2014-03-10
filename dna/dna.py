def count_nucleotides(nucs:str) -> [int, int, int, int]:
    return [nucs.count(it) for it in ['A', 'C', 'G', 'T']]


