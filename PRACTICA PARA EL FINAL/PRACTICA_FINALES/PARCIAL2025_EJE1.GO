ACCION ES 
	AMBIENTE 
		PEPE=registro 
			codigo:N(5)
			titulo:AN(30) 
			costo:N(4,2)
			cod_categoria:[1..10]
			segmento:["0","1","2"]
		finregsitro 
		arch:archivo de pepe 
		reg:pepe 

		arr:arreglo [1..10] de caractere 
		can:arreglo [1..10] de entero 
		noche:=arreglo [1..10] de entero 

		nodo=registro 
			codigo:N(5)
			titulo:AN(30) 
			costo:N(4,2)
			cod_categoria:[1..10]
			segmento:["0","1","2"]
			fecha:=regsitro
				dia:N(2)
				mes:N(2)
				año:=N(4)
			finregsitro 
			prox:=puntero a nodo 
		finregsitro 
		prim,q,p:punteros a nodo 

		costo:entero  
		procedimiento carga() es 
			si prim=nil entonces 
				prim:=q 
				*q.prox:=nil 
			sino 
				p:=prim 
				mientras p<>nil y *q.costo>*p.costo hacer 
					ant:=p
					p:=*p.prox 
				finsi 
				si p=prim entonces 
					prim:=q 
					*q.prox:=p 
				sino 
					si p=nil entonces 
						*ant.prox:=q
						*q.prox:=nil 
					sino 
						*ant.prox:=q
						*q.prox:=p 
					finis 
				fins 
			finsi 
		finprocedimiento 
		procedimeinto informar() es 
			p:=prim 
			mientras (p<>nil) hacer 
				si *p.segmento=0 y arr[*p.cod_categoria] = "Deportivo" entonces 
					cant:=cant + 1 
				finsi 
				can[*p.cod_categoria]:=can[*p.cod_categoria] + 1

				si *p.segmento=2 entonces 
					noche[*p.cod_categoria]=noche[*p.cod_categoria] + 1
					cant_noche:=cant_noche + 1 
				finsi 
				p:=*p.prox 
			finmientras
			para i:=1 a  10 hacer 
				si menor < can[i] entonces 
					menor:=can[i]
				finsi 
			finpara 
			para i:=1 a 10 hacer 
				total:=total + can[i]
			finpara 
			para i:=1 a 10 hacer 
				si max > noche[i] entonces 
					max:=noche:[i]
				finsi 
			finpara 
			total_total:=(cant_noche/total)*100 
			esc("Cantidad de eventos que se han emitido en el segmento -Matinal- en  Deportivo", can)
			esc("total invertido menor es ", menor)
			esc("el procentaje es ", total_total)


		finprocedimiento
		procedimeinto incializar() es 
			para i:=1 a 10 hacer
				cant[i]:=0
			finpara 
			para i:=1 a 10 hacer
				noche[i]:=0
			finpara 
		finprocedimiento
		proceso 
		incializar()
		abrir e/(arch)
		leer(arch,reg)
		esc("ingrese un valor ")
		leer(costo)
		prim:=nil

		Mientras ndfa(arch) hacer 
			si reg.costo>costo entonces 
				nuevo(q)
				*q.codigo:=reg.codigo
				*q.titulo:=reg.titulo
				*q.costo:=reg.costo 
				*q.cod_categoria:=reg.cod_categoria
				*q.segmento:=reg.segmento
				*q.fecha:FECHA_SISTEMA 
				*q.descripcion:=arr[reg.cod_categoria] "y " reg.segmento
				carga()
			finsi 
			leer(arch,reg)
		finmientras 
		informar()
		cerrar (arch)
	finproceso 
finaccion 


			



