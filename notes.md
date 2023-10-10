#### 10/10/2023

CURSO Go e Gin: criando API rest com simplicidade

@01-Instalando e criando a primeira rota com Gin

@@01
Apresentação

[00:00] Olá, meu nome é Guilherme Lima e eu serei o seu instrutor neste treinamento de API Rest com Gin, Gorm e Docker. Neste curso, o que será que vamos aprender? Nós vamos aprender a criar uma API Rest utilizando esse framework tão conhecido da linguagem Go, que é o Gin, nós vamos utilizar o Postgres rodando no Docker.
[00:19] Vamos realizar o Crud completo da nossa aplicação, vamos utilizar o Gorm como ORM do nosso projeto, vamos incluir funcionalidades na nossa API, como busca, por exemplo, e, claro, seguindo boas práticas de programação.

[00:32] Fazendo este curso você será capaz de criar suas próprias APIs, seguindo boas práticas de programação, e incluindo essas funcionalidades, sabendo como nós criamos uma API com o Gin. Para quem é esse curso? Esse curso é extremamente indicado para quem nunca utilizou o Gin e nunca criou nenhum projeto utilizando esse framework.

[00:54] Então se você nunca mexeu com Gin, nunca criou uma API Rest, mas tem uma certa familiaridade com o Go, esse curso é indicado para você. Sabendo disso, vamos começar?

@@02
Preparando o ambiente

Olá!
É muito bom receber você neste curso de Criando uma API Rest com Gin.

Espero que seja uma experiência de aprendizado incrível e que possamos vencer todos os desafios juntos. Neste curso, vamos aprofundar nosso conhecimento utilizando Go e criar uma API Rest do zero, seguindo boas práticas de programação. Além disso, vamos aprender também como usar um ORM no Go.

Preparando ambiente
Para conseguir acompanhar este curso, é recomendado que você tenha o Go instalado.
Durante o treinamento, usarei o VSCode para editar o código e para tornar o desenvolvimento ainda mais fácil, recomendo a instalação da extensão do Go desenvolvida pelo time da Google.

A Alura é formada por pessoas que gostam de tecnologia e acreditam no poder da educação através dela. Somos uma comunidade que ama compartilhar conhecimento. Em caso de dúvida na instalação ou durante o curso, conte sempre com o fórum. Caso não tenha dúvidas, não deixe de participar do fórum para ajudar outras pessoas e fazer da comunidade um lugar ainda melhor!

https://golang.org/

@@03
Servidor Gin

[00:00] Vamos iniciar então os nossos estudos? Vamos começar um projeto do zero. Eu vou criar uma pasta, uma nova pasta na área de trabalho, vou chamar de "gin-api-rest".
[00:11] Eu vou usar o Visual Studio Code como editor de código. Eu vou arrastar essa pasta para o VS Code. Nós vamos iniciar o nosso projeto assim como fizemos no projeto anterior, utilizando o "go.mod". Eu vou abrir o meu terminal integrado, com "Ctrl + J", vou jogar para cima, pra lermos melhor.

[00:32] Vou colocar aqui go mod init e vou colocar o meu usuário do GitHub, go mod init github.com/, o nome do meu usuário guilhermemorails/ e o nome que eu vou hospedar depois no GitHub esse repositório, vou chamar mesmo de /api-go-gin. Você pode escolher o nome que você quiser. Dou um "Enter" e o “go.mod” foi criado, nós podemos ver aqui que temos o "go.mod" com as dependências desse projeto.

[00:58] Nós só temos que é a versão 1.16 do Go. Vou fechar com "Ctrl + J". O que queremos fazer agora é instalar o Gin na nossa máquina e subir um servidor com ele, de uma forma muito simples. Vamos pesquisar na internet sobre ele? No navegador vou digitar "gin golang". Vou dar um "Enter", nesse primeiro link, observe que ele até trás "Gin Web Framework -GitHub".

[01:23] Eu vou acessar esse primeiro link. Scrollando para baixo, aqui tem toda a documentação, tem a parte de instalação, nesse terceiro link.

[01:32] Vou selecionar essa opção e aqui temos, ele fala: “para você conseguir instalar o Gin, você precisa ter o Go em uma versão 1.13 ou superior”. Nós já temos a 1.16, então eu vou selecionar, vou copiar essa linha do go get -u github.com/gin-gonic/gin. Dei um "Ctrl + C" nessa linha, volto na nossa aplicação, "Ctrl + V" do go get. Eu quero pegar o Gin para essa aplicação que estamos criando agora.

[02:02] O que ele está fazendo? Ele está trazendo todas as dependências para que o Gin funcione no nosso projeto. Ele terminou de instalar todas as dependências, nós podemos até ver elas aqui no "go.mod".

[02:14] Repare que a primeira dependência é o Gin. Sabendo disso, que o Gin já está na nossa aplicação, o que podemos fazer? Vamos usar o Gin. Minimizei a janela com o "Ctrl +J", vamos criar um novo arquivo chamado "main.go", para de fato subir o nosso servidor. Vou utilizar o pkgm como atalho para ele criar o pacote main e a nossa função main. Agora vamos utilizar o Gin.

[02:38] Vou criar uma variável chamada r :=, que vai configurar a nossa aplicação no Gin. Então nós podemos criar uma aplicação no Gin que o nosso servidor vai ter algumas características. Eu vou utilizar a característica mais simples do Gin, que é a default, para isso eu vou utilizar r : = gin.Default.

[02:58] Repare que não está aparecendo nada aqui. O que eu vou fazer? Deixa eu voltar, deixa eu importar o Gin primeiro, import(), entre parênteses, dei um "Enter". Vamos lá, do "github.com/gin-gonic/gin", esse é o cara que eu quero importar.

[03:19] Agora sim, na nossa função main, r := gin.Default(). Eu quero criar um servidor com as configurações default dele. Agora vou, de fato, subir o meu servidor, com r.Run(). Ele vai subir o nosso servidor.

[03:38] Essas linhas de código são as linhas que nós precisamos para subir um servidor. Deixa eu colocar aqui, para baixo, para conseguirmos visualizar ambas as partes.

[03:47] Então criei, instalei o Gin com o go get e criei uma função main, falando: olha, eu quero a configuração default do Gin e quero subir esse servidor, para rodar. Vamos rodar esse servidor para ver, no terminal, go run main.go. Damos um "Enter". Observe que apareceram algumas mensagens.

[04:04] Como é a primeira vez que estamos vendo essas mensagens, vamos ver. Ele fala aqui que ele está criando, nós estamos no modo de debug e tal, que podemos utilizar aqui para visualizar como a nossa aplicação está se comportando em fases de desenvolvimento, isso está bem bacana. E ele fala aqui, na última linha: ouvindo e servindo. O nosso servidor está de pé na porta 8080.

[04:27] Repare que essa porta é a porta padrão do Gin. Se eu for no navegador, no "local host:8080", apareceu um "404 page not found", página não encontrada. Se eu voltar na nossa aplicação, repare que o nosso servidor recebeu uma requisição, então temos aqui o 404.

[04:45] “Puxa, Gui, mas eu queria subir em outra porta”. Nós podemos especificar isso na nossa função run, no nosso primeiro parâmetro. Caso não coloquemos nada, será a porta 8080. Agora eu quero uma porta diferente, eu quero a porta r.Run(":5000").

[04:59] Salvando essa aplicação, vou parar o meu servidor com "Ctrl + C". Vou rodar mais uma vez. Agora estamos rodando na porta 5000. Se eu rodar aqui, no navegador, a porta 8080, repare que não vai aparecer nada. Só que, se eu for na porta ":5000", nós temos o page not found. E recebemos na aplicação, nesse horário recebeu uma requisição 404.

[05:21] Eu vou deixar, como nós estamos rodando no default, eu vou deixar o padrão, r.Run(). Mas, se você quiser uma porta diferente, você pode usar na sua aplicação também. Vou parar o meu servidor, subir aqui o go run mais uma vez. No navegador, vou voltar para a porta 8080, atualizando.

[05:38] No nosso servidor nós recebemos o 404. Isso ficou bacana. Só que não queremos carregar uma página e exibir um 404, nós queremos que, quando chegar uma requisição, nós retornamos alguma informação, de preferência no padrão que as APIs fornecem hoje, um JSON com alguma informação ali dentro. É isso o que vamos fazer na sequência.

@@04
Primeira rota com Gin

[00:00] Nós carregamos o nosso servidor, ele está funcionando, aceitando as requisições, respondendo, só que temos até agora um "404 page not found" boring, não está legal isso aqui, precisa ser um negócio mais da hora.
[00:12] Nós queremos que uma determinada requisição retorne alguma informação, sei lá, o ID de um aluno e o nome dele. Isso seria muito incrível. Como fazemos isso? Se formos na documentação oficial do Gin, do GitHub, no "Quick start", ele tem como retornamos uma informação, um JSON.

[00:30] Então eu tenho o R, que é a instância default. Por que eu usei o R? Eu usei porque ele é da própria documentação dele mesmo, é um padrão, é uma convenção, mas poderia usar uma outra letra, estou seguindo a convenção da documentação. r.GET, esse get aqui é o que já estamos muito, super acostumados, que já utilizamos no nosso curso anterior também, os verbos HTTP, já temos uma base disso.

[00:53] Temos aqui um determinado endpoint, uma vírgula, uma função com esse (c *gin.Context), que é o contexto, apontando para o gin.Context, para o contexto que nós queremos, e temos a nossa mensagem no formato JSON.

[01:09] O que eu vou fazer? Vamos criar essa nossa função na nossa aplicação. Nós teremos - eu vou tirar o terminal para ficar mais limpo para visualizarmos, vou trazer o código para cima.

[01:20] Nós teremos o r. e o verbo que queremos utilizar. Eu tenho o verbo get, o verbo post, put, path, todos os verbos aqui e vamos começar com o verbo r.GET(""). A primeira coisa que vamos fazer, segundo a documentação, é qual é o endpoint que queremos utilizar.

[01:35] Será ("/alunos", ). Vou colocar uma vírgula depois do fechamento das aspas, isso é muito importante. Repare que, na documentação, ele tem uma função anônima, não tem nome, ela já devolve tudo isso.

[01:49] Um pouco ruim. O que eu vou fazer? Eu vou fazer um exemplo um pouco mais didático. O que essa função "/alunos", ela fará? Ela retorna todos os alunos. Então eu vou criar uma função aqui em cima, chamada func ExibeTodosAlunos(). Essa nossa função ExibeTodosAlunos, ela será passada na função debaixo, vamos chamar ela por aqui.

[02:15] r.GET("/alunos", ExibeTodosAlunos). Deixa eu tirar esses parênteses, vou deixar só assim. Essa nossa função func ExibeTodosAlunos(), repare que na própria documentação ela precisará ter um parâmetro, que é da convenção, que é esse c *gin.Context apontando para esse gin.Context.

[02:32] Vamos fazer isso também. func ExibeTodosAlunos(c *gin.Context), que vai apontar para o gin.Context. O que a nossa função, que exibe todos os alunos, ela vai ter? Nós podemos exibir uma mensagem, então c.JSON, vai pegar a nossa variável de contexto, o status code da nossa resposta, então c.JSON() e aqui vamos devolver, quando chegar uma requisição em alunos, será (200, ).

[03:02] E vai ter uma mensagem, vamos devolver uma mensagem. Essa nossa mensagem, podemos utilizar essa função gin.H, para conseguirmos exibir a nossa mensagem de todos os alunos. Vamos fazer isso na nossa função.

[03:16] Então vou colocar no nosso código (200, gin.H{}), abrindo chaves, eu vou escrever o nosso JSON. Vai ter "id" do aluno, abrindo aspas eu vou escrever "id":"Gui Lima" - ID Gui Lima não, "id":"1",. Aqui sim, "nome":"Gui Lima",. Repare que eu - vou colocar vírgula para não dar ruim.

[03:41] Repare que agora nós conseguimos quebrar, está um pouco mais fácil do que o modelo que a própria documentação sugere. O que acontece? Quando chegar uma requisição get para alunos, quem vai atender será esse cara, o ExibeTodosAlunos. Essa função ExibeTodosAlunos vai transformar o nosso contexto, vai responder um 200 e vai devolver o ID 1 e o nome Gui Lima.

[04:09] Gui, está muito fácil isso, porque, olha só, em poucas linhas você criou uma função que devolve um JSON, que já devolve um aluno, a nossa API já teria um primeiro endpoint funcionando. Vamos ver isso na prática? No terminal, vou interromper o nosso servidor, com "Ctrl + C", e vou colocar o go run main.go mais uma vez. Ele subiu.

[04:31] Vou fechar o nosso servidor, "Ctrl +J", vou voltar na nossa aplicação, no navegador, "localhost:8080". 8080 continua page not found. Repare que no nosso servidor, nós recebemos o 404. Só que seu eu colocar no endereço do navegador aquele nosso endpoint "localhost:8080/alunos", ID 1, nome Gui.

[04:53] Olha só, o que precisamos fazer? Nós temos uma rota, temos o nosso servidor default e especificamos qual é o verbo que nós queremos. Depois nós criamos o endpoint e passamos isso para uma função, que faz de fato o que nós precisamos. Essa função faz o quê?

[05:12] Ela está devolvendo um aluno, nós temos, no caso aqui, um aluno só. Eu podia criar outros alunos, mas um aluno só já nos é suficiente. Só que, se pararmos para pensar em código, observe que o nosso código main, ele faz várias coisas - eu vou minimizar o terminal só para podermos entender melhor.

[05:29] O nosso código main, ele faz o import de tudo o que precisa do Gin, depois ele tem uma função que controla: quando chegar uma requisição e for para exibir todos os alunos, vai fazer tudo isso, vai pegar, responder o status code 200 e tal. E a nossa função main faz várias coisas também.

[05:44] Então existe uma forma de conseguirmos modularizar o nosso código para que ele fique mais fácil a manutenção e a edição, e garantirmos um código de qualidade. É isso o que vamos pensar nos próximos vídeos.

@@05
Exibir aluno

Uma pessoa fez um código para exibir um aluno e resolveu pedir a sua ajuda para avaliá-lo. Veja o código que a pessoa compartilhou:
package main

import (
    "github.com/gin-gonic/gin"
)

func ExibeTodosAlunos()  {
    c.JSON(200, gin.H{
        "id":"1",
        "nome":"Rodrigo Ferreira",
    })
}

func main() {
    r := gin.Default()
    r.GET("/alunos", ExibeTodosAlunos)
    r.Run()
}
COPIAR CÓDIGO
Analisando os trechos de código acima, podemos afirmar que:

Os códigos estão corretos e a API funcionará normalmente, sem qualquer problema.
 
Alternativa correta
Há um problema no código, e esse problema é a falta de parênteses “()” para chamar a função “ExibeTodosAlunos” no endpoint de alunos.
 
Alternativa correta
Há um problema no código, e esse problema é que a função “ExibeTodosAlunos” não está recebendo um ponteiro referente ao contexto do framework Gin.
 
Certo! O não recebimento do contexto do Gin por parâmetro faz com que não seja possível retornar o JSON. O contexto é muito importante e possui informações úteis, como por exemplo, cabeçalhos, cookies, etc.

@@06
Faça como eu fiz

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito isso, excelente. Se ainda não fez, é importante que você implemente o que foi visto no vídeo para poder continuar com a próxima aula.

Caso não encontre uma solução nas perguntas feitas por alunos e alunas deste curso, para comunicar erros e tirar dúvidas de forma eficaz, clique neste link e saiba como utilizar o fórum da Alura .
Não tem dúvidas? Que tal ajudar alguém no fórum?

https://cursos.alura.com.br/comunicando-erros-e-tirando-duvidas-em-foruns-c19

@@07
O que aprendemos?

Nesta aula:
Iniciamos o projeto do zero, criando uma pasta e instalando os módulos necessários para criar uma API com Gin;
Instalamos o gin em nosso projeto com o comando go get -u github.com/gin-gonic/gin;
Criamos um endpoint, que recebe uma requisição GET, retornando um Json exibindo recursos de um aluno.
Na próxima aula:
Vamos modularizar nosso projeto, aprender como manipular informações passadas por parâmetros e criar a struct modelo do nosso projeto!