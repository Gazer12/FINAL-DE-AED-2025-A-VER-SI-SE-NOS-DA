accion par es 

socursal=regsitro 
	cod_socursal:N(30)
	año_vta:[2001...2024]//creo que asi se definia un rango
	total:N(100)
	cantidad_facturas:N(4)
finregistro 
arch:archivo de socursal
reg:socursal

res_socursal
res_año


procedimiento corte_socursal()es 
	corte_año()
	total_general:=total_general + total_s
	si (total_s > 1.000.000) y (respaldo_año <= (respaldo_año+1)) entonces 
		esc("la sucursal",reg.sucursal,"vendio un monto de",total_s)
		total_destacado:=total_destacado + total_s
		si (total_s > max_sucursal) entonces
			max_sucursal:=total_s
			respaldo_sucursal:=reg.sucursal
		finsi
		segun reg.socursal hacer 
			1=		total1:=total1 + total_s
			2=		total2:=total2 + total_s
			3=		total3:=total3 + total_s
			4=		total4:=total4 + total_s
			5=		total5:=total5 + total_s
		finsegun
		total_s:=0
		res_socursal:=reg.sucursal
finprocedimiento

procedimiento corte_año() es 
	corte_socursal()
	prinf("en el año", reg.año,"se facturo",reg.cantidad_facturas)
	pepe:=reg.total
	total_s:=total_s + pepe
	pepe:=0
	respaldo_año:=reg.año

	si reg.año > max_año entonces
		max_año:=reg.año
		socursal:=reg.cod_socursal
	finsi 
	res_año:=reg.año
finprocedimiento


proceso 
abrir e/(arch)
leer(arch,reg)
res_socursal:=reg.cod_socursal
total:=0
total1:=0
total2:=0
total3:=0
total4:=0
total5:=0
total_destacado:=0
total_general:=0
total_s:=0
max_sucursal:=0
max_año:=0
mientras NFDA(arch) hacer 
	si res_socursal<> reg.socursal entonces 
		corte_socursal()
		total
	sino 
		si res_año<>reg.año_vta entonces
			corte_año()
		finisi
	finsi 
	leer(arch,reg)
finmientas
cerrar(arch)
esc("la sucursal con mayor facturacion total es",respaldo_sucursal,"con un monto de",max_sucursal)
esc("el promedio de ventas por la socursal es" ,1,"es" ((total1/total_general)*100) )
esc("el promedio de ventas por la socursal es" ,2,"es" ((total2/total_general)*100) )
esc("el promedio de ventas por la socursal es" ,3,"es" ((total3/total_general)*100) )
esc("el promedio de ventas por la socursal es" ,4,"es" ((total4/total_general)*100) )
esc("el promedio de ventas por la socursal es" ,5,"es" ((total5/total_general)*100) )
esc("el total general vedido es", total_general)
esc("el procentaje destado es", ((total_destacado/total_general)*100))

finaccion
//git status
//git add .
//git commit -m "Agrego el algoritmo de promedio"
//git push origin main
//si no anda nada a lo bruto git push origin main --force




