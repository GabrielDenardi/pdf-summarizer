# Resumo de PDF com IA

## Visão Geral

**Resumo de PDF com IA** é uma aplicação inovadora que permite ao usuário enviar um PDF e receber um resumo gerado por inteligência artificial. A ferramenta utiliza **Go** para o backend e a poderosa **API do ChatGPT** para criar resumos claros e concisos. A interface é moderna, responsiva e conta com animações, efeito glassmorphism e dark mode, tudo desenvolvido com HTML, CSS e JavaScript puro.

## Funcionalidades

- **Upload e Extração de PDFs:**  
  O usuário realiza o upload do arquivo PDF e o sistema extrai o conteúdo utilizando bibliotecas Go.

- **Integração com ChatGPT:**  
  O texto extraído é enviado para a API do ChatGPT, que gera um resumo detalhado e objetivo.

- **Interface Moderna e Interativa:**  
  - **Design Clean e Minimalista:** Com animações suaves (fade-in, transições) e efeito glassmorphism.
  - **Dark Mode:** Alternância fácil com um botão fixo no canto, que muda o ícone conforme o tema.
  - **Spinner de Carregamento:** Feedback visual durante o processamento do PDF.

- **Boas Práticas de Desenvolvimento:**  
  Uso de variáveis de ambiente via arquivo `.env` para proteger chaves sensíveis e uma estrutura modular que facilita a manutenção e evolução do projeto.

## Tecnologias Utilizadas

- **Backend:** Go (Golang)
- **IA:** API do ChatGPT
- **Frontend:** HTML, CSS e JavaScript (sem frameworks pesados)
- **Gerenciamento de Dependências:** Go Modules
- **Variáveis de Ambiente:** [godotenv](https://github.com/joho/godotenv)

## Pré-requisitos

- [Go](https://golang.org/doc/install) (versão 1.16 ou superior)
- Conta e chave de API na [OpenAI](https://platform.openai.com)
- Git

## Instalação e Execução

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/GabrielDenardi/pdf-summarizer.git
   cd pdf-summarizer
