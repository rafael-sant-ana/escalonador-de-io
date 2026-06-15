# Escalonador de I/O

Pequeno projeto feito para brincar com os escalonadores de I/O do livro do Silberschatz de sistemas operacionais:
- SCAN
- FCFS (First Come First Served)
- SSTF (Shortest Seek Time First)

Para executar:
```
go run .
```

NEXT STEPS:
- Colocar testes automatizados nos schedulings
- Fazer Consumer/Producer, com Producer gerando a cada Xms uma requisição aleatória, e com o Consumer sendo o nosso escalonador tentando atender as requisições, com um certo delay associado à movimentação (Yms por byte), e comparar em Z segundos como os escalonadores se contrastam entre si.
