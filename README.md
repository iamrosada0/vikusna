## EvaEats

`:rocket: → :rocket: (🚀)
:sparkles: → :sparkles: (✨)
:bug: → :bug: (🐛)
:memo: → :memo: (📝)
:art: → :art: (🎨)
:tada: → :tada: (🎉)
:construction: → :construction: (🚧)
:fire: → :fire: (🔥)
:books: → :books: (📚)
`

o fluxo geral da aplicação para chefs de pequenos negócios de comida:

1. **Cadastro de Chefs**:
   - Os chefs se cadastram na plataforma, fornecendo informações como nome, endereço, informações de contato e detalhes sobre seu negócio de comida.

2. **Gerenciamento de Pratos**:
   - Após o cadastro, os chefs podem adicionar os pratos que desejam oferecer aos clientes. Eles podem incluir detalhes como nome do prato, descrição, preço e disponibilidade.

3. **Visualização de Cardápio**:
   - Os pratos adicionados pelos chefs são exibidos em um cardápio virtual na plataforma. Os clientes podem navegar por esse cardápio e visualizar os pratos disponíveis.

4. **Recebimento de Pedidos**:
   - Quando os clientes fazem pedidos através da plataforma, os chefs recebem notificações em tempo real sobre os novos pedidos. Eles podem visualizar os detalhes do pedido, como os pratos solicitados, o endereço de entrega e as informações do cliente.

5. **Preparação e Confirmação do Pedido**:
   - Os chefs preparam os pratos conforme os pedidos recebidos e confirmam a preparação dos mesmos na plataforma. Isso permite que os clientes saibam que seus pedidos estão sendo preparados.

6. **Entrega dos Pedidos**:
   - Após a preparação, os chefs entregam os pedidos aos clientes. Eles podem marcar os pedidos como entregues na plataforma para manter os clientes atualizados sobre o status da entrega.

7. **Avaliações e Feedback**:
   - Os clientes têm a oportunidade de avaliar os pratos e a experiência geral de compra. Os chefs podem ver essas avaliações e usar o feedback para melhorar seus serviços.

8. **Gerenciamento da Conta**:
   - Os chefs podem acessar sua conta na plataforma para gerenciar suas informações, adicionar novos pratos, visualizar relatórios de vendas e interagir com os clientes.


 checklist para o fluxo completo da aplicação para chefs de pequenos negócios de comida:

Entendido. Vamos atualizar o checklist para refletir essa distinção:

1. **Cadastro e Configuração da Conta:**
   - [ ] Página de cadastro para novos chefs no site.
   - [ ] Formulário de cadastro para chefs com campos para nome, endereço, informações de contato e detalhes sobre o negócio de comida.
   - [ ] Opções para os chefs configurarem preferências da conta, como horários de funcionamento, áreas de entrega e métodos de pagamento aceitos.
   - [ ] Interface de cadastro para usuários no aplicativo móvel.

2. **Gerenciamento de Cardápio:**
   - [ ] Painel administrativo no site para os chefs gerenciarem seu cardápio.
   - [ ] Funcionalidade para os chefs adicionarem, editarem e removerem pratos do cardápio pelo site.
   - [ ] Interface no aplicativo móvel para os usuários visualizarem o cardápio dos chefs.

3. **Recebimento de Pedidos Online:**
   - [ ] Funcionalidade no site para os clientes visualizarem o cardápio e fazerem pedidos.
   - [ ] Capacidade para os chefs receberem pedidos diretamente no painel administrativo do site.
   - [ ] Interface no aplicativo móvel para os usuários fazerem pedidos.

4. **Preparação e Confirmação do Pedido:**
   - [ ] Painel administrativo no site para os chefs visualizarem e confirmarem pedidos.
   - [ ] Opções para os chefs confirmarem a preparação dos pedidos no painel administrativo do site.
   - [ ] Notificações em tempo real para os chefs sobre novos pedidos recebidos no site.

5. **Entrega dos Pedidos:**
   - [ ] Ferramentas para os chefs planejarem e executarem as entregas dos pedidos, acessíveis pelo site.
   - [ ] Funcionalidade para os chefs marcarem os pedidos como entregues no painel administrativo do site.

6. **Avaliações e Feedback:**
   - [ ] Opção para os clientes avaliarem os pratos e a experiência de compra no aplicativo móvel.
   - [ ] Funcionalidade para os chefs visualizarem e responderem às avaliações dos clientes no site.

7. **Gerenciamento da Conta:**
   - [ ] Acesso ao painel administrativo da conta para os chefs gerenciarem suas informações, cardápio e pedidos, disponível no site.
   - [ ] Interface no aplicativo móvel para os usuários acessarem seu perfil, histórico de pedidos e preferências.

Esse checklist aborda as necessidades específicas dos chefs e dos usuários, garantindo que ambas as partes tenham uma experiência fluida e eficiente ao utilizar o EvaEats.

## BACKEND WITH GOLANG

checklist para o desenvolvimento do backend da aplicação EvaEats:

1. **Configuração e Infraestrutura:**
   - [ ] Configuração do ambiente de desenvolvimento com as ferramentas necessárias (por exemplo, Go, GORM, PostgreSQL).
   - [ ] Configuração do banco de dados PostgreSQL.
   - [ ] Criação do projeto Go e estrutura de pastas.
   - [ ] Configuração do roteamento HTTP com um framework como Gin ou Chi.

2. **Modelagem de Dados:**
   - [ ] Definição dos modelos de dados necessários, incluindo User, Chef, Dish, Order, etc.
   - [ ] Mapeamento dos modelos para as tabelas do banco de dados utilizando GORM.
   - [ ] Implementação de validações nos modelos, como restrições de campos obrigatórios e tipos de dados.

3. **Endpoints da API:**
   - [ ] Implementação dos endpoints da API para as operações CRUD em cada modelo (por exemplo, criar, ler, atualizar, excluir).
   - [ ] Endpoints para autenticação e autorização de usuários e chefs.
   - [ ] Implementação de rotas protegidas que exigem autenticação para acesso.

4. **Autenticação e Autorização:**
   - [ ] Implementação de autenticação de usuários utilizando JWT (JSON Web Tokens).
   - [ ] Configuração de middleware para verificar a validade dos tokens JWT em rotas protegidas.
   - [ ] Implementação de lógica de autorização para garantir que os usuários só possam acessar recursos que lhes pertençam.

5. **Integração com o Banco de Dados:**
   - [ ] Configuração da conexão com o banco de dados PostgreSQL.
   - [ ] Implementação de funções para executar operações CRUD nos modelos utilizando GORM.
   - [ ] Criação de seeds para popular o banco de dados com dados iniciais (opcional, dependendo das necessidades de teste).

6. **Testes Unitários e de Integração:**
   - [ ] Implementação de testes unitários para as funções de manipulação de dados.
   - [ ] Implementação de testes de integração para os endpoints da API, incluindo casos de teste para diferentes cenários.

7. **Documentação da API:**
   - [ ] Documentação dos endpoints da API utilizando ferramentas como Swagger ou Postman.
   - [ ] Inclusão de exemplos de solicitação e resposta para cada endpoint, incluindo parâmetros de entrada e saída esperados.

8. **Configuração do Ambiente de Produção:**
   - [ ] Configuração de variáveis de ambiente para o ambiente de produção.
   - [ ] Configuração de servidores web como Nginx para lidar com o roteamento e a segurança.
   - [ ] Configuração de HTTPS e certificados SSL para comunicação segura.

9. **Monitoramento e Logging:**
   - [ ] Implementação de logging para registrar informações importantes sobre as solicitações recebidas e as operações realizadas.
   - [ ] Configuração de ferramentas de monitoramento para acompanhar o desempenho da aplicação e identificar possíveis problemas.

10. **Deploy e Implantação:**
    - [ ] Preparação da aplicação para implantação em um ambiente de produção.
    - [ ] Deploy da aplicação em um servidor de produção utilizando uma plataforma de hospedagem como Heroku ou AWS.


## BACKEND LOGICS 



1. **Endpoints Adicionais da API:**
   - [ ] Implementação de endpoints para que os usuários possam visualizar o cardápio dos chefs.
   - [ ] Implementação de endpoints para que os usuários possam fazer pedidos de pratos específicos.
   - [ ] Implementação de endpoints para que os chefs possam receber e gerenciar pedidos feitos pelos usuários.

2. **Gerenciamento de Pedidos:**
   - [ ] Implementação de lógica para adicionar pedidos à lista de pedidos de um chef quando um usuário faz um pedido.
   - [ ] Implementação de funcionalidades para que os chefs possam visualizar os pedidos recebidos e seu status.
   - [ ] Implementação de funcionalidades para que os chefs possam atualizar o status dos pedidos (por exemplo, confirmado, em preparo, entregue).

3. **Notificações em Tempo Real:**
   - [ ] Configuração de notificações em tempo real para os chefs quando um novo pedido é feito por um usuário.
   - [ ] Implementação de notificações para os usuários sobre o status do pedido (por exemplo, pedido confirmado, em preparo, entregue).

4. **Pagamento e Checkout:**
   - [ ] Implementação de funcionalidades para que os usuários possam pagar pelos pedidos através do aplicativo.
   - [ ] Integração com serviços de pagamento para processar transações de pagamento de forma segura.
   - [ ] Implementação de funcionalidades para que os chefs possam confirmar o pagamento e iniciar a preparação do pedido após a confirmação do pagamento.

5. **Histórico de Pedidos:**
   - [ ] Implementação de funcionalidades para que os usuários possam visualizar seu histórico de pedidos anteriores.
   - [ ] Implementação de funcionalidades para que os chefs possam acessar o histórico de pedidos recebidos e completados.

6. **Avaliações e Feedback:**
   - [ ] Implementação de funcionalidades para que os usuários possam avaliar os pratos e deixar feedback sobre a experiência.
   - [ ] Implementação de funcionalidades para que os chefs possam visualizar e responder às avaliações dos usuários.

