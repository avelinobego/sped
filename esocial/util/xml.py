class Builder(object):
    def __init__(self):
        self.valor = ""
        self.value = ""
        self.text = ""
        self.atributos = None
        self.root = None

    def __getattr__(self, name):
        self.__root = name

        def _build(atributos=None):
            self.atributos = "".join(f" {k}=\"{v}\"" for (k, v) in (
                atributos.items() if atributos != None else {}))
        self.root = name
        return _build

    def __str__(self):
        return f"<{self.root}{self.atributos}>{self.value}</{self.root}>"

    def __add__(self, other):
        return str(self) + str(other)

    def __lshift__(self, other):
        self.value = str(other)
