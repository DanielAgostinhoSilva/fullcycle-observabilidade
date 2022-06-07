# Curso de Obervabilidade da FullCycle

## 💻 Pré-requisitos

Antes de começar, verifique se você atendeu aos seguintes requisitos:

* Você instalou a versão mais recente de `<docker>`

## 🚀 Preparando o ambiente

Para rodar o <docker-compose.yaml>, siga estas etapas:

Antes de executar o docker-compose up, crie a rede observability com o comando
```
docker network create observability 
```

Rode o docker-compose
```
docker-compose up -d
```

Link para acessar o  [Kibana](http://localhost:5601).
