accion es 
	ambiente 
		sec:secuencia de caracter 
		v:caracter 
		

		datos=regsitro 
			N_patente:AN(6)
			chofer:AN(60)
		finregistro
		arch:archivo de datos
		reg:datos 


		camion_chan=regsitro 
			cant_entrega:N(6)
			N_patente:AN(6)
			chofer:AN(60)
			prox:puntero a camion_chan
		finregistro
		prim,q,p,ant:puntero a camion_chan

		funcion ver(num1.num2,aux:entero)-> boleano 
			si (num1=0) entonces 
				si (num2=aux) entonce
					ver:=verdader 
				sino	
					ver:=falso
				finsi 
			sino 
				aux:=aux*10 + num1 mod 10
				ver:=ver(num1 div 10,num2,aux)
			finsi
		finfuncion
		funcion conv(x:entero)->entero 
			segun x hacer 
				"0". conv:=0
				"1". conv:=0
				"2". conv:=0
				"3". conv:=0
				"4". conv:=0
				"5". conv:=0
				"6". conv:=0
				"7". conv:=0
				"8". conv:=0
				"9". conv:=0
			finsegun
		finfuncion
	proceso 
		arr(sec)
		avz(sec,v)
		abrir e/(arch)
		leer(arch,reg)
		prim:=nil
		mientras NFDA(arch) y (v<>"*") hacer 
			mientras v<>"+" hacer 
				num:=num *10 + conv(v)
				avz(sec,v)
			finmientas
			avz(sec,v)
			si ver(num,num,0) entonce 
				nuevo(q)
				*q.cant_entrega:=num
				*q.N_patente:=reg.N_patente
				*q.chofer:=reg.chofer
			finsi 
			leer(arch,reg)
		finsi 
		prim:=q
		mientras (p<>nil) hacer 
			si prim=p entonces 
				prim:=q
				*q.prox:=p
			sino 
				p:=prim 
				ant:=nil
				mientras (p<>nil) y (*q.cant_entrega > *p.cant_entrega) hacer 
					p:=*p.prox 
				finmientas

				si p=nil entoces 
					*ant.prox:=q 
					*q.prox:=nil 
				sino 
					*ant.prox:=q 
					*q.rpx:=p 
				finsi 
			finsi 
			p:=*p.prox 
		finmientras 
		cerrar(arch)
		p:=prim 
		para i:=1 a 10 hacer 
			si (p<>nil) hacer 
				esc("en el puesto", i "se encuentras ", *p.N_patente, "nombre:", *p.chofer, "con una entrega de:", *P.cant_entrega)
				p:=*p.prox 
			finsi 
		finpara 
	finprocedimiento
finaccion 


			

