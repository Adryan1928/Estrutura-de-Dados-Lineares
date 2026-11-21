class No:
    def __init__(self, valor):
        self.valor = valor
        self.proximo = None


class Pilha:
    def __init__(self):
        self.topo = None

    def empilhar(self, valor):
        novo_no = No(valor)
        novo_no.proximo = self.topo
        self.topo = novo_no

    def desempilhar(self):
        if self.esta_vazia():
            return None
        valor = self.topo.valor
        self.topo = self.topo.proximo
        return valor

    def esta_vazia(self):
        return self.topo is None


class Historico:
    def __init__(self):
        self.historico_voltar = Pilha()
        self.historico_avancar = Pilha()
        self.pagina_atual = None

    def visitar_pagina(self, pagina: str):
        if self.pagina_atual is not None:
            self.historico_voltar.empilhar(self.pagina_atual)
        self.pagina_atual = pagina
        self.historico_avancar = Pilha()

    def voltar(self):
        if not self.historico_voltar.esta_vazia():
            self.historico_avancar.empilhar(self.pagina_atual)
            self.pagina_atual = self.historico_voltar.desempilhar()

    def avancar(self):
        if not self.historico_avancar.esta_vazia():
            self.historico_voltar.empilhar(self.pagina_atual)
            self.pagina_atual = self.historico_avancar.desempilhar()


class ListaEncadeada:
    def __init__(self):
        self.cabeca = None
        self.cauda = None

    def add(self, valor):
        novo_no = No(valor)
        if not self.cabeca:
            self.cabeca = novo_no
            self.cauda = novo_no
        else:
            self.cauda.proximo = novo_no
            self.cauda = novo_no

    def pop(self):
        if not self.cabeca:
            return None
        valor = self.cabeca.valor
        self.cabeca = self.cabeca.proximo
        if not self.cabeca:
            self.cauda = None
        return valor

    def remover_por_valor(self, valor):
        atual = self.cabeca
        anterior = None
        while atual:
            if atual.valor == valor:
                if anterior:
                    anterior.proximo = atual.proximo
                else:
                    self.cabeca = atual.proximo
                if atual == self.cauda:
                    self.cauda = anterior
                return True
            anterior = atual
            atual = atual.proximo
        return False

    def exibir(self):
        atual = self.cabeca
        while atual:
            print(atual.valor)
            atual = atual.proximo


def question1():
    """"
    Questão 1 — O botão Avançar que leva à página errada (1,5)
    Uma startup lançou um navegador web próprio, e os usuários estão reclamando. Quem
    volta algumas páginas e depois abre um link novo descobre que o botão Avançar ainda
    leva a páginas antigas, que não fazem mais sentido. Além disso, o navegador trava quando
    o usuário clica em Voltar logo após abri-lo.
    Você foi contratado para reescrever o módulo de histórico. Ele deve guardar a página atual
    e oferecer três operações: visitar uma página nova, voltar e avançar. Os botões Voltar e
    Avançar nunca podem travar: quando não houver para onde ir, simplesmente mantêm a
    página atual.
    Implemente o módulo usando duas pilhas, com todas as operações em tempo constante.
    Explique, em uma frase, qual era a causa provável do defeito do botão Avançar e como sua
    solução o corrige.
    """

    historico = Historico()

    historico.visitar_pagina("Página 1")
    print(f"Página atual: {historico.pagina_atual}")

    historico.visitar_pagina("Página 2")
    print(f"Página atual: {historico.pagina_atual}")

    historico.voltar()
    print(f"Página atual após voltar: {historico.pagina_atual}")

    historico.avancar()
    print(f"Página atual após avançar: {historico.pagina_atual}")

    historico.voltar()
    print(f"Página atual após voltar novamente: {historico.pagina_atual}")

    historico.visitar_pagina("Página 3")
    print(f"Página atual após visitar nova página: {historico.pagina_atual}")

    historico.avancar()
    print(f"Página atual após tentar avançar: {historico.pagina_atual}")


    print("\nExplicação: O defeito do botão Avançar provavelmente ocorria porque o histórico de páginas antigas não era limpo ao visitar uma nova página, por isso o botão avançar leva a páginas antigas. A solução corrige isso limpando o histórico de avançar sempre que uma nova página é visitada, garantindo que o botão Avançar só leve a páginas que o usuário visitou naquela \"linha\" de navegação.")


def question2():
    """"
    Questão 2 — Pacientes que desistem da fila (1,5)
    Um posto de saúde chama os pacientes por ordem de chegada. Hoje a fila é um vetor, e
    cada chamada desloca todos os pacientes uma posição, o que deixa o sistema lento nos
    horários de pico. Há outro problema: pacientes cansados de esperar vão embora, e a
    recepção precisa tirá-los da fila pelo nome, em qualquer posição.
    Implemente uma fila com lista encadeada, mantendo referências para o primeiro e o
    último paciente, com três operações: registrar a chegada de um paciente, chamar o
    próximo e remover um paciente que desistiu. Chegada e chamada devem ser O(1).
    Garanta que a fila continue funcionando quando o paciente que desiste é o primeiro, o
    último ou o único.
    Informe a complexidade da desistência e explique por que ela não é O(1) nesta estrutura.
    """

    lista_encadeada = ListaEncadeada()

    lista_encadeada.add("Paciente 1")
    lista_encadeada.add("Paciente 2")
    lista_encadeada.add("Paciente 3")
    lista_encadeada.exibir()

    lista_encadeada.remover_por_valor("Paciente 2")
    print("\nFila após a desistência do Paciente 2:")
    lista_encadeada.exibir()

    lista_encadeada.pop()
    print("\nFila após chamar o próximo paciente:")
    lista_encadeada.exibir()

    print("\nComplexidade da desistência: O(n) no pior caso, ou seja, vai variar de O(1) a O(n), pois é necessário percorrer a lista encadeada para encontrar o paciente que desistiu, o que leva tempo proporcional ao número de pacientes na fila.")


ABERTURAS = "({["
FECHAMENTOS = ")}]"
CORRESPONDENCIAS = {')': '(', '}': '{', ']': '['}


def validar_balanceamento(texto):
    pilha = Pilha()

    for i, char in enumerate(texto):
        if char in ABERTURAS:
            pilha.empilhar((char, i))
        elif char in FECHAMENTOS:
            if pilha.esta_vazia():
                return i
            topo: tuple[str, int] = pilha.desempilhar()
            if CORRESPONDENCIAS[char] != topo[0]:
                return i

    if not pilha.esta_vazia():
        return pilha.topo.valor[1]

    return -1


def question3():
    """"
    Questão 3 — Sublinhando o erro no editor de código (1,5)
    Uma empresa desenvolve um editor de código para iniciantes. Quando o aluno esquece de
    fechar um parêntese, colchete ou chave, ou fecha com o símbolo errado, o editor deve
    sublinhar exatamente o símbolo culpado, e não apenas dizer que há um erro.
    Escreva uma função que recebe o texto do programa e, usando uma pilha, devolve a
    posição do símbolo a ser sublinhado, ou um valor especial indicando que está tudo certo.
    Regras do editor:
    Um fechamento sem abertura correspondente, ou que fecha o tipo errado, é o erro.
    Se o texto terminar com aberturas sem fechamento, sublinha-se a primeira delas, ou
    seja, a mais antiga.
    Informe a complexidade de tempo e de memória.
    """

    texto = input("Digite o texto do programa para ser balanceado ou aperte enter e use o texto de exemplo: ")
    if not texto:
        texto = "{ if (x > 0) else { return (4; }"

    resultado = validar_balanceamento(texto)
    if resultado == -1:
        print("O texto está balanceado.")
    else:
        print(f"Erro de balanceamento na posição: {resultado}; sublinhando o símbolo: '{texto[resultado]}'")


class Playlist(ListaEncadeada):
    def eliminar_repeticoes_consecutivas(self):
        atual = self.cabeca
        while atual and atual.proximo:
            if atual.valor == atual.proximo.valor:
                atual.proximo = atual.proximo.proximo
                if atual.proximo is None:
                    self.cauda = atual
            else:
                atual = atual.proximo

    def inverter_playlist(self):
        anterior = None
        atual = self.cabeca
        self.cauda = self.cabeca
        while atual:
            proximo = atual.proximo
            atual.proximo = anterior
            anterior = atual
            atual = proximo
        self.cabeca = anterior


def question4():
    """"
    Questão 4 — Playlists importadas com músicas repetidas (1,5)
    Um aplicativo de música guarda cada playlist como uma lista simplesmente encadeada,
    em que cada nó tem o título da música e a referência para a próxima. Ao importar playlists
    de outro serviço, um defeito fez algumas músicas aparecerem duplicadas ou triplicadas
    em sequência. Os usuários também pediram um botão para tocar a playlist de trás para
    frente.

    a) (0,75) Escreva uma rotina que elimine as repetições consecutivas de uma playlist, sem
    criar nós novos e sem usar estruturas auxiliares. Uma música que reaparece mais adiante,
    não colada a ela mesma, deve ser mantida. Explique por que esse caso não é tratado como
    repetição.

    b) (0,75) Escreva uma rotina que inverta a playlist no lugar, apenas trocando referências,
    em O(n) de tempo e O(1) de memória extra.
    As duas rotinas devem funcionar com playlist vazia ou com uma única música.
    """

    playlist = Playlist()
    playlist.add("Música 1")
    playlist.add("Música 2")
    playlist.add("Música 2")
    playlist.add("Música 3")
    playlist.add("Música 3")
    playlist.add("Música 4")
    playlist.add("Música 3")
    playlist.add("Música 5")
    playlist.exibir()

    playlist.eliminar_repeticoes_consecutivas()
    print("\nPlaylist após eliminar repetições consecutivas:")
    playlist.exibir()

    playlist.inverter_playlist()
    print("\nPlaylist após inverter:")
    playlist.exibir()

    print("\nExplicação: O caso em que uma música reaparece mais adiante, mas não colada a ela mesma, não deve ser tratada como repetição, já que o erro exposto na questão é sobre duplicações ou triplicações consecutivas, e não sobre músicas que aparecem novamente em posições diferentes na playlist. Portanto, apenas as repetições consecutivas são eliminadas e que esse caso de reeaparecer novamente na playlist pode ser por escolha do usuário.")


class ListaEncadeadaCircular(ListaEncadeada):
    def __init__(self):
        super().__init__()
        self.no_atual = None

    def add(self, valor):
        novo_no = No(valor)
        if not self.cabeca:
            self.cabeca = novo_no
            self.cauda = novo_no
            novo_no.proximo = novo_no
        else:
            novo_no.proximo = self.cabeca
            self.cauda.proximo = novo_no
            self.cauda = novo_no

    def proximo_servidor(self):
        if not self.cabeca:
            return None
        if not self.no_atual:
            self.no_atual = self.cabeca
        else:
            self.no_atual = self.no_atual.proximo
        return self.no_atual.valor

    def remover_por_valor(self, valor):
        if not self.cabeca:
            return False

        atual = self.cabeca
        anterior = None
        while True:
            if atual.valor == valor:
                if anterior:
                    anterior.proximo = atual.proximo
                else:
                    if atual == atual.proximo:
                        self.cabeca = None
                        self.cauda = None
                        self.no_atual = None
                        return True
                    else:
                        self.cabeca = atual.proximo
                        self.cauda.proximo = self.cabeca

                if atual == self.cauda:
                    self.cauda = anterior

                if atual == self.no_atual:
                    self.no_atual = anterior

                return True

            anterior = atual
            atual = atual.proximo

            if atual == self.cabeca:
                break

        return False

    def exibir(self):
        atual = self.cabeca

        while True:
            print(atual.valor)
            atual = atual.proximo
            if (atual == self.cabeca):
                break


def question5():
    """"
    Questão 5 — Servidor que caiu continua recebendo acessos (2,0)
    Um site de vendas distribui as requisições entre seus servidores em rodízio (round-robin):
    cada nova requisição vai para o servidor seguinte e, depois do último, volta ao primeiro.
    Durante uma promoção, um servidor caiu, mas continuou recebendo requisições porque o
    sistema não sabia retirá-lo do rodízio. Em outro incidente, remover um servidor que não
    existia deixou o balanceador preso em um laço infinito.
    Implemente o balanceador com uma lista encadeada circular, guardando apenas a
    referência para o servidor que atendeu por último. Ele deve permitir:
    adicionar um servidor em O(1); o novo servidor só atende depois de todos os que já
    estão na roda;
    obter o servidor da próxima requisição em O(1);
    remover um servidor pelo nome, informando se ele foi encontrado.
    Sua solução deve tratar: balanceador sem servidores, remoção do único servidor, remoção
    do servidor que atendeu por último e remoção de um nome inexistente, que deve terminar
    sem laço infinito.
    """

    servidores = ListaEncadeadaCircular()
    servidores.add("Servidor 1")
    servidores.add("Servidor 2")
    servidores.add("Servidor 3")
    servidores.add("Servidor 4")
    servidores.exibir()

    servidor_atual = servidores.proximo_servidor()
    print(f"\nPróximo servidor a atender: {servidor_atual}")

    servidores.remover_por_valor("Servidor 2")
    print("\nLista de servidores após remover o Servidor 2:")
    servidores.exibir()

    servidor_atual = servidores.proximo_servidor()
    print(f"\nPróximo servidor a atender: {servidor_atual}")

    has_servidor = servidores.remover_por_valor("Servidor 20")
    if (has_servidor):
        print("Servidor Removido")
    else:
        print("Servidor não existe")


class NoDeque:
    def __init__(self, valor):
        self.valor = valor
        self.anterior = None
        self.proximo = None


class Deque:
    def __init__(self):
        self.inicio = None
        self.fim = None
        self.tamanho = 0

    def adicionar_no_fim(self, valor):
        novo_no = NoDeque(valor)
        if self.fim is None:
            self.inicio = novo_no
            self.fim = novo_no
        else:
            novo_no.anterior = self.fim
            self.fim.proximo = novo_no
            self.fim = novo_no
        self.tamanho += 1

    def remover_do_inicio(self):
        if self.inicio is None:
            return None

        valor = self.inicio.valor
        self.inicio = self.inicio.proximo
        self.tamanho -= 1
        if self.inicio is None:
            self.fim = None
        else:
            self.inicio.anterior = None
        return valor

    def remover_do_fim(self):
        if self.fim is None:
            return None

        valor = self.fim.valor
        self.fim = self.fim.anterior
        self.tamanho -= 1
        if self.fim is None:
            self.inicio = None
        else:
            self.fim.proximo = None
        return valor

    def esta_vazio(self):
        return self.inicio is None


class HistoricoEditor:
    def __init__(self, limite):
        if limite < 0:
            raise ValueError("O limite deve ser maior ou igual a zero.")
        self.limite = limite
        self.acoes_feitas = Deque()
        self.acoes_desfeitas = Pilha()

    def registrar_acao(self, acao):
        self.acoes_desfeitas = Pilha()
        if self.limite == 0:
            return
        if self.acoes_feitas.tamanho == self.limite:
            self.acoes_feitas.remover_do_inicio()
        self.acoes_feitas.adicionar_no_fim(acao)

    def desfazer(self):
        if self.acoes_feitas.esta_vazio():
            return None
        acao = self.acoes_feitas.remover_do_fim()
        self.acoes_desfeitas.empilhar(acao)
        return acao

    def refazer(self):
        if self.acoes_desfeitas.esta_vazia():
            return None
        acao = self.acoes_desfeitas.desempilhar()
        self.acoes_feitas.adicionar_no_fim(acao)
        return acao


def question6():
    """"
    Questão 6 — Desfazer que consome memória demais (2,0)
    Um editor de textos on-line guarda todas as ações do usuário para os comandos Desfazer
    e Refazer. Em documentos editados por horas, esse histórico consome tanta memória que
    o navegador fecha a aba. A equipe decidiu guardar apenas as últimas L ações: quando
    chega uma ação nova com o histórico cheio, a mais antiga é descartada. Além disso, como
    em qualquer editor, uma ação nova apaga tudo o que poderia ser refeito.
    Implemente o histórico com as operações registrar ação, desfazer e refazer, todas em O(1).
    Use um deque para as ações feitas e uma pilha para as desfeitas. Você pode usar o deque
    da biblioteca da sua linguagem ou implementar o seu, mas não pode usar um recurso
    pronto de tamanho máximo automático.
    Explique por que um vetor comum, removendo a ação mais antiga pela primeira posição,
    não atende ao requisito de O(1).
    """
    limite = int(input("Digite o limite de ações do histórico (L): "))
    historico = HistoricoEditor(limite)

    for acao in ("Criar questão 1", "Criar questão 2", "Criar questão 3", "Criar questão 4"):
        historico.registrar_acao(acao)
        print(f"Ação registrada: {acao}")

    acao = historico.desfazer()
    print(f"\nAção desfeita: {acao}" if acao is not None else "\nNão há ações para desfazer.")

    acao = historico.refazer()
    print(f"Ação refeita: {acao}" if acao is not None else "Não há ações para refazer.")

    historico.desfazer()
    historico.registrar_acao("Criar questão 5")
    acao = historico.refazer()

    if acao is None:
        print(f"\nApós uma nova ação, não é possível refazer: {acao}")
    else:
        print(f"\nAção refeita: {acao}")

    print("\nAs operações são O(1): o deque mantém referências às duas pontas,e a pilha também insere/remove pelo topo. Ao atingir L ações, a mais antiga é removida da ponta inicial. Um vetor comum exigiria deslocar os demais elementos ao remover a primeira posição, operação O(n) para realocar em um novo espaço de memória.")


def main():
    print("Interface para selecionar uma questão da prova!")

    while True:
        q = input("Digite o número da questão que deseja selecionar(1-6): ")


        if q == "1":
            question1()
        elif q == "2":
            question2()
        elif q == "3":
            question3()
        elif q == "4":
            question4()
        elif q == "5":
            question5()
        elif q == "6":
            question6()
        elif q.lower() == "sair" or q.lower() == "0":
            print("Programa encerrado.")
            break
        else:
            print("Questão inválida! Por favor, selecione uma questão entre 1 e 6.")


main()