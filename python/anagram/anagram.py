def sort(word):
    return ''.join(sorted(word.lower()))

def detect_anagrams(word, ps):
    return [i for i in ps if sort(word) == sort(i) and i.lower() != word.lower()]