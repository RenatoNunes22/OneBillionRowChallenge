# 🌡️ One Billion Row Challenge - Go

Implementação do desafio "One Billion Row Challenge" (1BRC) em Go.

## 📋 Sobre o Desafio

O **One Billion Row Challenge** consiste em processar um arquivo de texto contendo **1 bilhão de linhas** de medições de temperatura de estações meteorológicas, calculando as estatísticas (mínima, média e máxima) para cada estação da forma mais eficiente possível.

### Formato do Arquivo

Cada linha do arquivo `measurements.txt` segue o formato:

```
<nome da estação>;<temperatura>
```

Exemplo:

```
São Paulo;23.5
Rio de Janeiro;31.2
Curitiba;-5.3
```

### Saída Esperada

O programa deve gerar uma saída ordenada alfabeticamente com as estatísticas de cada estação:

```
{Curitiba=-5.3/15.2/28.4, Rio de Janeiro=18.5/27.3/38.1, São Paulo=12.0/22.5/35.8}
```

Formato: `{estação=min/média/max, ...}`

## 🚀 Como Executar

### Pré-requisitos

- [Go](https://golang.org/dl/) 1.21 ou superior
- [Python](https://www.python.org/downloads/) 3.x (para gerar o arquivo de teste)

### 1. Gerar o Arquivo de Medições

```bash
python3 create.py 1_000_000_000
```

Para testes menores:

```bash
python3 create.py 100_000
```

### 2. Executar o Programa

```bash
go run main.go
```

Ou compile e execute:

```bash
go build -o 1brc main.go
./1brc
```

## 📊 Benchmark

| Registros     | Tempo |
| ------------- | ----- |
| 100.000       | TBD   |
| 1.000.000     | TBD   |
| 1.000.000.000 | TBD   |

## 🛠️ Técnicas de Otimização

- [ ] Leitura com buffer otimizado
- [ ] Processamento paralelo com goroutines
- [ ] Memory-mapped files (mmap)
- [ ] Parsing manual de números (sem `strconv`)
- [ ] Hash map customizado
- [ ] Alocação de memória otimizada

## 📁 Estrutura do Projeto

```
OneBillionRowChallenge/
├── main.go              # Código principal
├── create.py            # Script para gerar measurements.txt
├── measurements.txt     # Arquivo de dados (gerado)
├── data.csv             # Dados auxiliares
└── README.md            # Este arquivo
```

## 🔗 Referências

- [Desafio Original (Java)](https://github.com/gunnarmorling/1brc)
- [One Billion Row Challenge](https://1brc.dev/)

## 📄 Licença

Este projeto está sob a licença MIT.
