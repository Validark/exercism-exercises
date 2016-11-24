def distance(seq1, seq2):
    return len([i for i, _ in enumerate(seq1) if seq1[i] != seq2[i]])