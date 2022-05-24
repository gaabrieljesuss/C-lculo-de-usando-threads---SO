# Cálculo-de-pi-usando-threads-SO

ALUNOS: GABRIEL DE JESUS SANTOS E MICLEDSON CARLOS SANTOS DE LIMA

Processador utilizado para execução dos testes: i7 11800H | 8 núcleos e 16 Threads

## Como executar:
- Clone o repositório;
- Navegue até a pasta do projeto clonado;
- Crie um arquivo chamado `.env` seguindo o modelo do arquivo `env-file.example` e defina:
    - `NUMERO_THREADS` - Representa o número de threads utilizadas para o cálculo do valor de PI;
    - `NUMERO_TERMOS` - Representa o número de termos da série utilizados para o cálculo do valor de PI.
- Por fim, execute o seguinte comando:
```
$ go run main.go
```

## Resultados Encontrados
### **1 THREAD**
```
  Número de Threads: 1
  Número de Termos: 1000000000
  Valores de pi: [
      3.1415926525880504,
      3.1415926525880504,
      3.1415926525880504,
      3.1415926525880504,
      3.1415926525880504
  ]
 Tempo médio de execução: 99.513872s
 Desvio padão: 5.376023010105659

```
### **2 THREADS** 
```
  Número de Threads: 2
  Número de Termos: 1000000000
  Valores de pi: [
      3.141592652589258,
      3.141592652589258,
      3.141592652589258,
      3.141592652589258,
      3.141592652589258
  ]
 Tempo médio de execução: 46.51430028s
 Desvio padão: 0.08000093799576222
```
### **4 THREADS**
```
  Número de Threads: 4
  Número de Termos: 1000000000
  Valores de pi: [
      3.1415926525892104,
      3.1415926525892104,
      3.1415926525892104,
      3.1415926525892104,
      3.1415926525892104
  ]
 Tempo médio de execução: 26.37716896s
 Desvio padão: 1.5294024652022349
```
### **8 THREADS** 
```
  Número de Threads: 8
  Número de Termos: 1000000000
  Valores de pi: [
      3.141592652589324,
      3.141592652589324,
      3.141592652589324,
      3.141592652589324,
      3.141592652589324
  ]
 Tempo médio de execução: 14.774088019999999s
 Desvio padão:  0.14865471946797237
```
### **16 THREADS** 
```
  Número de Threads: 16
  Número de Termos: 1000000000
  Valores de pi: [
      3.141592652590205,
      3.141592652590205,
      3.141592652590205,
      3.141592652590205,
      3.141592652590205
  ]
 Tempo médio de execução: 10.20393166s
 Desvio padão:  0.22853949908677562
```
### **32 THREADS**
```
  Número de Threads: 32
  Número de Termos: 1000000000
  Valores de pi: [
      3.141592652590108,
      3.141592652590108,
      3.141592652590108,
      3.141592652590108,
      3.141592652590108
  ]
 Tempo médio de execução: 9.73416038s
 Desvio padão:  0.17827926522369772
```
