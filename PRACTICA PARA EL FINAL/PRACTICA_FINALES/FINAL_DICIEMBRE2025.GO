//numero_clinica(4)nombre_cliente_(finaliza con "-")numero_especialida[10...11]codigo_cobertura[0..9]dia_atendido(L,M,V)
ACCION ES 
	AMBITEN 	
		sec:secuencia de carcater 
		v:caracter 
		info=regsitro 
			nombre:AN(30)
			dia:N(2)
			numero_especialida:N(2)
			codigo_cobertura:N(1)
			n_historial:AN(4)
			prox:puntero a info 
		finregistro 
		prim,q,puntero a info
		especialida:arreglo de [10..29] de entero  
		cobertura: arreglo de [0..9] de entero  
		funcion conv_entero(v:caracter):entero 
			="0": conv_entero:= 0
			="1": conv_entero:= 1
			="2": conv_entero:= 2
			="3": conv_entero:= 3
			="4": conv_entero:= 4
			="5": conv_entero:= 5
			="6": conv_entero:= 6
			="7": conv_entero:= 7
			="8": conv_entero:= 8
			="9": conv_entero:= 9
    	FinFuncion
		procedimiento conv(x:carcater) es 
			aux:=conv_entero(x)
			espe:=espe*10 + aux
			aux:=0
		finprocedimeinto
		procedimiento inicializar() es 
			arr(sec)
			avz(sec,v)
			para i:=10 a 29 hacer 
				especialida[i]:=0
			finpara 
			para j:=0 a 9 hacer 
				Cobertura[j]:=0
			finpara 

		finprocedimeinto
		procedimiento crear() es 
			nuevo(q)
			*q.nombre:=nombre 
			*q.dia:=dia 
			*q.numero_especialida:=espe
			*q.n_historial:=historial
			especialida[espe]:=especialida[espe] + 1
			cobertura[covertura]:=cobertura[covertura] + 1
		finprocedimeinto
		procedimeinto carga_ordenada() es 
			
			si prim=nil entonces 
				prim:=q
				*q.prox:=nil
			sino
				aux:prim
				si prim=q entonces 
					prim:=q 
					*q.prox:=aux
				sino
					p:=prim
					mientras (p<>nil) y (*q.dia>*p.dia) y (*q.n_historial>*p.n_historial) hacer 
						ant:=p
						p:=*p.prox
					finmientas
					si p=nil entonces 
						*ant.prox:=q
						*q.prox:=nil 
					sino 
						*ant.prox:=q 
						*q.prox:=p
					finsi 
				fins 
			finsi 
		finprocedimeinto
		procedimeinto buscar() es 
			esc("ingrese el nombre del paciente a buscar")
			leer(paciente)
			a:=prim
			mientras a<>nil  y (paciente <> *a.nombre) hacer
				a:=*a.prox
			fin mientras 
			si (paciente =*a.nombre) entonces 
				esc("paceinte encontrado")
				esc("especialidad es ", *a.numero_especialida)
				esc("se atendio el dia de ", *a.dia)
			sino 
				esc("no ta tu loca aca")
			finsi 

		finprocedimeinto
		procedimiento mostrar() es
			para i:=10 a 29 hacer 
				esc("la cantidad de paciente que se atendiero de la espceialda" i ,"son " especialida[i])
				si max>cobertura[i] entonces 
					max:=cobertura[i]
					tipo_espe:=i
				finsi 
			finpara 
			para i:=0 a 9 hacer 
				esc("la cantidad de paciente que se atendiero de la cobertura" j ,"son " cobertura[j])
				si min<cobertura[j] entonces 
					min:=cobertura[j]
					tipo_cover:=j
				finsi 
			finpara 
			esc("la cobertura menos usada es", tipo_cover)
			esc("la especialidad mas usada es", tipo_espe)
		
		finprocedimeinto
		procedimiento pepe() es 
			Mientas NFDS(sec) hacer 
			para i:=1 a 4 hacer 
				historial:=concatenar(sec,v)
				avz(sec,v)
			finpara 
			mientras v<>"-" hacer 
				nombre:=concatenar(sec,v)
				avz(sec,v)
			finmientras 
			avz(sec,v)
			para i:=1 a 2 hacer 
				conv(v)
				avz(sec,v)
			finpara 
			avz(sec,v)
			covertura:=conv_entero(v)
			avz(sec,v)
			dia:=v 
			crear()
			carga_ordenada()
			avz(sec,v)
		finmientas
		finprocedimeinto

		proceso 
			inicializar()
		
			repetir
				esc("1 para cargar, 2 para buscar, 3 para informar")
				leer(tipo)
				segun tipo hacer 
					=1 :	pepe()
					=2 :	buscar()
					=3 :	informar()
				finsegun
				esc("quiere salir a la mierda si/no")
				leer(opcion)
			hasta que opcion="si"
		finproceso
finaccion









