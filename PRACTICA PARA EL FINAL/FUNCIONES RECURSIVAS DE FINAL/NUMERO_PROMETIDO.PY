def casi(A,B,resA,resB,sumaA,sumaB:int)-> bool:
    if (A==0) and (B==0):
        if (sumaA==resB+1) and (sumaB==resA+1):
            return True
        else:
            return False
    else:
        if (A > 0):
            if ((resA % A)==0):
                return casi(A-1,B,resA,resB,sumaA + A,sumaB)
            else:
                return casi(A-1,B,resA,resB,sumaA,sumaB)
        if (B > 0):
            if ((resB % B)==0):
                return casi(A,B-1,resA,resB,sumaA,sumaB + B)
            else:
                return casi(A,B-1,resA,resB,sumaA,sumaB)
            

print("ingrese un numero a")
A=int(input())
print("ingrese un numero b")
B=int(input())

if casi(A-1,B-1,A,B,0,0):
    print("es numero perfecto")
else:
    print(" no es numero perfecto")

"""
funcion casi(A,B,resA,resB,sumaA,sumaB:enteor)->boleano
	si (A=0) y (B=0) entoces 
		si (sumaA=resB+1) y (sumaB=resA+1) entoces
			casi:verdadero
		sino 
			casi:falso 
	sino 
		si A>0 entoces
			si ((resA mod A)=0) entoces
				casi:(A-1,B,resA,resB,sumaA + A,sumaB)
			sino 
				casi:(A-1,B,resA,resB,sumaA,sumaB)
			finsi 
		finsi
		si B>0 entoces
			si ((resB mod B)=0) entoces
				casi:(A,B-1,resA,resB,sumaA,sumaB + B)
			sino 
				casi:(A,B-1,resA,resB,sumaA,sumaB)
			finsi 
		finsi
	finsi 
finfuncion
"""