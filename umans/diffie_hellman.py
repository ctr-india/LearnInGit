# use Python 3 print function
# this allows this code to run on python 2.x and 3.x
from __future__ import print_function
 
# Variables Used
sharedPrime = 23    # p
sharedBase = 5      # g
 
aliceSecret = 6     # a
bobSecret = 15      # b
 
# Begin
print( &amp;amp;amp;amp;amp;amp;quot;Publicly Shared Variables:&amp;amp;amp;amp;amp;amp;quot;)
print( &amp;amp;amp;amp;amp;amp;quot;   Publicly Shared Prime: &amp;amp;amp;amp;amp;amp;quot; , sharedPrime )
print( &amp;amp;amp;amp;amp;amp;quot;   Publicly Shared Base:  &amp;amp;amp;amp;amp;amp;quot;, sharedBase )
 
# Alice Sends Bob A = g^a mod p
A = (sharedBase**aliceSecret) % sharedPrime
print( &amp;amp;amp;amp;amp;amp;quot;\n Alice Sends Over Public Chanel: &amp;amp;amp;amp;amp;amp;quot; , A )
 
# Bob Sends Alice B = g^b mod p
B = (sharedBase ** bobSecret) % sharedPrime
print( &amp;amp;amp;amp;amp;amp;quot;   Bob Sends Over Public Chanel: &amp;amp;amp;amp;amp;amp;quot;, B )
 
print( &amp;amp;amp;amp;amp;amp;quot;\n------------\n&amp;amp;amp;amp;amp;amp;quot; )
print( &amp;amp;amp;amp;amp;amp;quot;Privately Calculated Shared Secret:&amp;amp;amp;amp;amp;amp;quot; )
# Alice Computes Shared Secret: s = B^a mod p
aliceSharedSecret = (B ** aliceSecret) % sharedPrime
print( &amp;amp;amp;amp;amp;amp;quot;   Alice Shared Secret: &amp;amp;amp;amp;amp;amp;quot;, aliceSharedSecret )
 
# Bob Computes Shared Secret: s = A^b mod p
bobSharedSecret = (A**bobSecret) % sharedPrime
print( &amp;amp;amp;amp;amp;amp;quot;   Bob Shared Secret: &amp;amp;amp;amp;amp;amp;quot;, bobSharedSecret )

