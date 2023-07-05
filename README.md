**README.md - pt-br**

# Monitorador de Sites

Este é um projeto simples em Go (Golang) que permite monitorar o status de sites verificando seus códigos de status. Ele oferece opções para iniciar o monitoramento, visualizar os logs e sair do programa.

## Funcionalidades

- Iniciar o monitoramento de sites.
- Exibir logs com os resultados do monitoramento.
- Encerrar o programa.

## Como usar

1. Certifique-se de ter o Go (Golang) instalado em seu sistema.
2. Clone ou faça o download deste repositório.
3. No diretório do projeto, crie um arquivo "sites.txt" e coloque os sites que quer monitorar dentro do mesmo ( um site por linha ) e execute o seguinte comando para compilar e executar o programa:

   ```shell
   go run main.go
   ```

4. Siga as instruções apresentadas no menu para interagir com o programa.
5. O arquivo `sites.txt` contém a lista de sites a serem monitorados. Adicione os URLs dos sites desejados, um por linha, nesse arquivo.
6. Os logs de monitoramento serão registrados no arquivo `log.txt`.

## Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões para melhorar o monitorador de sites, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---

**README.md - en-us**

# Website Monitor

This is a simple project in Go (Golang) that allows you to monitor the status of websites by checking their status codes. It provides options to start monitoring, view logs, and exit the program.

## Features

- Start monitoring websites.
- Display logs with monitoring results.
- Exit the program.

## How to Use

1. Make sure you have Go (Golang) installed on your system.
2. Clone or download this repository.
3. In the project directory, create a "sites.txt" file and place the sites you want to monitor inside it (one site per line) and run the following command to compile and run the program:

   ```shell
   go run main.go
   ```

4. Follow the instructions presented in the menu to interact with the program.
5. The `sites.txt` file contains the list of websites to monitor. Add the URLs of the desired sites, one per line, in this file.
6. The monitoring logs will be recorded in the `log.txt` file.

## Contribution

Contributions are welcome! If you find any issues or have suggestions to improve the website monitor, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
