# Desafio técnico DeliveryMuch

### Utilização com Docker
1. Construa o projeto a partir do diretório raiz do repositório: 
	```
	docker build . -t delivery
	```

2. Execute o container expondo a porta ```:8181```:
	```
	docker run -p 8181:8181 delivery
	```

3. Considerando execução local, envie requisões GET ao *endpoint* **/recipes/**. Exemplo:
	```
	http://127.0.0.1:8181/recipes/?i=onion,tomato
	```

### Utilização padrão
1. Exporte a chave em varíavel de ambiente:
	```
	export GIPHY_APIKEY=aXQy4rfovFa6J18bXZHt6MKJ8hQQjkXd
	```

2. Compile e execute a partir do diretório raiz:
	```
	go build
	./delivery
	```

3. Considerando execução local, envie requisões GET ao *endpoint* **/recipes/**. Exemplo:
	```
	http://127.0.0.1:8181/recipes/?i=onion,tomato
	```

### Observação
A chave de autorização à Giphy API encontra-se publicamente exposta, documentada neste arquivo README conforme requisitado. Caso necessário, basta alterar o valor informado no comando ```export``` no passo 1. ou no [Dockerfile](./Dockerfile).
