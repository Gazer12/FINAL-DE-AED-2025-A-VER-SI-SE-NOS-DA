BC123-A3SZ1+XYZ789-C5NZ2+DEF456-A2NZ1+GH789PP-C7SZ3+*
ACCION ES 
	AMBIENTE
		SEC,SALID1,SALIDA2:secuencia de caracter 
		v:carcter 

		nodo1=regsitro 
			dato:=AN(30)
			prox:puntero a nodo1 
		finregistro
		prim1:punteor a nodo 

		nodo2=regsitro 
			dato:AN(30)
			prox:puntero a nodo
		finregistro
		prim2 
		C:arreglo [1..4] de entero 

		procedimiento cargar(x,pri:punteor nodo,)//incerto uno despues de otro
			si pri=nil entonces 
				pri:=x 
				*x.prox:=nil 
			sino 
				p:=prim 
				mientras *p.prox<>nil hacer 
					a:=p
					p:=*p.prox 
				finminetras 
				*a.prox:=x 
				*x.prox:=nil 
			finsi 
		finprocedimeinto

		prcocedimiento eliminar(pri:punteor a nodo) es //elimino uno despues de otro 
			mientras pri<>nil hacer 
				aux:=pri 
				pri:=*pri.prox 
				disponer(aux)
			finmientras 
		finprocedimeinto
		procedimiento crear(x,pri:puntero a nodo) es
		
		finprocedimeinto crear() es 
			nuevo(q1)
				*q1.dato:=v 
				nuevo(q2)
				*q2.dato:=v 
				cargar(q1,prim1)
				cargar(q2,prim2)
		finprocedimeinto
		funcion entero(x:entero)->entero 
			segun x hacer 
				"0". conv:=0
				"1". conv:=1
				"2". conv:=2
				"3". conv:=3
				"4". conv:=4
				"5". conv:=5
				"6". conv:=6
				"7". conv:=7
				"8". conv:=8
				"9". conv:=9
			finsegun
		finfuncion
		funcion conv(x:caracter):entero 
		 	segun x hacer 
				='A'	:conv:=1
				='B'	:conv:=2
				='C'	:conv:=3
				='D'	:conv:=4
			finsegun 
		finfuncion
		funcion mostrar(x:caracter):caracter
		 	segun x hacer 
				'1'	conv:=A
				'2'	conv:=B
				'3'	conv:=C
				'4'	conv:=D
			finsegun 
		finfuncion  

		proceso 
		arr(sec,v)
		leer(sec,v)
		prim1:=nil
		prim2:=nil

		mientras NFDS(sec) hacer 
			meintras v="-" hacer 
				crear()
			fin
			crear()
			avz(sec,v) //parado en el tipo
			C[conv(v)]:=C[conv(v)] + 1
			crear()
			avz(sec,v)//parado en las horas
			si v>2 entonces 
				bandera1:=verdader 
			finsi 
			crear()
			avz(sec,v)//parado en la infacion
			si (v="s") y (bandera1) entonces 
				bandera2:=verdadero 
				c2:=c2 + 1 
			sino 
				si v="s" entonce
					c2:=c2+1
				finsi
			fisi 
			crear()
			avz(sec,v) //toy en la z 
			crear()
			avz(secc,v)//toy en el ultimo dijitos 
			SI (v="1") entonces 
				bandera3:=verdadero
				c:=c + 1
			finsi 
			crear()
			avz(sec,v)//toy en el + 
			C:=
			crear()
			si bandera2 entonces 
				p2=prim2
				meintaras p2<> nil hacer 
					grabar(SALIDA2,*p2.dato)
					p2:=*p2.prox 
				fin mientras 
				eliminar(prim2)
				prim2:=nil
			sino 
				eliminar(prim2)
				prim2:=nil
			finSI 
			SI bandera3 entoces 
				p1=prim1
				meintaras p1<> nil hacer 
					grabar(SALIDA1,*p1.dato)
					p1:=*p1.prox 
				fin mientras 
				eliminar(prim1)
				prim1:=nil
			sino 
				eliminar(prim1)
				prim1:=nil
			finSI 
		FINMIENTRAS

		PARA I:=1 A HACER HACER 
			total:=total + c[I]
		finpara 
		
		
		para i:=1 a 4 hacer 
			si max>C[i] entonce 
				max:=C[i]
				pos:=i
			finsi 
		finpara 
		pepe:=mostrar(pos)
		
		
		
		esc("cantidad de auto que pasaron pro z1 es ", c )
		esc("la cantidad de autos que cometion una infracciones " c2)
		esc("el tipo de autoq ue ams circulo por la ciudad es", pepe ,"con " max )
		cerrar(sec)
		cerrar(SALID1)
		cerrar(SALIDA2)
finaccion

