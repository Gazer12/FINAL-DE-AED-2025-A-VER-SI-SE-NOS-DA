ACCION ES 
	AMBIENTE 
		SEC,SALIDA:secuencia de caracter 
		v:caracter
		cant_pal,car_inv,total_car:entero 
		digitos = {'0','1','2','3','4','5','6','7','8','9'} 
		vocal = {'a','e','i','o','u'}
		consonante =  {'b','c','d','f','g','h','j','k','l','m','n','p','q','r','s','t','v','w','x','y','z'}

		funcion numerito(x:entero)->entero 
			si ((x mod 2)<>0) entonces 
				numerito:(x*x*x) 
			sino 
				numerito:(x*2)

			finsi 
		finfuncion 
		funcion entero(x:caracter)->entero 
			segun x hacer 
				="0"	:	entero:=0 
				="1"	:	entero:=1
				="2"	:	entero:=2 
				="3"	:	entero:=3 
				="4"	:	entero:=4 
				="5"	:	entero:=5 
				="6"	:	entero:=6 
				="7"	:	entero:=7 
				="8"	:	entero:=8
				="9"	:	entero:=9 
			finsegun 
		finfuncion
		PROCEDIMIENTO vocales(x:caracter) es 
			segun x hacer 
				="a"	:	grabar(salida,"u")
							car_vali:=car_vali + 1
				="e"	:	grabar(salida,"a")
							car_vali:=car_vali + 1
				="i"	:	grabar(salida,"e")
							car_vali:=car_vali + 1
				="o"	:	grabar(salida,"i")
							car_vali:=car_vali + 1
				="u"	:	grabar(salida,"o")
							car_vali:=car_vali + 1
			finsegun 
		FINPROCEDIMIENTO informar() es 
			esc("cantidad de palabras del mensaje", cant_pal)
			esc("porcentaje de caracter invalidos",((car_inv/total_car)*100))
			si ((car_inv/total_car)*100) > 10 entonces 
				esc("mensaje conrrupto")
			sino 
				si (car_inv=car_vali) entoces  
					esc("mensaje alerta")
				sino 
					si (car_inv<(car_vali/2)) entoces 
						esc("mesnaje curado")
					finsi 
				finsi 
			finsi 
		FINPROCEDIMIENTO
		
		

	
	
		proceso 
		arr(sec)
		crear(SALIDA)
		avz(sec,v)
		cant_pal:=1
		car_inv:=0
		total_car:=0
		car_vali:=0
		numero:=0

		MIENTRAS ndfs(sec) hacer 
			si v=" " entonces 
				cant_pal:=cant_pal + 1 
				total_car:=total_car + 1
				car_vali:=car_vali + 1
				grabar(salida,v)
			finsi 
			si  v en digitos entonces 
				total_car:=total_car + 1
				car_vali:=car_vali + 1
				aux:=entero(v)
				numero:=numerito(aux)
				grabar(salida,numero)
			sino 
				si v en vocal entonces 
					vocales()
				sino 
					si v en consonante entonces 
						grabar(SALIDA,"1")
						car_vali:=car_vali + 1
						total_car:=total_car + 1
					sino 
						car_inv:=car_inv + 1 
						total_car:=total_car + 1
						grabar(SALIDA,"@")
					finsi 
				finsi 
			finsi 
			avz(SEC,v)
		FINMIENTRAS
		informar()
		cerrar(sec)
		cerrar(SALIDA)
finaccion 




