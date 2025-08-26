def casi(num1,suma,num2:int)-> bool: 
    if (num2==0):
        if (suma==num1-1):
            return True
        else:
            return False
    else:
        if ((num1 % num2)==0):
           
            return casi(num1,suma + num2,num2-1)
        else:
            return casi(num1,suma,num2-1)

    
    

print("Ingrese un numero para ver si es casi perfecto o no: ")
numero=int(input())
if casi(numero, 0, numero-1):
    print("El numero es casi perfecto.")
else:
    print("El numero no es casi perfecto.")

"""
funcion casi(num1,suma,num2:entero)->boleano 
    si (num2=0) entonces 
        si (suma=num1-1) entonces 
            casi:=verdadero
        sino 
            casi:=falso
        finsi
    sino 
        si ((num1 mod num2)=0) entonces 
            casi(num1,suma+num2,num2-1)
        sino 
            casi(num1,suma,num2-1)
        finsi 
    finsi 
finfuncion
"""
    
    