# Unvoid Chess

Unvoid Chess é um jogo de xadrez personalizado desenvolvido em Go, com regras e peças exclusivas. Nesta implementação:
  
- **ProductOwner (PO):** Movimenta-se como o rei (uma casa em qualquer direção).  
- **Developer (Torre):** Movimenta-se na horizontal e vertical até 3 casas, obedecendo às restrições de bloqueio.  
- **Designer (Cavalo):** Movimenta-se em "L", similar ao cavalo do xadrez clássico.

A aplicação foi estruturada para ser extensível, de fácil manutenção e leitura, atendendo os critérios de avaliação de código (passagem do spec, legibilidade, eficiência, etc).

---

## Índice

- [Requisitos e Critérios de Avaliação](#requisitos-e-critérios-de-avaliação)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Descrição dos Pacotes](#descrição-dos-pacotes)
- [Regras e Movimentações das Peças](#regras-e-movimentações-das-peças)
- [Como Rodar a Aplicação](#como-rodar-a-aplicação)
- [Exemplos de Comandos e Testes](#exemplos-de-comandos-e-testes)

---

## Requisitos e Critérios de Avaliação

- **Passagem do Spec:** O jogo deve permitir que cada peça se movimente de acordo com a regra definida e o vencedor é o jogador que captura o `ProductOwner` adversário.
- **Legibilidade e Manutenção:** O código está dividido em pacotes com responsabilidades bem definidas.
- **Extensibilidade:** Novas peças ou regras podem ser adicionadas sem modificar a estrutura base.
- **Performance e Eficiência:** Validações e cálculos de movimentação foram feitos com eficiência, utilizando arrays e loops controlados.
- **Estilo de Programação:** A implementação adota um estilo funcional e declarativo, com foco em modularidade.

---

## Estrutura do Projeto

```
Unvoid-Chess/
│
├── cmd/
│   └── main.go           // Entrada principal que chama game.Run()
│
├── internal/
│   ├── board/
│   │   └── board.go      // Representação e lógica do tabuleiro (inicialização, exibição, validação de coordenadas)
│   │
│   ├── game/
│   │   ├── cli.go        // Loop principal de comando, processamento de comandos move, restart, help, exit
│   │   └── game.go       // Estrutura Game (tabuleiro, turno atual e método para alternar turno)
│   │
│   └── pieces/
│       ├── piece.go      // Interface da peça e definições básicas (cores, tipos de peças)
│       ├── product_owner.go  // Representa o ProductOwner (movimento 1 casa por vez em qualquer direção)
│       ├── developer.go  // Representa o Developer que atua como Torre (movimento horizontal/vertical até 3 casas)
│       └── designer.go   // Representa o Designer que atua como Cavalo (movimentos em "L")
│
├── go.mod                // Módulo do Go
└── README.md             // Documentação deste arquivo
```

---

## Descrição dos Pacotes

### cmd
- **main.go:** Inicializa a aplicação chamando `game.Run()`.

### internal/board
- **board.go:**  
  - Cria o tabuleiro com dimensões definidas pelo usuário (largura e altura entre 6 e 12).  
  - Inicializa as peças nas posições iniciais:
    - PO branco na posição A1  
    - PO preto na posição H6 (ou conforme os parâmetros de tamanho definidos)
  - Possui métodos para mostrar o tabuleiro (`Display()`), validar movimentos e converter coordenadas (ex.: "A1" para índices).

### internal/game
- **game.go:** Define a estrutura do jogo, contendo o tabuleiro e o turno atual.
- **cli.go:** Contém o loop principal para lidar com os comandos do usuário:
  - Pergunta as dimensões do tabuleiro no início;
  - Aceita comandos do tipo `move <from> <to>`, `help`, `restart`, `exit`;
  - Alterna o turno e trata a vitória quando o PO adversário é capturado.

### internal/pieces
- **piece.go:** Define a interface `Piece` e constantes para cores e tipos de peças.
- **product_owner.go:**  
  - Define a peça `ProductOwner` que se move uma casa em qualquer direção (semelhante ao movimento do rei).
  - Exemplo de símbolo: ♔ (branco) e ♚ (preto).
- **developer.go:**  
  - Representa o Developer, funcionando como Torre.  
  - Movimenta-se horizontal ou verticalmente até 3 casas, parando se houver bloqueio.  
  - Símbolos: ♖ (branco) e ♜ (preto).
- **designer.go:**  
  - Representa o Designer, funcionando como Cavalo.  
  - Se move em "L" (movimentos do cavalo do xadrez).  
  - Símbolos: ♘ (branco) e ♞ (preto).

---

## Regras e Movimentações das Peças

- **ProductOwner:**  
  - Só pode mover uma casa por vez na diagonal, horizontal ou vertical.
  - Pode capturar qualquer peça que esteja na casa de destino.
  
- **Developer (Torre):**  
  - Pode mover-se na horizontal ou vertical até 3 casas.
  - Não pode ultrapassar um obstáculo; se encontra uma peça adversária, pode capturá-la e encerra a movimentação.
  
- **Designer (Cavalo):**  
  - Movimento em “L”: 2 casas em uma direção e 1 casa perpendicular.
  - Pode pular sobre outras peças, mas o destino deve estar vazio ou conter uma peça adversária.

---

## Como Rodar a Aplicação

1. **Pré-requisitos:**  
   - [Go](https://golang.org/doc/install) instalado na máquina.

2. **Clonando o Repositório:**  
   Clone ou baixe o código-fonte do projeto.

3. **Compilando e Executando:**

   Abra o terminal/PowerShell e navegue até a pasta do projeto, então execute:
   ```
   go run cmd/main.go
   ```

4. **Configuração Inicial:**  
   Ao iniciar, o jogo solicitará as dimensões do tabuleiro:
   - Informe uma largura entre 6 e 12.
   - Informe uma altura entre 6 e 12.
  
   Exemplo:
   ```
   Enter board width (X): 8
   Enter board height (Y): 6
   Starting match on an 8x6 board...
   ```

---

## Exemplos de Comandos e Testes

- **Mover uma peça:**  
  Use o comando:
  ```
  move <origem> <destino>
  ```
  Exemplo para o ProductOwner branco mover de A1 para B2:
  ```
  move A1 B2
  ```

- **Comandos adicionais:**  
  - `help`: Mostra os comandos disponíveis.  
  - `restart`: Reinicia o jogo, pedindo as dimensões novamente.  
  - `exit`: Encerra o jogo.

- **Exemplo de Sequência de Movimentos para Vitória (White vence):**

  Com o PO branco começando em A1 e PO preto em H6, uma sequência possível seria:
  ```
  move A1 B2
  move H6 G5
  move B2 C3
  move G5 F4
  move C3 D4
  move F4 E4
  move D4 E5
  move E4 F4
  move E5 F6     // Captura o PO preto e vence
  ```

- **Testando Developer e Designer Movimentos:**  
  - **Developer (Torre):**  
    Exemplo: Se o Developer estiver na célula D3, você pode testar:
    ```
    move D3 D6     // Movimento vertical para até 3 casas, desde que o caminho esteja livre
    ```
  - **Designer (Cavalo):**  
    Exemplo: Se o Designer estiver em B1, teste:
    ```
    move B1 C4     // Movimento em "L" (2 casas em uma direção, 1 em outra)
    ```

---

## Considerações Finais

- **Validações:**  
  Cada movimento é validado conforme a lógica específica de cada peça. Movimentos que ultrapassem os limites ou que não estejam de acordo com as regras definidos resultarão em mensagens “Invalid move for this piece.”

- **Customização:**  
  O projeto foi desenvolvido com foco em legibilidade e facilidade de manutenção. Caso novas regras ou peças sejam necessárias, crie/adicione novas implementações de acordo com a interface `Piece`.

- **Feedback:**  
  As mensagens exibidas no console ajudam o jogador a entender o que está ocorrendo e a corrigir comandos inválidos.

---

Esperamos que essa documentação ajude a entender, rodar e testar a aplicação. Se surgirem dúvidas ou pontos para ajuste, sinta-se à vontade para modificar e expandir conforme os critérios de avaliação.