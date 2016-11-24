RNA_MAP = {
    'G': 'C',
    'C': 'G',
    'T': 'A',
    'A': 'U'
}

def to_rna(dna):
    return ''.join([RNA_MAP[i] for i in dna])