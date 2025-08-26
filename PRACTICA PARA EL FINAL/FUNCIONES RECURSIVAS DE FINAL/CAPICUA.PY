def es_capicua(num1,num2,num_og:int)-> bool: 
    if (num1 == 0):
        if (num2 ==num_og):
            return True
        else:
            return False
    else:
        num2= num2 * 10 + num1 % 10
        num1 = num1 // 10 
        return es_capicua(num1, num2, num_og)

print("Ingrese un numero para ver si es capicua o no: ")
numero=int(input())
 
if es_capicua(numero, 0, numero):
    print("El numero es capicua")
else:
    print("El numero no es capicua")
"""
funcion es_capicua(num1,num2,num_og:entero)->boleano
    si num=0 entonces
        si num2=num_og entonces
            es_capicua= verdadero
        sino
            es_capicua= falso
        finsi 
    sino
        num2= num2 * 10 + num1 mod 10
        num1 = num1 div 10 
        es_capicua= es_capicua(num1, num2, num_og)
    finsi
finfuncion
"""

        

   