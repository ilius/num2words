#!/usr/bin/python
# -*- coding: utf-8 -*-
##   File: num2fa-0.1.1.py
##
##   Author: Saeed Rasooli <saeed.gnu@gmail.com>  (ilius)
##
##   This library is free software; you can redistribute it and/or
##   modify it under the terms of the GNU Lesser General Public
##   License as published by the Free Software Foundation; either
##   version 2.1 of the License, or (at your option) any later version.
##
##   This library is distributed in the hope that it will be useful,
##   but WITHOUT ANY WARRANTY; without even the implied warranty of
##   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
##   Lesser General Public License for more details.

import sys

faBaseNum={
1:'یک',
2:'دو',
3:'سه',
4:'چهار',
5:'پنج',
6:'شش',
7:'هفت',
8:'هشت',
9:'نه',
10:'ده',
11:'یازده',
12:'دوازده',
13:'سیزده',
14:'چهارده',
15:'پانزده',
16:'شانزده',
17:'هفده',
18:'هجده',
19:'نوزده',
20:'بیست',
30:'سی',
40:'چهل',
50:'پنجاه',
60:'شصت',
70:'هفتاد',
80:'هشتاد',
90:'نود',
100:'صد',
200:'دویست',
300:'سیصد',
500:'پانصد',
1000:'هزار',
10**6:'میلیون',
10**9:'میلیارد'
}
faBaseNumKeys = faBaseNum.keys()

def split3(n):
  parts=[]
  while n>999:
    parts.append(n%1000)
    n = int(n/1000)
  parts.append(n)
  return parts

def num2fa(n):
  if n in faBaseNumKeys:
    return faBaseNum[n]
  if n>999:
    parts = split3(n)
    fa=''
    k=len(parts)
    for i in range(k):
      p = parts[i]
      if i==0:
        fa += num2fa(p)
        continue
      order = 10**(3*i)
      if order in faBaseNumKeys:
        faOrder = faBaseNum[order]
      else:
        (d,m) = divmod(i,3)
        ls = [faBaseNum[10**9]] * d
        if m!=0:
          ls = [faBaseNum[10**(3*m)]] + ls
        faOrder = '‌'.join(ls)
      if i==1 and p==1:
        fa = faOrder + ' و ' + fa
      else:
        fa = num2fa(p) + ' ' + faOrder + ' و ' + fa
    return fa
  ## now assume that n <= 999
  y = n%10
  d = int((n%100)/10)
  s = int(n/100)
  dy = 10*d+y
  fa=''
  if s!=0:
    if s*100 in faBaseNumKeys:
      fa += faBaseNum[s*100]
    else:
      fa += (faBaseNum[s]+faBaseNum[100])
    if d!=0 or y!=0:
      fa += ' و '
  if d!=0:
    if dy in faBaseNumKeys:
      fa += faBaseNum[dy]
      return fa
    fa += faBaseNum[d*10]
    if y!=0:
      fa += ' و '
  if y!=0:
    fa += faBaseNum[y]
  return fa






if __name__=='__main__':
  for arg in sys.argv[1:]:
    try:
      i = int(arg)
    except:
      pass
    else:
      print '%s\t%s'%(i,num2fa(i))
  ############################
  """
  f=file('fa_num.txt', 'w')
  for i in range(10000):
    f.write('%s\t%s\n'%(i,num2fa(i)))
  f.close()
  """ 
