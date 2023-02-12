def NumeroEntero():
    estabien = False
    numero = 0
    while (not estabien):
        try:
            numero = int(input("Elige una opcion: "))
            estabien = True
        except ValueError:
            print("Error, Valor no correcto")
    return numero

Nombre = []
Actores = []
Year = []
Genero = []


def CargaArchivos():
    Delimitador = ";"
    print("Bienvenido a la carga de archivos.")
    # ruta = input("Por favor ingrese la ruta del archivo: \n")
    # print("Ruta: " + ruta)
    try:
        with open("ArchivoPrueba.lfp") as archivo:
            for linea in archivo:
                # print(linea)
                # nueva1 = linea.replace(" ", "")
                nueva2 = linea.replace("\n", "")
                temp = nueva2.split(Delimitador)
                # print(temp)
                for i in range(len(Nombre)):
                    if Nombre[i] == temp[0]:
                        Nombre.pop(i)
                        Actores.pop(i)
                        Year.pop(i)
                        Genero.pop(i)
                        break
                Nombre.append(temp[0])
                Actores.append(temp[1])
                Year.append(temp[2])
                Genero.append(temp[3])

    except:
        print("Error prra, revise la ruta")


print("*-----------------------------------------*")
print("*  Lenguajes Formales y de Programación   *")
print("*  Seccion: B                             *")
print("*  Carnet: 201901907                      *")
print("*  Byron Estuardo Caal Catun              *")
print("*-----------------------------------------*")

input("Presione cualquier tecla para continuar\n")
estamos = False

while not estamos:
    print("*---------------------------------*")
    print("*  1. Cargar archivo de entrada   *")
    print("*  2. Gestionar peliculas         *")
    print("*  3. Filtrado                    *")
    print("*  4. Grafica                     *")
    print("*  5. Salir                       *")
    print("*---------------------------------*")

    opcion = int(input("Elige una opcion: \n"))

    if opcion == 1:
        print("Entramos a Carga de archivo")
        CargaArchivos()
        print("Se guardo correctamente el archivo!")
        # print(Nombre)
        # print(Actores)
        # print(Year)
        # print(Genero)
    elif opcion == 2:
        print("Entramos Gestionar")
    elif opcion == 3:
        print("Entramos Filtrado")
    elif opcion == 4:
        print("Entramos Grafica")
    elif opcion == 5:
        print("Salimos prro")
        estamos = True
    else:
        print("Por favor ingrese un numero del 1 al 5")
print("Termino como la relacion con mi ex")
