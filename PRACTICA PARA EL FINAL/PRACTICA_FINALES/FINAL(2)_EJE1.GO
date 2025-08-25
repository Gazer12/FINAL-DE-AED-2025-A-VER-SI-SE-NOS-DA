Accion es 
	ambiente 
		viajes=regsitro
			N_viaje:N(3)
			ciudad_origen:[1..5]
			ciudad_destino:[1..5]
			plaza_disponible:N(2)
		finregistro
		arch:archivo de viajes
		reg:viajes
		nodo=regsitro 
			datos:=viajes//preguntar a la profe si es legal
			vendidos:N(6)
			tipo: arreglo [1..2] de entero 
			prox:=puntero a nodo 
		finregistro
		prim,p,q:puntero a nodo 
		A:arreglo[1..2] de entero 

		procedimeinto verificar() es 
			si tipo="E" y (*p.tipo > 0) entonces
				*p.tipo[1]:=*p.tipo[1] - 1
				A[1]:=A[1] + 1
				*p.vendidos:=*p.vendidos + 1 
				sino 
				si tipo="V" y (*p.tipo > 0) entoces 
					*p.tipo[2]:=*p.tipo[2] - 1 
					A[2]:=A[2] + 1 
					*p.vendidos:=*p.vendidos + 1 
				sino 
					esc("cagaste capo no hay pasaje")
				finis 
			finsi 
		finprocedimiento
		total,total_genera,cant_menor,mayor:entero 
		
		proceso 
		abrir e(arch)
		leer(arch,reg)
		prim:=nil
		para i:=1 a 2 hacer 
			A[i]:=0
		finpara 
		total:=0
		total_genera:=0
		cant_menor:=0
		mayor:=HV
		Mientras NFDA(arch) hacer 
			nuevo(q)
			*q.dato:=reg
			*q.vendidos:=0
			si reg.plaza_disponible="1" entonces 
				*q.tipo[1]:=*q.tipo[1] + reg.plaza_disponible
			sino 
				*q.tipo[2]:=*q.tipo[2] + reg.plaza_disponible
			finsi 
			total:=total + 1 
			leer(arch,reg)
		finmientas
		prim:=q
		repetir 
			esc("ingrese la ciudad de origen")
			leer(origen)
			esc("ingrese la ciudad de destino")
			leer(destino)
			esc("ingrese el tipo de asiento V:vip o E:economico ")
			leer(tipo)

			p:=prim
			Mientras (p<>nil) y(*q.dato.ciudad_origen <> origen) y (*q.dato.ciudad_destino <> destino) hacer 
				p:=*p.prox 
			finmientras 

			si (*q.dato.ciudad_origen = origen) y (*q.dato.ciudad_destino = destino) entonces 
				verificar()
			finsi 
			esc("desea seguir 1:si 2:no")
			leer(opcion)
		hasta que (opcion=2)

		mientras (p<>nil) hacer 
			total_genera:=total_genera + *p.vendidos
			si (*p.dato.plaza_disponible < *p.vendidos*0.5) entonce 
				cant_menor:=cant_menor + 1
			finsi

			si mayor > *p.vendidos entonces 
				mayor:=*p.vendidos
				res_num:=*p.dato.N_viaje
			finsi 
			p:=*p.prox 
		finmientas
		cerrar(arch)
		esc("la cantidada de vuelos vip son:",A[2], "y economico son"A[1])
		esc("el numero de viaje con mayor boleto vendido es", res_num)
		esc("el procenta es de " ((cant_menor/total)*100))
		esc("total general de voleto vendido es de", total_genera)
finaccion




			



