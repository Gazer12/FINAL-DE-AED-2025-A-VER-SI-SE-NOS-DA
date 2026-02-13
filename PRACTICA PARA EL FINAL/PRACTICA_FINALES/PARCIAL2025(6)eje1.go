accion es 
	ambiente 
		lista=regsitro 
			dato:N(30)
			max:=N(2)
			pos:N(2)
			min:=N(2)
			promedio:R(4)
			prox:=puntero a lista 
		finregsitro 
		prim,p,q,ant:puntero a lista 
		num,numerito,ver:entero

		procedimiento incertar() es 
			si prim=nil entocnes 
				prim:=nil 
				*q.prox:=nil 
			sino 
				p:=prim 
				mientras p<>nil hacer 
					ant:=P
					p:=*p.prox 
				finmientras
				ant:=q 
				*q.prox:=nil
			finsi 
		finprocedimeinto 
		procedimeinto poner_0() es 
			nuevo(q)
			nuevo(q)
			*q.dato:=0
			*q.max:=nil
			*q.pos:=nil
			*q.min:=nil 
			*promedio:=nil 
			incertar()

		finprocedimeinto
			

		procedimiento crear() es 
			nuevo(q)
			*q.dato:=numerito
			*q.max:=nil
			*q.pos:=nil
			*q.min:=nil
			*promedio:=nil 
			incertar()
		finprocedimeinto
		funcion max(x:entero)

		funcion prom(x:entero,suma:entero,cant:entero):real 
			si x=0 entonces 
				prom:=((suma/cant)*100)
			sino 
				suma:=suma + (x mod 10)
				prom:(x div 10,suma , cant + 1)
			finsi
				
		finfuncion

		funcion maximo(x:entero,max:entero)
			si x=0 entocnes 
				maximo:=max 
			sino 
				si (x mod 10)> max entoces 
					maximo:=maximo(x div 10, x mod 10)
				sino 
					maximo:=maximo(x div 10, max)
				finsi 
			fisni 
		finfuncion

		funcion minimo(x: entero, min: entero): entero
			si x = 0 entonces
				minimo := min
			sino
				si (x mod 10) < min entonces
					minimo := minimo(x div 10, x mod 10)
				sino
					minimo := minimo(x div 10, min)
				finsi
			finsi
		finfuncion

		funcion posicion(x:entero,max:entero,pos:entero)
			si x=0 entocnes 
				posicion:=pos
			sino 
				si (x mod 10)> max entoces 
					posicion:=posicion(x div 10, x mod 10,pos+1)
				sino 
					posicion:=posicion(x div 10, max,pos)
				finsi 
			fisni 
		finfuncion

		procedimeinto cargar() es 
			p:=prim 
			mientras p<>nil hacer 
				si *p.dato<>0 entonces 
					*p.max:=maximo(*p.dato,0)
					*P.pos:=posicion(*p.datp,0,0)
					*p.min:=minimo(*p.dato,9999)
					*p.promedio:=prom(*p.dato,0,0)
				fisni
				p:=*p.prox 
			finmientras
		finprocedimeinto 

		procedimeinto mostrar() es 
			p:=prim 
			mientras p<>nil hacer 
				si *p.dato<>0 entocnes 
					esc("el promedio del sub conjunto", cant_sub)
					esc("maximo del conjunto",*p.max, "en la poscion", *p.pos)
					esc("el minimo del conjunto", *p.min)
					p:=*p.prox 
				sino 
					futuro:=*p.prox 
					si (futuro<> nil y (*p.dato=0 y *futuro.dato=0)) entoces 
						cant_vacio:=cant_vacio + 1 
					finsi 
					cant_sub:=cant_sub + 1
					p:=*p.prox 
					futuro:=nil 
				finsi 
			finmientras
			esc("la cantidad de subconjunto son", cant_sub)
		finprocedimeinto



		proceso 
		prim:=nil
		ant:=nil
		numerito:=0
		ver:=0
		repetir
			esc("ingrese numero y use 0 para separa conjuntos y ingrese un numero negarivo para salir")
			si (num<>ver) entonces 
				ver:=num
				numerito:=numerito*10 + num 
			sino 
				si num=0 entonces
					crear()
					poner_0()
					numerito:=0
				finsi 
				esc("error no podes ingresar el mismo numero 2 veces")
				numerito:=0
			fisni 
		hasta (num<0)
		cargar()
		mostrar()
finaccion 