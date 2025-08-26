ACCION es 
	ambiente 
	//id(5)titulo finaliza con &num[01..39]num[0..2]estado[v,r,s]
	sec:secuencia de caracter 
	v:caracter 

	info=regsitro 
		id:N(5)
		titulo:AN(30)
		area:[1...39]
		tipo_proyecto:[0..2]
		estado:["v","r","s"]
		prox:puntero a info 
	finregistro
	prim,q,p:puntero a info 

	cosa=registro 
		nombre_area:N(2)
		id:N(5)
	finregistro 
	arch:archivo de cosa 
	reg:cosa 
	proceso 
		avz(sec,v)
		mientras NDS(sec) hacer 
			si v

