#### 10/10/2023

CURSO Go e Gin: criando API rest com simplicidade

```
export GO111MODULE="on" 

Error = package github.com/lib/pq is not a main package
go env -w GO111MODULE=off
go get github.com/lib/pq
go mod init

lsof -i:8080
Kill -9

react-scripts --openssl-legacy-provider start
go get github.com/stretchr/testify
go get -u gorm.io/gorm
go get github.com/lib/pq

rm -rf go.mod go.sum
go mod init "panel"  # <== replace by your own project name
go mod tidy
go mod vendor
```

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

####

@02-Modularizando o código, modelos e banco de dados

@@01
Projeto da aula anterior

Aqui você pode baixar o zip da aula 01 ou acessar os arquivos no Github!

https://github.com/alura-cursos/api_rest_gin_go/archive/refs/heads/aula_1.zip

https://github.com/alura-cursos/api_rest_gin_go/tree/aula_1

@@02
Modularizando o código

[00:00] Para tornar o nosso projeto escalável e facilitar a edição e a manutenção do nosso código, vamos dividir aqui algumas responsabilidades, que nós deixamos toda no nosso código main. O que eu vou fazer?
[00:13] Repare que nós temos esse ExibeTodosAlunos, é um código muito típico de um controller. E esse debaixo que cuida, que está lidando com as rotas, é um código que podíamos colocar em routes, um arquivo específico que só lida com as rotas. Pensando nisso, eu vou criar duas pastas, uma chamada "routes", para lidar só com as rotas.

[00:32] E uma outra pasta chamada "controllers", para lidar apenas com o que é controle da nossa aplicação. Dentro do "routes" eu vou criar um arquivo chamado "routes.go". Nesse nosso arquivo "routes.go", ele será do pacote de routes, e o que vamos fazer? Eu quero passar toda essa responsabilidade de subir o servidor default e as rotas para esse nosso arquivo. Eu vou dar um "Ctrl + X" neste código, nessa primeira linha também.

[01:05] "Ctrl + X", vou tirar ele daqui e vou passar ele para uma nova função, em "routes.go". Essa função eu vou chamar de uma função que lida com as requisições, em inglês func HandleRequests(). Vou dar um "Ctrl + V" daquele nosso código.

[01:24] Assim que eu salvo, ele traz essas outras coisas, ele traz os imports necessários, só que o ExibeTodosAlunos não está aqui, está na nossa função main. O que eu vou fazer? Eu vou tirar ele da função main, vou em "controllers" e vou criar uma pasta, um arquivo chamado "controller.go".

[01:46] Ele será do pacote de controller, package controllers, e eu terei uma função aqui, toda aquela nossa função vem para cá. Assim que eu salvo ele já traz todas as informações certas.

[01:58] Agora sim. E no "routes.go", no ExibeTodosAlunos, repare que precisaremos trazer ele de onde? Do nosso controller, então r.GET("/alunos", controllers.ExibeTodosAlunos). Assim que eu salvo, está funcionando, só que o nosso código main não faz nada. O que o main vai fazer?

[02:15] Assim que o nosso programa inicia, eu quero iniciar o nosso servidor default subindo quem lida com todas as requisições, essa função HandlerRequests, e ela está no nosso arquivo de routes, nosso pacote routes. Então em routes.HandlerRequests(), eu quero executar esse código.

[02:34] Repare que - vou parar o meu servidor, interrompi. Executando mais uma vez, na porta 8080, assim que eu atualizo o navegador eu tenho o mesmo resultado. Repare que chegaram 3 requisições. E se eu for, no navegador, para "localhost" sem ser "8080/alunos", temos o page not found 404, 404.

[02:54] Então 200 em "/alunos", que já tínhamos colocado, sucesso, e 404 caso receber uma página local. Repare que dessa forma ficará um pouco mais fácil de organizar o nosso código. Eu preciso de uma rota nova, eu crio essa rota, e quem será o controller responsável em atender todo esse código, nós já deixamos ele aqui.

[03:16] E podemos minimizar e criar outras funções, outros tipos de controle para a nossa aplicação. Isso é muito bom, porque conseguiremos garantir que o nosso código será um código escalável. “Gui, eu posso criar, dentro de "controller", um arquivo para cada função?” Você pode também, geralmente isso é o que acontece.

[03:35] Nesse nosso exemplo didático, o que eu vou fazer? Eu vou tomar a liberdade de criar só nessa nossa pasta de "controller" mesmo, então não teremos vários arquivos em "controller", para conseguirmos manter um arquivo só que exiba os alunos, um arquivo só com a responsabilidade de criar alunos, outro de deletar.

[03:52] Nós faremos isso neste nosso arquivo "controller", porque o nosso projeto tem o escopo limitado, nós sabemos até onde ele vai crescer. Então fique ligado, faça essas alterações, modularize o seu código que, na sequência, vamos tornar a nossa API Rest ainda mais incrível com o Go.

@@03
Params no Gin

[00:00] Uma ação muito comum quando estamos trabalhando com uma API Rest é passarmos um determinado valor no endereço do navegador, vou passar, sei lá, passo o meu nome, "localhost:8080/guilherme" e, de repente, quando eu dou um "Enter", exibe alguma mensagem, alguma informação.
[00:13] Ou seja, esse valor, passado como parâmetro aqui para um determinado endpoint, eu recupero esse valor e faço alguma coisa, alguma verificação do banco de dados, e assim por diante. Como fazemos isso no Go? Repare que eu não tenho ainda uma struct, eu não tenho o modelo vinculado, e eu quero mostrar isso, porque é uma ação muito importante quando vamos criar uma API Rest.

[00:35] Como pegamos esse valor passado, esse Guilherme passado, fazemos alguma coisa, por exemplo, exibimos uma mensagem na tela? Podia exibir uma mensagem de saudação, por exemplo, e o nome que passamos: “E aí, Guilherme, tudo beleza?” Vamos fazer isso? Então vamos lá. A primeira coisa que vamos fazer será criar um controller responsável que tenha esse comportamento.

[00:58] Vou chamar essa função de saudação, func Saudacao(). Ela, assim como as nossas outras funções, ela vai apontar para (c *gin.Context), esse é um padrão que temos aqui. A nossa função saudação, a primeira coisa que vamos fazer será pegar o nome que está vindo como parâmetro.

[01:18] Eu vou fechar o "main.go", não vamos precisar dele, só o "controller.go" e "routes.go". Então, o que teremos? Nós teremos mais uma requisição get, essa requisição get, eu vou colocar ela como r.GET("/") e eu preciso pegar o nome que será passado para ela. Como eu faço isso?

[01:36] Eu posso especificar assim ("/:nome", ). Vou colocar uma vírgula aqui do lado - deixa eu dar um zoom na tela para visualizarmos melhor, "Ctrl +" e vou fechar o terminal com "Ctrl + J", para conseguirmos enxergar melhor. Então, em ("/nome", ), esse valor, ele será um valor que eu poderei mudar, alterar, eu vou passar alguma informação.

[01:57] E quem será responsável por essa ação, por essa requisição get, será o ("/nome", controllers.Saudacao). Salvei esse cara, agora temos que entender como pegamos a informação que temos no parâmetro, naquele determinado endpoint? Primeira coisa, no "controller.go" eu vou criar uma nova variável, que eu vou chamar de nome, nome :=.

[02:17] Repare que tudo o que passamos por parâmetro, seja no nosso endpoint ou alguma informação que queremos pegar, quem tem esse poder é essa nossa variável C. Ela tem o contexto de tudo, então eu vou colocar := c. - repare que quando eu digito c.Params já aparece o params.

[02:34] Nos parâmetros eu posso pegar o parâmetro por nome, c.Params.ByName(""). Qual é o nome que demos mesmo? Entre aspas duplas eu vou pegar o ("nome"). Ele está dando um erro aqui, ele está falando: você declarou o nome, mas você ainda não está usando. Vamos usar ele agora.

[02:49] c.JSON(), porque eu quero devolver uma mensagem JSON. A resposta será semelhante à de cima, 200 também, e eu vou chamar o (200, gin.H{}) para criarmos a nossa mensagem. Aqui vai ser um negócio bem legal, eu vou criar uma mensagem assim: "API diz:" e vou colocar outro dois pontos, que será a nossa mensagem.

[03:16] Vou escrever assim : "E aí ", dar um espaço, vou colocar aqui fora o símbolo de mais, o nome que temos ali, vou colocar mais um símbolo de mais, para concatenar, colocar umas strings, vírgula, + nome + ", tudo beleza". O que vai acontecer?

[03:37] Olha a mensagem que vamos criar: quando alguém colocar um nome e aparecer esse nome em cima do navegador, igual eu fiz com Guilherme, vai estar escrito API - quando for nesse endpoint, no "/", a API vai falar: API diz. E o nosso conteúdo será "E aí", eu vou concatenar, eu vou juntar esse nome com essa frase toda: “E aí, tudo beleza?” Com o nome que nós passamos.

[03:59] Vamos ver. Parei tudo aqui, limpando o terminal, executando o nosso servidor mais uma vez, go run main.go. No navegador, vou utilizar esse mesmo endpoint, passando o "/guilherme". Quando eu atualizar, repare, a API diz: “E aí, guilherme, tudo beleza?”.

[04:13] Mas será mesmo, que seu mudar esse nome? Vou mudar para "/Ana". Dei um "Enter". “E aí, Ana, tudo beleza?” E se eu não passar nome nenhum? Repare que nós temos um page not found.

[04:21] Então, se colocarmos no endereço "/" e passar qualquer coisa, vou escrever "/qualquercoisa", tudo junto, repare que vai ficar escrito: "qualquercoisa", tudo beleza.

[04:30] “E aí, qualquercoisa, tudo beleza?” Dessa forma nós conseguimos pegar as informações. A primeira coisa, e o que eu quero deixar muito claro na sua mente, no seu coração, é que todo o contexto da requisição que está vindo, quem terá o controle é essa nossa propriedade C.

@@04
Struct de aluno

[00:00] Na nossa API estamos exibindo todos os alunos, aqui no caso só tem um aluno e esse aluno sou eu, criando na hora, ID 1, passando o nome. Isso não é tão legal.
[00:10] Seria legal se tivéssemos uma estrutura de alunos, uma struct de alunos. Pensando nisso, vamos criar um modelo de aluno, o que um aluno tem, quais são os dados desse aluno? O aluno, ele tem um nome, ele tem um RG e um CPF, por exemplo. Nós criamos esse modelo, depois instanciamos alguns alunos e pedimos para que esses alunos, quando essa função for invocada, ela exiba esses alunos com base no modelo que já temos, na struct que temos.

[00:37] Vamos fazer isso? A primeira coisa que eu vou fazer será criar uma pasta nova, que eu vou chamar de "models". Essa minha pasta "Models". Nessa minha pasta "models" eu vou criar um arquivo chamado "aluno.go". Esse aluno faz parte do package model, e o que eu vou fazer será criar a nossa estrutura.

[00:57] Lembrando que, como eu estou utilizando o VS Code, existem alguns atalhos que vão me auxiliar. Se eu escrevo "this", por exemplo, e dou um "Enter", ele dá o tipo, que será o nome, type Aluno struct que vamos criar, e ele já fala que é uma estrutura. Então o aluno, ele terá um nome, que será do tipo Name string e vamos concatenar através do acento grave.

[01:17] Concatenar não, no acento grave vamos falar como queremos os dados, Nome string json:. Aqui, muito importante, abre aspas duplas e eu vou escrever como eu vou vincular esse campo a resposta JSON. Eu vou vincular através de json:"nome". Repare que aqui é muito importante, nós não podemos errar. Vou dar até um "Ctrl +" para não errarmos isso.

[01:42] Acento grave, a crase invertida, json:"nome", é o nome do campo que eu vou vincular, com a letra minúscula, fecha aspa do nome e depois fecha o acento grave. E eu terei dois campos muito parecidos com esse, que será o CPF string e o RG string. Só que eu não vou chamar eles de nome, o primeiro será o json:"cpf" e o outro json:"rg".

[02:10] Então temos aqui os nossos alunos. Assim que eu salvo, ele já indenta e fica bem bonito, bem legal. Só para fins educacionais, didáticos, eu vou criar uma variável nova que eu vou chamar de alunos, que é a ideia que nós queremos. Eu quero criar uma lista, então abre e fecha o sinal de colchete, var Alunos[], eu vou chamar de var Alunos []Aluno.

[02:34] Eu vou criar uma lista de alunos, que é referente à essa struct de alunos. O que eu vou fazer? Eu vou no nosso código main, eu vou trazer essa nossa lista de alunos para conseguirmos concatenar e criar os nossos alunos. Eu vou falar que nos nossos modelos existe um cara chamado alunos, e esse alunos, o que ele é?

[02:55] Ele é uma lista, do meu model de aluno, models.Alunos = []models.Aluno{}. Eu vou criar aqui e instanciar alguns alunos. Vou dar um "Enter", deixa eu minimizar o menu lateral só para você poder ver. Criei uma lista de modelos, que eu vou chamar de alunos. O meu primeiro aluno serei eu mesmo, eu vou colocar, o aluno, ele tem um nome, o nome será {Nome: "Gui Lima", }.

[03:21] O aluno, ele também tem um CPF, o CPF é uma string também, deixa eu ver, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, , CPF: "00000000000", }, eu não sei se o CPF tem tudo isso de dígito, mas será algo parecido com isso. E o aluno também tem um RG, RG: "470000000"},. Esse é o meu RG. No final, nós colocamos uma vírgula.

[03:43] Olha que interessante, nesse RG eu esqueci de colocar aspas duplas, que ele também é uma string, agora sim. Eu posso criar aqui quantos alunos eu quiser. Eu vou criar dois alunos, será o Gui e a Ana. {Nome: "Ana",, o CPF dela invés de ser tudo isso de 0 será 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, , CPF: "11111111111",, eu acho que o CPF é isso. O RG dela, ao invés de ser o 47, será , RG: "480000000"},.

[04:07] Tenho esses dois alunos. Gui, como eu faço agora para pegar esses dois alunos e falar para o controller exibir esses dois alunos? O que vamos fazer? Vamos no nosso "controller.go".

[04:20] Está vendo que estamos fazendo aqui o c.JSON(200, gin.H) e o Gin fazendo tudo isso? Nós vamos tirar esse cara, eu vou tirar toda essa parte do gin.H e vou falar: quando chegar uma requisição para exibir todos os alunos, vá no meu modelo, no meu model, e traz um cara chamado c.JSON(200, models.Alunos).

[04:35] Salvei. Só isso? Só isso. Vou para o meu servidor, com "Ctrl + C", vou iniciar o meu servidor mais uma vez, go run main.go. Vou voltar no navegador, na nossa aplicação. Se eu colocar no endereço, só para vermos que está tudo funcionando, "localhost:8080/gui", a API diz: “e aí, gui, tudo beleza?”

[04:52] Vamos lá, eu quero saber se eu colocar "localhost:8080/alunos" vai exibir aqueles nossos dois alunos que nós criamos, o aluno Gui e o aluno Ana. Dei um "Enter" e está lá, temos o nome do aluno 1 e do aluno 2.

[05:07] Isso ficou muito incrível, ficou bem legal mesmo. Nós temos agora uma forma de conseguir estruturar esses nossos alunos. Só que nós temos um novo problema: ficou legou, temos essa estrutura de alunos, só que todas essas informações de alunos, elas estão aqui no nosso main, no nosso arquivo main.

[05:27] Isso ficou interessante, a forma que fizemos para conseguir visualizar, mas esses dados não ficam em um arquivo. O que nós fazemos? Toda a informação que passamos para um site, os seus cadastros na Aluna, os cadastros de todos os cursos que temos na Alura, eles estão em uma base de dados, e é isso que vamos começar a atacar nos próximos vídeos, vamos conectar a nossa API com um banco de dados.

@@05
Modularização do código

Durante as aulas passadas, vimos o que é a modularização do código e o que ela acarreta. Sabendo disso, podemos afirmar que:

A modularização traz muitas vantagens, como a reutilização do código, a facilidade de manutenção e a melhor legibilidade.
 
Isso! A modularização nos traz todos estes benefícios e muitos outros. Devemos sempre modular a construção de nosso software, assim obtemos todos os benefícios citados.
Alternativa correta
A modularização faz com que o compilador leia mais facilmente o código e o compile mais tranquilamente.
 
Alternativa correta
A modularização deixa os códigos mais leves e mais rápidos para serem executados. Ou seja, temos um ganho de performance da aplicação, principalmente no ambiente de produção.

@@06
Faça como eu fiz

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito isso, excelente. Se ainda não fez, é importante que você implemente o que foi visto no vídeo para poder continuar com a próxima aula, que tem como pré-requisito todo o código escrito até o momento.

Caso não encontre uma solução nas perguntas feitas por alunos e alunas deste curso, para comunicar erros e tirar dúvidas de forma eficaz, clique neste link e saiba como utilizar o fórum da Alura .
Não tem dúvidas? Que tal ajudar alguém no fórum?

https://cursos.alura.com.br/comunicando-erros-e-tirando-duvidas-em-foruns-c19

@@07
O que aprendemos?

Nesta aula:
Modularizamos nosso código criando pacotes de controles e rotas, tornando nosso código mais fácil de manter e editar;
Aprendemos como pegar informações passadas por parâmetros e retornar uma mensagem personalizada;
Criamos a struct de aluno, que vamos disponibilizar como recurso da nossa API.
Na próxima aula
Vamos aprender como conectar nosso projeto com banco de dados e exibir as informações armazenadas utilizando um Gorm!

#### 11/10/2023

@03-Struct, banco de dados e ORM

@@01
Projeto da aula anterior

Aqui você pode baixar o zip da aula 02 ou acessar os arquivos no Github!

https://github.com/alura-cursos/api_rest_gin_go/archive/refs/heads/aula_2.zip

https://github.com/alura-cursos/api_rest_gin_go/tree/aula_2

@@02
Preparando o ambiente

No próximo vídeo, vamos conectar nossa aplicação com banco de dados. Para isso, vamos utilizar uma imagem Postgres rodando no Docker.
Preparando ambiente
Para acompanhar esta aula, é recomendado que você tenha o Docker instalado.
Caso não tenha o Docker instalado e precise de ajuda:

Instalando o Docker no Windows
Instalando o Docker no Mac
Após a instalação do Docker, nos vídeos 3 e 4 desta aula, vamos executar este arquivo, para criar uma imagem do Postgres:

version: '3'
services:
  postgres:
    image: "postgres"
    environment:
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=root
      - POSTGRES_DB=root      
    ports:
      - "5432:5432"
    volumes:
      - ./postgres-data:/var/lib/postgresql/data  

  pgadmin-compose:
    image: dpage/pgadmin4
    environment:
      PGADMIN_DEFAULT_EMAIL: "gui@alura.com"
      PGADMIN_DEFAULT_PASSWORD: "123456"
    ports:
      - "54321:80"
    depends_on:
      - postgresCOPIAR CÓDIGO
Fique à vontade para alterar suas as credenciais!

https://www.docker.com/products/docker-desktop

https://cursos.alura.com.br/course/docker-e-docker-compose/task/29235

https://cursos.alura.com.br/course/docker-e-docker-compose/task/29237

https://github.com/alura-cursos/api-go-rest/blob/aula_2/docker-compose.yml

@@03
Docker e Postgres

[00:00] Os dados dos alunos que nós criamos, eles estão no nosso arquivo main, e não é essa a ideia. Geralmente as aplicações reais, elas utilizam um banco de dados, seja ele relacional ou não, para armazenar todas as informações e manter os dados persistentes no banco. É isso o que queremos fazer também.
[00:18] Assim como o exemplo do nosso curso anterior, vamos utilizar o Docker para executar e armazenar todas as nossas informações dos nossos alunos. Então o que vamos fazer? Na atividade anterior a esse vídeo você terá o script do Docker que vamos subir, com a imagem que vamos subir do Docker, para gerarmos esses alunos.

[00:38] Terá uma diferença: no curso anterior nós utilizamos um script que já inseria alguns dados, que já criava a tabela através de um script SQL. Nesse curso nós teremos uma outra abordagem, eu quero mostrar uma outra forma para vocês. O que vamos fazer? Primeira coisa, vamos criar um novo arquivo. Deixa eu sair dessas pastas.

[01:01] Vou criar um novo arquivo, que eu vou chamar de "docker-compose.yml". Esse arquivo, ele vai descrever tudo o que precisamos na nossa imagem. Eu vou dar um "Ctrl + V" e esse código todo que apareceu aqui, ele está na atividade anterior a esse vídeo.

[01:19] Só para entendermos o que esse código faz, ele fala qual é a versão o Docker que estamos utilizando e cria dois serviços, o serviço "postgres", assim como fizemos no curso anterior, com o user root, password root, DB root.

[01:34] Eu recomendo que, se você está fazendo esse curso pela primeira vez, que você não altere essas informações, para que você chegue no mesmo resultado que eu. Se você já tem uma determinada vivência, experiência, você pode ficar à vontade para alterar todas essas informações de quem será o usuário, colocar em uma variável de ambiente também.

[01:51] Ele cria duas imagens, dois serviços. O primeiro com uma imagem do Postgres e o segundo com uma imagem do pgAdmin, para conseguirmos visualizar as tabelas e as informações que estamos movimentando no nosso banco de dados. No pgAdmin ele usa a imagem, do pgAdmin, e ele tem como default esse "gui@alura.com" e essa senha, "123456".

[02:11] Lembrando que estamos utilizando o ambiente de desenvolvimento, é óbvio que no ambiente de produção nós jamais usaríamos essas senhas, você pode alterar essa senha e ela vai servir para quê? Essa senha, esse usuário, esse e-mail e essa senha, vão servir para acessarmos o pgAdmin na porta 54321:80. O banco de dados mesmo está na porta 5432:5432.

[02:34] Então vamos subir do "docker-compose.yml" duas imagens, dois serviços, um do Postgres e outro do pgAdmin. Vamos fazer isso então? Eu já tenho o meu "docker-compose.yml", o que eu vou fazer? Como vamos dar o foco nele, eu vou abrir um outro terminal e digitar docker-compose up. Dei um "Enter".

[02:53] Ele vai começar a subir, a criar essas nossas imagens, essas duas imagens. Essa ação pode demorar um pouco, então eu recomendo que depois que você fizer isso, se quiser fazer um suco de laranja, passar um café para a família, dar um abraço no gato, no cachorro, fique à vontade também.

[03:11] Ele está subindo, ele está trazendo tudo o que necessário para carregar a nossa aplicação. Deixa eu ver se ela já subiu aqui. No navegador, se eu colocar "localhost:54321", para ver o nosso pgAdmin, vamos ver o que acontece? Ainda não subiu, ele está carregando. Vamos ver, maravilha subiu, eu estou usando 200 de zoom para conseguirmos enxergar melhor.

[03:33] Aqui nós temos o pgAdmin e um caminho com e-mail - lembrando, qual é esse e-mail mesmo? Esse e-mail que nós criamos no código. Então antes de você subir o seu "docker-compose", use outro e-mail - ou se você quiser usar o meu também como base, só para aprender, está tranquilo.

[03:48] Então gui@alura.com e a senha mais difícil do universo: 123456. Eu vou tirar o zoom, dou um "Login". Observe que interessante, ele vai criar esse nosso servidor.

[04:02] Temo aqui no "Servers", deixa eu ver, aqui no "Servers" não tem absolutamente nada. Nós precisaremos criar um servidor novo. Como fazemos isso? Primeira coisa, venho no menu lateral esquerdo, no "Servers", clico com o botão direito, "Create", coloco "Server...". Eu quero criar um novo servidor que vai ficar ouvindo aquele banco de dados que nós temos.

[04:18] Para isso, o que a vamos fazer? Precisamos passar algumas informações para ele. Por exemplo, qual é o nome desse servidor? Vamos pesquisar no código o nome que vamos utilizar para ele.

[04:31] Nós demos o nome para ele de - eu vou dar um nome para esse banco de dados de "alunos", vai se chamar alunos. Ele será um servidor mesmo, eu vou deixar esse background, essa parte toda normal. Na aba "Connections" nós temos uma parte importante, o host.

[04:43] Nós precisamos falar qual é o caminho, qual é o endpoint que vai conter as informações desse banco de dados. Para isso, nós vamos fazer o quê? Eu vou criar um novo terminal no VS Code, repare que eu estou com o servidor do Go, o "docker-compose" rodando e um novo terminal, mais um terminal.

[05:00] Eu vou executar um comando para conseguirmos visualizar qual é a porta que estamos utilizando para esse nosso serviço do Postgres. Eu vou falar assim: docker-compose exec, eu quero executar o Postgres e vou colocar SH, exec postgres sh, para acessarmos essa máquina. Eu vou digitar hostname -i.

[05:31] Dou um "Enter" e ele dá um valor, 172.21.0.2. Eu vou copiar esse hostname, "Ctrl + C", vou na configuração e coloco ele no "Hostname/adress", vai ser esse IP. A porta será a porta que já estamos acostumados, a 5432, que estamos utilizando. O "username" será o "root", o username é exatamente esse que temos no código, POSTGRES_USER=root.

[05:55] O password root e o DB root também. Então user "root", "Password" "root" e o DB, cadê o DB? Não está aqui. Aqui, "Maintance", aqui também vai ser "root".

[06:06] Eu vou salvar e temos agora um servidor para os alunos. Se eu vier em "Database > root > Schemas", repare que aqui em "Tables", tabelas, não temos nada ainda.

[06:21] É isso o que vamos ver na sequência. Como é que agora nós pegamos esse banco de dados que nós conectamos no pgAdmin com o Postgres mesmo e conectamos com a nossa aplicação, para utilizarmos o banco de dados rodando no Docker.

@@04
GORM e migração dos dados

[00:00] Para facilitar a comunicação do banco de dados com a nossa aplicação GO, nós vamos usar um ORM que nós vimos no curso anterior, que é o Gorm. Só que eu quero mostrar uma outra forma de conseguirmos criar as tabelas e os registros na nossa aplicação. A primeira coisa que vamos fazer será instalar o Gorm na nossa API Go e conectar com o banco de dados.
[00:21] Para instalar o Gorm eu estou com a documentação dele aberta, se você quiser pesquisar "gorm golang", aparece esse primeiro link. Aqui tem a forma como instalamos o Gorm, é go get -u gorm.io/gorm.

[00:37] Vou copiar essa linha, no nosso código eu vou instalar ele aqui, dei o Gorm para ele instalar e ele adicionou o Gorm. Agora, o que vamos fazer será criar um pacote responsável por realizar essa comunicação com o banco de dados, essa conexão com o banco de dados. Vou criar um pacote chamado "database".

[00:58] Dentro de "database" eu vou criar um arquivo chamado "db.go". Esse arquivo "db.go", ele faz parte do pacote package database e todo o código que faz a conexão com o banco de dados, se formos na documentação do Gorm - vou colocar aqui: conectando com database, conectando com o banco de dados. Nós estamos utilizando o banco de dados Postgres, então seria esse código aqui.

[01:22] Só que eu tenho um código um pouco mais fácil, para conseguirmos visualizar, quebrando algumas etapas. Eu vou deixar esse código na descrição do vídeo e na atividade anterior, no Preparando o Ambiente. O código vai ser esse, vou dar um "Ctrl + V", que eu já tinha copiado.

[01:36] Vou salvar esse código, ele trouxe aqui tudo o que nós precisamos. Repare que no "gorm.io/driver/postgres" ele está dando uma mensagem, ele está falando: “olha, eu não encontrei nenhum driver chamado Postgres”.

[01:51] Vou segurar o "Ctrl" e vou clicar em cima do link aqui, para visualizarmos esse erro que ele fala. O que acontece? Esse erro, ele diz: “olha, você não instalou o driver do Postgres”. Então precisamos instalar. No VS Code, aqui do lado esquerdo, quando eu clico nessa linha, essa lâmpada aparece aqui.

[02:09] Quando eu clico nela, ele fala: "go get package gorm.io/driver/postgres". Caso essa lâmpada não apareça para você, você pode colocar exatamente esse comando no seu terminal, go get package gorm.io/driver/postgres. Eu vou clicar nessa lâmpada para instalar, repare que ele está pensando, ele já está fazendo essa instalação com o go get. Depois não teremos nenhuma mensagem, essa marcação vai sumir.

[02:39] Agora que ela sumiu, olha que interessante, no curso anterior nós criávamos a tabela do nosso modelo, com base na nossa struct, colocávamos alguns registros utilizando a linguagem SQL. Existe uma outra forma de realizarmos isso com o Go e com o Gorm. Nós podemos falar: eu tenho essa struct, com base nela, crie uma tabela no banco de dados para mim. E ele coloca outros campos como o ID, como quando aquele campo foi criado ou não.

[03:09] E eu quero mostrar isso na documentação para vocês. Vou voltar na documentação do Gorm, vou colocar "gorm.io/docs", declarando modelos, nesse segundo link.

[03:21] Repare que aqui ele tem uma descrição de um modelo, um user com nome, e-mail e todos os campos aqui. Ele fala de algumas convenções e podemos falar qual será a primary key, que o ID será a primary key quando queremos definir isso manualmente. Mas existe uma forma de conseguirmos declarar o Gorm “embedando” alguns códigos. Olha que interessante.

[03:44] Se formos na nossa struct e colocarmos gorm.Model, o que vai acontecer? Repare que ele tem um tipo, user, uma struct de user, que ele só tem um nome, ele falou gorm.Model. Isso é equivalente a ele colocar no ID quando aquele registro foi criado, quando o registro foi atualizado, quando o registro foi deletado e a última vez que ele foi atualizado.

[04:11] Então repare que ele vai guardar uma série de informações só de colocarmos esse gorm.Model na nossa struct. Então eu vou fazer isso, vou na nossa struct, no nosso modelo de aluno, "aluno.go" e vou falar: o nosso aluno tem um nome, um CPF, um RG e ele será do tipo gorm.Model. Vou dar um "Enter" aqui.

[04:32] Ele vai fazer um import aqui em cima para mim. Só de eu realizar essa linha ele vai inserir para nós o ID, quando foi criado, a última vez que foi atualizado e quando foi deletado. Agora, o que eu quero fazer é o seguinte, eu quero pedir para o Gorm criar essa tabela no banco de dados, no nosso Postgres. Como faremos isso?

[04:54] Aqui, na nossa string que realiza a conexão, se a conexão não tiver nenhum erro e ela for realizada com sucesso, nós podemos fazer o seguinte, vou pegar o nosso DB, DB.AutoMigrate(). Com esse auto migrate nós poderemos inserir qual é a struct que queremos inserir no nosso banco de dados.

[05:18] Para isso vou passar o endereço de memória dos nossos modelos, DB.AutoMigrate(&models.Aluno). Dessa forma ele vai falar: eu vou migrar - deixa eu salvar aqui, para ele trazer esse pacote, isso, models.Aluno, ele é uma instância dele, esqueci de pegar a instância dele, (&models.Aluno{}).

[05:39] Agora sim. Para eu criar uma tabela no banco de dados, com base na minha struct, utilizando o Gorm, eu posso executar esse comando DB.AutoMigrate(&models.Aluno{}), passando o endereço de memória da struct que eu quero criar uma instância dela. Ele vai gerar para nós, no nosso Postgres - deixa eu abrir o Postgres, ele está no "localhost:54321".

[06:07] Aqui no Postgres, deixa eu fazer a minha senha, 123456, com aquele e-mail. Ele está carregando o meu servidor, o meu "docker-compose" está de pé, deixa eu só ter certeza aqui do lado do terminal. O "docker-compose" está de pé.

[06:27] Está carregando o servidor do Postgres. O que eu vou fazer agora é conectar com o pgAdmin, vou colocar o meu e-mail, a minha senha, 123456, senha muito difícil. Vou esperar ele carregar aqui, não precisa checar, não precisa guardar a senha.

[06:43] Vou em servidores, temos o servidor de alunos e ele está pedindo a senha. A senha é "root", com base no que tínhamos criado naquele momento.

[06:51] Carregou, venho em tabelas e não tem nada. Então o que nós fizemos no nosso código? Nós criamos um pacote de database que realiza a conexão e faz esse DB.AutoMigrate, no nosso banco de dados, com base no nosso modelo de aluno. E nós inserimos esse gorm.Model, para ele incluir o ID, o created at, o updated at e o deleted at.

[07:14] O que eu preciso fazer agora é executar esse código. Quando? Quando o meu programa iniciar, eu já quero que ele tenha essa migração dos modelos no banco de dados. Então o que eu vou fazer? No meu código main, eu vou dar um "Enter". A primeira coisa que eu quero que ele faça é: vá no meu database, pegue e faça a conexão com o meu banco de dados, database.ConectarComBancoDeDados().

[07:34] Essa conexão que ele fará com o banco de dados já vai permitir que ele crie a nossa tabela de alunos no Postgres. Eu vou no terminal do Go e vou executar esse comando, go run main.go. Quando eu der um "Enter", repare o que vai acontecer, ele vai dar uma pensada e vai falar: nós estamos ouvindo na porta 8080.

[07:55] Vou no nosso banco de dados, na nossa tabela, clicar com o botão direito, dar um "Refresh". Olha que interessante que vai aparecer: não apareceu nada. Deixa eu dar um refresh no meu banco de dados. Refresh no meu database de alunos, cliquei com o botão direito, "Refresh". Venho em "Database > root", que estamos utilizando, "Schemas > Tables" e temos aqui a tabela de alunos. Quando eu clico em "alunos", olha que interessante.

[08:21] Eu vou em colunas, olha só, ele vai mostrar "cpf", "created_at", "deleted_at", "id", "nome", "rg" e "updated_at".

[08:29] Nós podemos até ver de uma outra forma, clicando com o botão direito em "alunos", eu vou em "View/Edit Data", vou ver todas as linhas e aqui teremos todos aqueles campos que eu tinha comentado com vocês. Nós temos o ID, quando foi criado, quando foi atualizado, deletado, nome, CPF e RG.

[08:46] O que vamos fazer, na sequência, é criar, de fato, um aluno, utilizando o Gorm em parceria com o Gin.

@@05
Criando alunos

[00:00] Agora que conectamos a nossa API com o nosso servidor, o que queremos fazer é criar novos alunos. Se olharmos o nosso código, nós não temos nenhum aluno registrado. Nós já temos os campos de ID, nome, CPF, mas não conseguimos gerar alunos. O que eu quero fazer?
[00:15] Eu quero enviar uma solicitação post, com os dados desse novo aluno no corpo da requisição, e inserir esse novo aluno no banco de dados. Vamos ver como fazemos isso? Para começar, vamos precisar de dois arquivos, alterar dois arquivos, o arquivo de rotas, porque vamos gerar uma nova rota, "Ctrl + P" aqui. E o arquivo de controller. O "Ctrl + P" é só para abrirmos essas abas que aparecem, os arquivos específicos.

[00:41] Então terá um cara responsável apenas em criar esse novo aluno, um novo controller responsável só por criar. Em "controller.go" eu vou criar uma nova função, vou chamar de func CriaNovoAluno(). Esse CriaNovoAluno, ele vai manter a convenção do que nós já temos, (c *gin.Context).

[01:03] Nas nossas rotas eu vou criar um novo endpoint responsável por criar um novo aluno. r. - repare que nós usamos o .GET para recuperarmos um determinado recurso ou realizar uma determinada ação.

[01:17] O que queremos fazer agora, com é criar, é uma requisição r.POST(), assim como já vimos no nosso curso anterior, utilizando o Gorilla Mux. Então é uma requisição post em r.POST("/alunos", ), nós já sabemos que quem vai atender essa requisição é esse , controllers.CriaNovoAluno).

[01:37] Então já temos o endpoint, falamos que a requisição é uma ação post, o que vamos fazer agora é implementar essa função para que ela de fato crie um novo aluno. Em "controller.go", a primeira coisa de tudo, precisaremos de um endereço de memória de um aluno, para ele conseguir vincular o corpo da requisição com, de fato, a struct que nós já temos.

[01:58] Para isso eu vou criar uma nova variável chamada aluno, ela será do tipo var alunos models.Aluno. Existe uma função, aqui no Gin, que conseguimos empacotar todo o corpo da solicitação, da entrada, com base na struct que nós já temos. Isso vai facilitar muito, não precisaremos falar: pegue, do corpo da requisição, o nome e vincule com o nome desse aluno, o CPF vinculando com a struct CPF.

[02:26] Ele já faz isso para nós, então eu vou mostrar isso para vocês. Primeira coisa, eu vou verificar se tem alguma mensagem de erro nessa ação que ele vai fazer, if err := c.Should, que é essa função que eu estou falando para vocês, c.ShouldBind.

[02:47] Repare que ele tem várias coisas, várias ações que ele pode fazer. A requisição que vai vir para nós é uma requisição JSON, então eu vou deixar c.ShouldBindJSON(). Ou seja, pegue todo o corpo da solicitação e empacote com base nessa struct que eu vou te passar agora. Qual struct vamos passar agora para ele? A struct de (&aluno); que nós temos ali em cima.

[03:07] Ele já vai fazer esse vínculo de todos os alunos, que nós precisamos. Vamos verificar só se o erro não é nulo? Ponto e vírgula, se o err != nil{}, não for igual a nil, se tivermos um erro, o que vai acontecer? Nós podemos exibir uma mensagem como nós já conhecemos, por exemplo, c.JSON(), eu vou passar aqui, vou indicar o status code e vou falar: olha, é uma requisição que não deu certo.

[03:32] Vou pegar (http.StatusBadRequest, ), nós não conseguimos fazer o que nós queríamos, e vou colocar aqui o , gin.H{}), que já conhecemos também. Dentro das chaves eu vou falar: nós temos um erro, vou colocar "error": e vou imprimir a mensagem de erro que vamos receber, : err.Error().

[03:53] Vou trazer essas chaves para cima. Nós já indicamos que temos um erro na nossa mensagem, um erro no nosso retorno da requisição. Vou colocar um return vazio.

[04:04] Caso não tenhamos erro - esse foi um status que demos erro, nós pegamos, empacotamos os dados e alguma coisa não deu certo, ele vai nos informar qual é o erro. Caso não tenhamos nenhum erro significa que conseguimos empacotar todo o corpo da requisição, os nomes, o nome do aluno, o RG e CPF, deu tudo certo, e queremos salvar isso no banco de dados.

[04:26] Para isso usamos o nosso database.DB.Create(). Eu vou falar: crie um novo aluno. Endereço de memória (&aluno). Olha como fica muito simples.

[04:39] Para finalizar e deixar a nossa API ainda melhor, o que eu posso fazer é informar uma mensagem falando: deu tudo certo, você criou esse aluno no banco de dados. Eu posso fazer aqui no c.JSON(), que já utilizamos na nossa mensagem de erro, só que aqui é um caso de sucesso, então (http.StatusOK, ), tudo junto. Vírgula e aqui eu vou passar o aluno que nós geramos, (http.StatusOK, aluno).

[05:08] Pronto, é isso, para criarmos de fato um novo registro no nosso banco de dados, vamos precisar dessa função que pega, esse ShouldBindJSON, que é super importante, ele vai pegar todo o corpo da solicitação de entrada com base na struct de alunos que nós temos e depois colocamos uma rota post.

[05:27] Vou desligar, derrubar o meu servidor, ligar ele, rodar ele mais uma vez, go run main.go, vou puxar esse tela para o lado, e vou testar isso no Postman. Você pode testar em outros locais também se você quiser, se tiver base de conhecimento em outra forma de testar a API, eu vou testar no Postman.

[05:49] Vamos para o Postman. No Postman eu vou colocar que a nossa requisição. Será uma requisição "POST" para "localhost:8080/alunos/", GG, está tudo certo. Agora o que precisamos fazer é criar o corpo da requisição. Aqui eu tenho o corpo de uma requisição.

[06:09] Abre chaves, nome, RG e CPF, vou colocar um RG com a quantidade de dígitos mesmo, então "rg":"1234563258", um RG fictício, o "cpf":"40255111230", coloquei qualquer coisa aqui, espero que esse CPF não seja de ninguém. Maravilha, tenho a nossa requisição, tenho essas informações.

[06:32] O que eu vou fazer agora é enviar essa solicitação. O que vai acontecer? Quando eu clicar no "Send", essa nova aluna, a Ana, ela será criada no nosso banco de dados. Então estamos vinculando o nosso servidor já com o banco de dados, com a requisição post e essa maluquice toda. Vamos lá, enviei no "Send".

[06:50] ID 1, olha que interessante, created at, updated e deleted at, ele falou, não foi deletado, foi criado nesse horário, nesta data, e tem aqui o nome Ana, CPF e o RG.

[07:05] Até aqui tudo bem, é o comportamento que esperamos, vamos para a prova de fogo. O que eu vou fazer? No pgAdmin, em "alunos", vou colocar para visualizar todas as linhas mais uma vez, ele está pensando. Quando ele visualiza, está lá. ID 1, o deleted at null e o nome Ana, aquele CPF maluco que eu digitei e o RG também muito doido que eu digitei.

[07:24] Qual é o meu desafio para você agora? Meu desafio é: crie novos alunos para essa nossa base de dados, depois nós vamos fazer coisas muito legais com isso. Eu espero que você tenha gostado, é uma forma de conectar com o banco de dados, criar uma requisição com o verbo certo para conseguirmos gerar novos alunos agora no Postgres.

@@06
Método db.AutoMigrate

Durante esta aula, vimos como criar uma tabela no banco de dados com base na struct de aluno que temos no código, e como adicionar um aluno a partir de um endpoint de método POST.
Sabendo disso, marque apenas a opção correta que representa a funcionalidade do método ‘db.AutoMigrate’ que temos no código:

O método db.AutoMigrate possui a funcionalidade automática de gerar e persistir dados aleatórios no banco de dados.Estes dados aleatórios são gerados a partir da referência de memória de uma struct que é passada como parâmetro para o método.
 
Alternativa correta
O método db.AutoMigrate possui a funcionalidade de criar automaticamente novos endpoints para a aplicação.
 
Alternativa correta
O método db.AutoMigrate possui a funcionalidade de migrar modelos(structs) que estão em código Go para o banco de dados. Os modelos são criados no banco de dados utilizando como base os dados acessados a partir da referência de memória de uma ou mais structs que são passadas como parâmetro para o método db.AutoMigrate.
 
Certo! O método db.AutoMigrate realmente possui a funcionalidade de migrar modelos que estão em código para o banco de dados.

@@07
Faça como eu fiz

Chegou a hora de você seguir todos os passos realizados por mim durante esta aula. Caso já tenha feito isso, excelente. Se ainda não fez, é importante que você implemente o que foi visto no vídeo para poder continuar com a próxima aula, que tem como pré-requisito todo o código escrito até o momento.

Caso não encontre uma solução nas perguntas feitas por alunos e alunas deste curso, para comunicar erros e tirar dúvidas de forma eficaz, clique neste link e saiba como utilizar o fórum da Alura.
Não tem dúvidas? Que tal ajudar alguém no fórum?

https://cursos.alura.com.br/comunicando-erros-e-tirando-duvidas-em-foruns-c19

@@08
O que aprendemos?

Nesta aula:
Instalamos o Gorm para não escrever código sql, facilitando a comunicação da aplicação com o banco de dados;
Conectamos a API com banco de dados e realizamos uma migração com base na struct de aluno;
Alteramos o controle para exibir os alunos registrados no banco de dados!
Na próxima aula
Vamos aprender como criar, editar e deletar o cadastro de um aluno ou aluna!

