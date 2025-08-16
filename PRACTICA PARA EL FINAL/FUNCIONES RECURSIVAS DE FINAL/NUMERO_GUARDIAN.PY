def casi(a, b, resA, resB, sumaA, sumaB):
    # Caso base: cuando a y b llegan a 0
    if a == 0 and b == 0:
        return sumaA == resB and sumaB == resA
    
    # Procesar a (divisores impares de resA)
    if a > 0:
        if (resA % a == 0) and (a % 2 != 0):
            return casi(a - 1, b, resA, resB, sumaA + a, sumaB)
        else:
            return casi(a - 1, b, resA, resB, sumaA, sumaB)
    
    # Procesar b (divisores pares de resB)
    if b > 0:
        if (resB % b == 0) and (b % 2 == 0):
            return casi(a, b - 1, resA, resB, sumaA, sumaB + b)
        else:
            return casi(a, b - 1, resA, resB, sumaA, sumaB)

# Ejemplo de uso
a = int(input("Ingrese valor de a: "))
b = int(input("Ingrese valor de b: "))

if casi(a - 1, b - 1, a, b, 0, 0):
    print("Cumple la condición")
else:
    print("No cumple la condición")

"""
funcion casi(a,b,resA,resB,sumaA,sumaB:entero)->boleano
	si (a=0) y (b=0) entonces 
		si (sumaA=resB) y (sumaB=resA) entonces
			casi:verdadero
		sino 
			casi:falso 
		finsi 
	sino 
		si a>0 entonces 
			si ((resA mod a)=0) y ((a mod 2)=0) entonces
				sumaA:=sumaA + a
				casi:(a-1,b,resA,resB,sumaA,sumaB)
			sino 
				casi:((a-1,b,resA,resB,sumaA,sumaB))
			finsi 
		finsi 
		si b>0 entonces 
			si ((resB mod b)=0) y ((b mod 2)<>0) entonces
				sumaB:=sumaB + b
				casi:(a,b-1,resA,resB,sumaA,sumaB)
			sino 
				casi:((a,b-1,resA,resB,sumaA,sumaB))
			finsi 
		finsi
	finsi 
Método matemático
1. Un número d es n si n mod d= O
2. Es par si d mod 2 = O.
3. Es impar si d mod 2 <>0
"""






