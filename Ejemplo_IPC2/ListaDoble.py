class Nodo:
    def __init__(self, dato=None, semestre=None, meses=None):
        self.dato = dato
        self.semestre = semestre
        self.meses = meses
        self.prev = None
        self.next = None

class Lista:
    def __init__(self):
        self.cabeza = None
        self.cola = None
        self.contador = 0

    def IsEmpty(self):
        if self.cabeza is None:
            return True
        else:
            return False

    def Insertar(self, dato, meses, semestres):
        temp = Nodo(dato, semestres, meses)
        if self.IsEmpty() is True:
            self.cabeza = temp
            self.cola = temp
        else:
            temp.prev = self.cola
            self.cola.next = temp
            self.cola = temp
        self.contador += 1

    def Iterar(self):
        actual = self.cabeza
        while actual:
            dato = actual
            actual = actual.next
            yield dato

    def BuscarExiste(self, mes):
        for item in self.Iterar():
            if mes == item.meses:
                return True
        return -1

    def BuscarNodo(self, mes):
        for item in self.Iterar():
            if mes == item.meses:
                return item
        return -1

    def ObtenerIndice(self, mes):
        for item in self.Iterar():
            conta = 0
            if mes == item.meses:
                return conta
            conta += 1
        return -1

    def __getitem__(self, item):
        if item>=0 and item < self.contador:
            actual = self.cabeza
            for _ in range(item-1):
                actual = actual.next
            return actual.dato
        else:
            raise Exception("Indice no valido")

    def __setitem__(self, indice, value):
        if indice>=0 and indice < self.contador:
            actual = self.cabeza
            for _ in range(indice-1):
                actual = actual.next
            actual = value
        else:
            raise Exception("Indice no valido")

    def MostrarLista(self):
        actual = self.cabeza
        while actual:
            dato = actual
            actual = actual.next
            print("El dato es: " + dato.dato + "  El Mes es: " + str(dato.meses) + "  El semestre es: " + str(dato.semestre))




