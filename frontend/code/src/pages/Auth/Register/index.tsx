import React, { useState } from "react";
import axios from 'axios';

import enrollService from '../../../services/enroll';

import { TextInput, Form, LabelInput, RegisterButton } from "./styles";

import { toast } from 'react-toastify';

interface RegisterFormProps {
  backToHomeScreen: () => void;
}

export const RegisterForm = ({ backToHomeScreen }: RegisterFormProps): JSX.Element => {
  const [valuesToRegister, setValuesToRegister] = useState({
    name: "",
    username: "",
    email: "",
    password: "",
  });

  const updateFormValues = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { value, name } = event.target;
    setValuesToRegister({ ...valuesToRegister, [name]: value });
  };

  const passwordIsValid = (password: string): boolean => {
    return /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/.test(password)
  }

  const onSubmitForm = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!passwordIsValid(valuesToRegister.password)) {
      toast.error('Sua senha deve ter conter pelo menos 8 caracteres entre números e letras')
      return
    }
    let response;
    try {
      response = await enrollService.post('users/new', {
        ...valuesToRegister,
      })
    } catch (err) {
      if (axios.isAxiosError(err)) {
        if (err.response?.data?.data?.message === 'User already exists') {
          toast.error("Usuário já cadastrado")
          setTimeout(() => {
            backToHomeScreen()
          }, 3000)
        } else {
          toast.error("Erro no servidor")
          setTimeout(() => {
            backToHomeScreen()
          }, 3000)
        }
      } else {
        toast.error("Erro inesperado")
        setTimeout(() => {
          backToHomeScreen()
        }, 3000)
      }
    }
    if (response && response.status === 201) {
      toast.success("Usuário criado com sucesso!")
      setTimeout(() => {
        backToHomeScreen()
      }, 3000)
    } else if (response?.status === 400 && response.data.message === 'User alreay exists') {
      toast.error("Usuário já cadastro com esses dados")
      setTimeout(() => {
        backToHomeScreen()
      }, 3000)
    }
  };


  return (
    <Form onSubmit={onSubmitForm}>
      <div>
        <LabelInput>E-mail</LabelInput>
        <TextInput
          type="email"
          name="email"
          placeholder="Digite seu e-mail"
          onChange={updateFormValues}
          value={valuesToRegister.email}
        />
      </div>
      <div>
        <LabelInput>Username</LabelInput>
        <TextInput
          type="text"
          name="username"
          placeholder="Escolha um username"
          onChange={updateFormValues}
          value={valuesToRegister.username}
        />
      </div>
      <div>
        <LabelInput>Nome completo</LabelInput>
        <TextInput
          type="text"
          name="name"
          placeholder="Digite seu nome completo"
          onChange={updateFormValues}
          value={valuesToRegister.name}
        />
      </div>
      <div>
        <LabelInput>Senha</LabelInput>
        <TextInput
          type="password"
          name="password"
          placeholder="Defina uma senha"
          onChange={updateFormValues}
          value={valuesToRegister.password}
        />
      </div>
      <RegisterButton type="submit">Criar conta!</RegisterButton>
      <RegisterButton onClick={() => backToHomeScreen()}>Voltar</RegisterButton>
    </Form>
  );
};
