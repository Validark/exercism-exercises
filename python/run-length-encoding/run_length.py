def encode(seq):
    temp = ''
    last_letter = seq[:1]
    count = 0
    for i in seq:
        if i == last_letter:
            count += 1
        else:
            if count == 1:
                count = ''
            temp += str(count) + last_letter
            last_letter = i
            count = 1
    if count == 1:
        count = ''
    temp += str(count) + last_letter
    return temp

def decode(seq):
    temp = ''
    last_letters = ''
    for i in seq:
        if i.isnumeric():
            last_letters += i
        else:
            if last_letters == '':
                last_letters = 1
            temp += ''.join([i for l in range(int(last_letters))])
            last_letters = ''
    return temp