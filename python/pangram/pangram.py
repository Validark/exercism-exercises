def is_pangram(phrase):
    alpha = set(i for i in phrase.lower() if i.isalpha())
    return len(alpha) == 26