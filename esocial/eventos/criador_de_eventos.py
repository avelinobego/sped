# Usar a Tabela 9 para criar os esqueletos de eventos, derivados de chain

if __name__ == "__main__":
    from string import Template
    import os

    registros = None

    with open(file=r"C:\Users\aveli\git\sped\esocial\tables\Tabela 9 - Tipos de Arquivo - eSocial.txt", encoding="utf-8") as t:
        next(t)
        registros = [n.rstrip().split("|") for n in t]

    texto = """class $nome$sufixo:
    def __init__(self):
        self.codigo = "$codigo"
        self.descricao = "$descricao"
        self.identificacao = "$identificacao"
        self.data_inicio = $datinicio
        self.data_fim = $datafim
        self.id = $id
        self.orbig_pessoa_fisica = "$orbig_pessoa_fisica"
        self.orbig_pessoa_juridica = "$orbig_pessoa_juridica" """

    template = Template(texto)
    for r in registros:
        sufixo = "_" + r[0].replace("-", "")
        class_name = r[5]
        datinicio = f"\"{r[2]}\""
        datfim = f"\"{r[3]}\"" if r[3] else "None"
        codigo = r[0]
        descricao = r[1]
        identificacao = r[6]
        id = r[4]
        orbig_pessoa_fisica = r[15]
        orbig_pessoa_juridica = r[16]

        # print(f"eventos\\{class_name.capitalize()}.py")

        with open(f"eventos\\{class_name.capitalize()}{sufixo}.py", mode="x",  encoding="utf-8-sig") as f:
            print(template.substitute(
                sufixo=sufixo,
                nome=class_name.capitalize(),
                datinicio=datinicio,
                datafim=datfim,
                codigo=codigo,
                descricao=descricao,
                identificacao=identificacao,
                id=id,
                orbig_pessoa_fisica=orbig_pessoa_fisica,
                orbig_pessoa_juridica=orbig_pessoa_juridica), file=f)
