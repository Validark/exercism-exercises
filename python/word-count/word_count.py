#Python3
from collections import Counter

import re

def word_count(phrase):
    cnt = Counter()
    for i in re.findall('[^\W_]+', phrase.lower()):
        cnt[i] += 1
    return dict(cnt)