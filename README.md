<body>
    <h1>API Rest em Golang utilizando Gin e Gorm</h1>
    <p>Este repositório contém uma API Restful desenvolvida em Golang, utilizando os frameworks Gin Gonic e Gorm. A aplicação foi construída do zero, implementando rotas, endpoints, modelos de dados e integração com um banco de dados Postgres por meio do Docker.</p>
    <h2>Funcionalidades</h2>
    <ul>
        <li>CRUD de Alunos: Implementação completa das operações CRUD (Create, Read, Update, Delete) para entidades de alunos.</li>
        <li>Busca por CPF: Adição de uma funcionalidade que permite buscar alunos por meio do número do CPF.</li>
    </ul>
    <h2>Estrutura do Projeto</h2>
    <p>A estrutura do projeto segue o padrão de organização MVC (Model-View-Controller), distribuindo os arquivos em suas respectivas camadas para uma melhor organização e escalabilidade.</p>
    <h2>Pré-requisitos</h2>
    <p>Certifique-se de ter o Docker instalado para execução do banco de dados Postgres.</p>
    <h2>Configuração e Execução</h2>
    <ol>
        <li><strong>Clonar o repositório:</strong>
            <pre>
                git clone https://github.com/jonathantx/api-go-gin.git
            </pre>
        </li>
        <li><strong>Executar o Banco de Dados:</strong>
            <pre>
                docker-compose up -d
            </pre>
        </li>
        <li><strong>Executar a Aplicação:</strong>
            <pre>
                go run main.go
            </pre>
        </li>
    </ol>
    <p>A aplicação estará disponível em <code>http://localhost:PORT</code>, onde <code>PORT</code> é a porta configurada para a API.</p>
    <h2>Rotas Disponíveis</h2>
    <ul>
        <li><code>GET /alunos</code>: Retorna todos os alunos cadastrados.</li>
        <li><code>GET /alunos/{cpf}</code>: Retorna o aluno correspondente ao CPF fornecido.</li>
        <li><code>POST /alunos</code>: Cria um novo aluno.</li>
        <li><code>PUT /alunos/{id}</code>: Atualiza os dados de um aluno existente.</li>
        <li><code>DELETE /alunos/{id}</code>: Remove um aluno do sistema.</li>
    </ul>
    <hr>
    <br>
   <small>Desenvolvido por <b>Jonathan Teixeira</b> - (Curso Alura - Go e GIN: Criando API Rest)</small>
</body>
