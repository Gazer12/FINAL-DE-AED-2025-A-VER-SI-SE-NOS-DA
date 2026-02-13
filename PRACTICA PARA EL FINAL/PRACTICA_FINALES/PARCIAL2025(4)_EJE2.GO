accion es 
	ambiente 
		fecha=regsitro 
			dia:N(2)
			mes:N(2)
			año:N(4)
		finregistro

		incripciones=regsitro 
			id_equipo:N(8)
			nombre_equipo:AN(30)
			fecha_incripcion:fecha()
			tegnologia_utilizada:arr de[1..5] de en entero 
		finregistro
		arch:=archivo de incripciones ordenado por id_equipo
		reg:incripciones
		
		puntaje=registro 
			id_equipo:N(8)
			puntaje_criterio:punto [1..3] de entero 
			puntaje_lider:N(4)
		finregistro
		arch_i:archivo indexado de puntaje ordenado por id_equipo 
		reg_i:puntaje
		
		nodo=regsitro
			nombre_equipo:AN(30)
			puntaje_total:N(4)
			puntaje_lider:N(4)
			prox,ant:puntero a nodo  
		finregsitro 
		prim.p,q:puntero a nodo 

		mas_de_3,total:entero

		procedimiento insertar() es 
			si prim=nil entonces 
				prim:=q
				ult:=q
				*q.prox:=nil 
				*q.ant:=nil 
			sino 
				p:=prim 
				ant:=nil
				mientras p<>nil y (*p.puntaje_total < *q.puntaje_total O ( *p.puntaje_total = *q.puntaje_total Y *p.puntaje_lider < *q.puntaje_lider )) hacer 
					ant:=p 
					p:=*p.prox 
				finmientras
				si prim=p entonces 
					*q.prox:=p 
					*p.ant:=q
					*q.ant:=nil 
					prim:=q 
				sino 
					si p=nil entonces 
					*q.prox:=nil 
					*q.ant:=ult 
					*ult.prox:=q 
					ult:=q 
					sino 
						*q.prox:=p 
						*q.ant:=*p.ant 
						*(*p.ant).prox:=q
						*p.ant:=q
					fisni 
				dinsi 
		finprocedimiento
		procedimiento crear() es 
			nuevo(q)
			*q.nombre_equipo:=reg.nombre_equipo
			*q.puntaje_total:=total 
			*q.puntaje_lider:=lider 
			incertar()
		finprocedimiento
		procedimiento calcular() es 
			para i:1 a 5 hacer 
				si arr[i]=1 entoces 
					tecno:=tecno + 1 
				finsi 
			fin para 
			para j:1 a 3 hacer 
				criterio:=criterio + punto[j]
			fin para 

			si tecno >= 3 entonces 
				cant:=cant + 1 
			finsi 
			total:=criterio + tecno 
			criterio:=0
			tecno:=0
		finprocedimiento
		procedimiento mostar() es 
			p:=ult
			i:=1
			mientras i<=3 Y p<>nil hacer 
				esc(i, "puesto:", *p.nombre_equipo, "puntaje="*p.puntaje_total)
				i:=i+1
				p:=*p.ant 
			finmientras
			
		finprocedimiento

		proceso 
		abrir e/(arch)
		abrir e/(arch_i)
		leer(arch,reg)
		leer(arch_i,reg_i)
		prim:=nil
		ult:=nil
		mientras nfda(arch) hacer 
			reg.id_equipo:=reg_i.id_equipo
			leer(arch_i,reg_i)
			Si existe entonces 
				lider:=reg_i.puntaje_lider
				calcular()
				crear()
			finsi 
			leer(arch,reg)
		finmientras 
		mostar()
	FINPROCESO 
FINACCION 


// ESTA MIERDA ES LO QUEHACE QUE ORDENE POR PUNTAKJE TOTAL Y SI EMPATAN ORDENA POR PIUNTAJE DE LIDER 
mientras p<>nil y ( *p.puntaje_total < *q.puntaje_total o ( *p.puntaje_total = *q.puntaje_total y *p.puntaje_lider < *q.puntaje_lider ))
hacer
//


