digit_text = {'00': 'Zero', '01': 'One', '02': 'Two', '03': 'Three',
'04': 'Four', '05': 'Five', '06': 'Six', '07': 'Seven', '08': 'Eight',
'09': 'Nine', '10': 'Ten', '11': 'Eleven', '12': 'Twelve', '13': 'Thirteen',
'14': 'Fourteen', '15': 'Fifteen', '16': 'Sixteen', '17': 'Seventeen',
'18': 'Eighteen', '19': 'Nineteen', '20': 'Twenty', '30': 'Thirty',
'40': 'Forty', '50': 'Fifty', '60': 'Sixty', '70': 'Seventy', '80': 'Eighty',
'90': 'Ninety'}

cats = ['Billion', 'Million', 'Thousand']

def num2text2(n):
    n = str(n)
    if len(n) < 2: n = '0' + n
    if n[0] == '1' or n[1] == '0':
        return digit_text[n]
    else:
        if int(n) < 10:
            return digit_text['0' + n[1]]
        else:
            return digit_text[n[0]+'0'] + ' ' + digit_text['0'+n[1]]

def num2text3(n):
    assert n < 1000
    if n < 10:
        return digit_text['0' + str(n)]
    elif n < 100:
        return num2text2(n)
    else:
        n = str(n)
        return "%s Hundred%s" % (digit_text['0'+n[0]], (' ' + num2text2(int(n[1:])), '')[n[1:] == '00'])

def num2text(n):
    n = str(n)
    n = n.zfill(12)
    s = []
    for (i, cat) in enumerate(cats):
        start, end = i * 3, (i + 1) * 3
        if int(n[start:end]) > 0:
            s.append(num2text3(int(n[start:end])) + ' ' + cat)
    if int(n[-3:]) > 0:
        s.append(num2text3(int(n[-3:])))
    if s:
        return ', '.join(s)

import sys, random

if len(sys.argv) < 2:
    n = random.randrange(999999999999)
else:
    n = int(sys.argv[1])

print(n)

if n > 999999999999:
    print('Number must be less than 999999999999')
    sys.exit()

print(num2text(n))
