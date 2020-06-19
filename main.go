package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/GiterLab/urllib"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type infoInical struct {

	url string
	concurrencia int
	solicitudes int
	rutaCarga string
	timeOut int


}

type strucData struct {
	Nombre string
	Departamento string
	Edad int
	Formadecontagio string
	Estado string

}


func main() {

	fmt.Println(". . . . inicio")
	//leerData()

	/*var pac []paciente
	var prueba string = "prueba.txt"

	pac = getParams(prueba)
	fmt.Printf("pacientes : %+v", pac)*/



}

func leerData(){

	var info infoInical
	var option int = -1

	for true {

		fmt.Println("------menu-----")
		fmt.Println("1. Ingresar url:")
		fmt.Println("2. Ingresar Concurrencia:")
		fmt.Println("3. Ingresar Solicitudes:")
		fmt.Println("4. Ingresar ruta del archivo que se desea cargar:")
		fmt.Println("5. Ingresar timeout:")
		fmt.Println("6. Aceptar")
		fmt.Println("7. Mostrar info ingresada:")
		fmt.Println("8. Cancelar")
		fmt.Println("9. Salir")

		fmt.Scanf("%d", &option)

		if option == 0{
			break
		}

		optionSwitch(option, &info)

	}

}

func optionSwitch(option int, info *infoInical) {
	switch option {
	case 1:
		readUrl(&info)
		break
	case 2:
		readConcurrence(&info)
		break
	case 3:
		readSolicitodes(&info)
		break
	case 4:
		readRuta(&info)
		break
	case 5:
		readTimeOut(&info)
		break
	case 6:
		break
	case 7:
		mostrarInfo(&info)
		break
	case 8:
		break

	}

}


func readUrl(info **infoInical) {
	fmt.Println("url...")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}

	(*info).url= scanner.Text()


}

func readConcurrence(info **infoInical){
	fmt.Println("no. concurrencia...")
	var concu int
	fmt.Scanf("%d", &concu)
	(*info).concurrencia = concu

}

func readSolicitodes(info **infoInical){

	fmt.Println("no. solicitudes...")
	var sol int
	fmt.Scanf("%d", &sol)

	(*info).solicitudes = sol



}

func readTimeOut(info **infoInical){

	fmt.Println("timeout...")
	var timeOut int
	fmt.Scanf("%d", &timeOut)

	(*info).timeOut = timeOut

}

func readRuta(info **infoInical){

	fmt.Println("ruta de carga...")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	(*info).rutaCarga = scanner.Text()

}

func mostrarInfo(info **infoInical)  {

	fmt.Println("url: "+ (*info).url )
	fmt.Println("concurrencia: ", (*info).concurrencia)
	fmt.Println("solicitudes: ", (*info).solicitudes)
	fmt.Println("ruta carga: "+ (*info).rutaCarga )
	fmt.Println("timeout: ", (*info).solicitudes)

}

func validarNumero(value int, valorMenor int, valorMax int) int {


	if value < valorMenor || value > valorMax{
		return -1
	}

	return value
}

func execConcurrence(info **infoInical){

	var url string = (*info).url
	var urlParams string = (*info).rutaCarga
	var concurrence int = validarNumero((*info).concurrencia, 1, math.MaxInt64 )
	var solcitudes int = validarNumero((*info).solicitudes, 1, math.MaxInt64 )
	var time int = validarNumero((*info).timeOut, 1 ,math.MaxInt64)


	if len(strings.TrimSpace(url)) == 0{

		println("La url no puede estar vacia")
		return
	}

	if len(strings.TrimSpace(urlParams)) == 0{

		println("La ruta de carga no puede estar vacia")
		return
	}

	if concurrence == -1{

		println("La concurrencia esta fuera de rango")
		return

	}

	if solcitudes == -1{

     	println("La solicitud esta fuera de rango")
		return

	}

	if time == -1{

		println("La solicitud esta fuera de rango")
		return

	}

	//queryString := getParams(urlParams)



}

func peticionExterna(concurrence int, url string, query []strucData, solicitudes int, timeout int){

}

func getParams(url string) []strucData {

	bytesLeidos, err := ioutil.ReadFile(url)

	if err != nil{
		fmt.Println(". . .error al leer archivo")

	}

	var packs []strucData
	//repleace porque si no no se va leer bien el JSON
	contenido := strings.ReplaceAll(string(bytesLeidos),"Forma de contagio","Formadecontagio")
	json.Unmarshal([]byte(contenido), &packs);
	return packs
}

func envio_data() {

	str, err := urllib.Get("https://jsonplaceholder.typicode.com/users/1").String()
	if err != nil {
		// error
	}
	fmt.Println(str)

}