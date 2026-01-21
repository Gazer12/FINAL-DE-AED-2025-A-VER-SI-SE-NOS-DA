accion es 
	 ambiente 
	 	 sec,sal1,sal2:secuencia de caracter 
		 v,v1,v2:caracter
		 A:arreglo [1..13] de caracter 
		 B:arreglo [1..13] de caracter
		 C:arreglo [1..4] de entero 
		 procedimeinto puesta() es
		 	para i:=1 a 13 hacer 
				A[i]:=nil
			fin para 
		 	para j:=1 a 13 hacer 
				A[j]:=nil
			fin para 
		 finprocedimeinto 
		 funcion conv(x:caracter):entero 
		 	segun x hacer 
				='A'	:conv:=1
				='B'	:conv:=2
				='C'	:conv:=3
				='D'	:conv:=4
			finsegun 
		finfuncion 
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
		bandera1,bandera2:booleano
		funcion mostrar(x:caracter):caracter
		 	segun x hacer 
				'1'	conv:=A
				'2'	conv:=B
				'3'	conv:=C
				'4'	conv:=D
			finsegun 
		finfuncion 
		proceso 
		ARR(sec)
		avz(sec,v)
		crear(sal1)
		crear(sal2)
		puesta()
		bandera1:=falso
		bandera2:=falso
		cant1:=0
		cant2:=0
		total:=0
		max:=LV
		mientras (v<>*) hacer 
			para i:=1 a 7 hacer 
				A:=[i]:=v
				B:=[i]:=v
				avz(sec,v)
			finpara
			C[conv(v)]:=C[conv(v)] + 1
			A[8]:=v
			B[8]:=v
			avz(sec,v)
			aux2:=entero(v)
			si (v > aux) entonces 
				bandera2:=verdader 
			finsi
			para i:=1 a 10 hacer 
				A[i]:=v
				B[i]:=v
				avz(sec,v)
			finpara 
			si (v="Z") entonces 
				bandera1:verdadero 
			finsi 
			A[11]:=v
			B[11]:=v
			avz(sec,v)
			aux:=entero(v)
			si aux=1 enotnces 
				bandera1:=verdader 
				A[12]:=v
				B[12]:=v
			sino 
				bandera1:=falso 
				B:=[12]:=v
			finsi 
			avz(sec,v)
			A[13]:=v
			B[13]:=v
			total:=total + 1
			si bandera1 entonces 
				para i:=1 a 13 hacer 
					grabar(sal1,A[i])
				finpara 
				para i:=1 a 13 hacer 
					A[i]:=nil
				finpara 
				bandera1:=falso 
				cant1:=cant1 + 1
			sino 
				para i:=1 a 13 hacer 
					A[i]:=nil
				finpara
			finsi 

			si bandera2 entonces 
				para i:=1 a 13 hacer 
					grabar(sal2,B[i])
				finpara 
				para i:=1 a 13 hacer 
					B[i]=0
				finpara 
				bandera2:falso
				cant2:=cant2 + 1
			sino 
				para i:=1 a 13 hacer 
					B[i]=0
				finpara 
			finsi 
			avz(sec,v)
		finmientras 

		para i:=1 a 4 hacer 
			si max > C[i] entonces 
				max:=C[i]
				tipo:i
			finsi 
		finpara 

	esc("la cantidad de autos que pasaron por la zona z1 es", cant1, "y el procentaje es "((cant1/total)*100))
	esc("la cantidad en infracion es", cant2,"y el procentaje es"((cant2/total)*100))
	esc("el tipo de auto", mostrar(tipo), "con un total de", max)

	cerrar(sec)
	cerrar(sal1)
	cerrar(sal2)
finaccion

			


				


			



			
			









		 


		






