import abc

class Chain(metaclass=abc.ABCMeta):
    """
    Este objeto e a base do Chain of Responsability
    """

    def __init__(self, context=None):
        self.next = None
        self.context = context

    def addNext(self, next_chain):
        self.next = next_chain
        return next_chain

    def start(self):
        if self.execute():
            if self.next:
                self.next.context = self.context
                self.next.start()

    @abc.abstractmethod
    def execute(self):
        """
        Colocar aqui o codigo consultando o contexto. 
        Retornar False quando desejar parar
        """
