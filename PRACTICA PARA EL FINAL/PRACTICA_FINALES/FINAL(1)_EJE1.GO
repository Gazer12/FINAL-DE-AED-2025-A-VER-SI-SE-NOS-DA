Accion es 
	//numero_clinica(4)nombre_cliente_(finaliza con "-")numero_especialida[10...11]codigo_cobertura[0..9]dia_atendido(L,M,V)
	sec:secuencia de caracter
	v:caracter
	opcion:entero

	entrada=regsitro 
		N°_historial:AN(4)
		nom_paciente:AN(30)
		tipo_especialiodad:AN(3)
		tipo_cobertura:AN(3)
		dia:AN(1)
		prox:puntero a entrada 
	finregistro 
	prim_E,p,,q,ant:puntero a entrada
	
	salida=regsitro 
		N°_historial:AN(4)
		nom_paciente:AN(30)
		tipo_especialiodad:AN(3)
		tipo_cobertura:AN(3)
		dia:AN(1)
		prox:puntero a salida
	finregistro 
	prim_s,p2,q2,ant_s:puntero a salida

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
    FFuncion
	especialidad:arreglo[10..29] de entero 
	cobertura:arreglo[0...9] de entero 
	princedimeinto cargar() es 
		mientras NFDS(sec) hacer 
			para i:=1 a 4 hacer 
				num_his:=concatenar(num_his,v)
				avz(sec,v)
			finpara 
			mientras (v <> "-") hacer 
				nom:=concatenar(nom,v)
				avz(sec,v)
			finmientas
			avz(sec,v)//me salto el "-"
			para i:=1 a 2 hacer 
				aux:=conv_entero(v)
				espe:=(espe*10)+ aux 
				aux:=0
				avz(sec,v)
			finpara 
			cover:=conv_entero(v)
			avz(sec,v)
			dias:=(v)
			Cobertura[cover]:=cobertura[cover] + 1
			especialidad[espe]:=especialidad[espe] + 1 
			nuevo(q)
			*q.N°_historial:=num_his
			*q.nom_paciente:=nom
			*q.especialidad:=espe
			espe:=0
			*q.cobertura:=cover
			*q.dia:=dias
			*q.prox:=prim_E
			prim_E:=q
			avz(sec,v)
		finmientas
		cerrar(sec)
	finprocedimiento
	procedimiento carga_ordenada() es 
		si prim2=p2 entonces 
			prim2:=q2 
			*q2.prox:=p2 
		sino 
			mientras ((p2<>nil) y (*q.dia < *p2.dia ) y (*q.N°_historial < *P2.NFDS)) hacer
				*ant2.prox:=q2
				*q2.prox:=p2
			finmientas
		sino 
			si p2=nil entonces 
				*ant2:prox:=q2 
				*q2.prox:=nil
			finsi 
		finsi 
	finprocedimiento

	procedimiento pasar() es
		mientras (*q.prox <> nil) hacer 
			nuevo(q2)
			*q2.N°_historial:=*p.N°_historial
			*q2.nom_paciente:=*p.nom_paciente
			*q2.especialidad:=*p.especialidad
			*q2.cobertura:=*p.cobertura
			*q2.dia:=*p.dia
			si prim_s=nil entonces
				prim_s:=q
				*q.prox:=nil
			sino 
				carga_ordenada()
			finsi 

	finprocedimiento

	procedimiento Buscar() es 
		mientras (p2<>nil) y (*p2.nom_paciente<>nombre) hacer
			p2:=*p2.prox 
		finmientas
		si (*p2.nom_paciente=nombre) entonces 
			esc("paceinte encontrado")
			esc("la especialidad en la que se atendio es",*p2.tipo_especialiodad,"y se atendio un dia", *p2.dia)
		sino 
			esc("pasiente no encontrado")
		finsi 
	finprocedimiento

	procedimiento informar() es 
		para i:=0 a 9 hacer 
			esc("el paciente ", i , 'y tiene la cobertura de', cobertura[i] ) 
			finprocedimiento
			si menor_cober < cobertura[i] entonces 
				menor:menor_cober
				menor_cober:cobertura[i]
			finsi 
		finpara 
		
		para i:=10 a 29 hacer 
			esc("el paciente",  i "tiene la especialidad", especialida[i])
			
			si espe_mayor > especialidad[i] entonces 
				mayor:espe_mayor 
				espe_mayor:especialida[i]
			finsi 
		finpara 
	finprocedimeinto 

	proceso 

		prim_E:nil 
		prim_s:=nil
		menor_cober:LV
		espe_mayor:HV
		menor:=0
		mayor:=0
		cargar()
		repetir 
			esc("ingrese 1= para ordenar,2=para bsucar un paciente, 3=para informar")
			leer(opcion)
			segun (opcion) hacer 
				"1" = pasar()
				"2" = Buscar()
				"3" = informar()
			finsegun
			esc("desea salir si/no ")
			leer(opcion2)
		hasta opcion2="si"
	finproceso
finaccion 
		
	






