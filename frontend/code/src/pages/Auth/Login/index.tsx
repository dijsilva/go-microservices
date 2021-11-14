import React, { useState } from "react";
import axios from 'axios';

import enrollService from '../../../services/enroll';

import { FiLogIn } from 'react-icons/fi';

import { useNavigate } from 'react-router-dom';

import { TextInput, Form, LabelInput, RegisterButton } from "./styles";

import { toast } from 'react-toastify';

interface RegisterFormProps {
  backToHomeScreen: () => void;
}

export const LoginForm = ({ backToHomeScreen }: RegisterFormProps): JSX.Element => {
  const navigate = useNavigate()
  const [valuesToRegister, setValuesToRegister] = useState({
    email: "",
    password: "",
  });

  const updateFormValues = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { value, name } = event.target;
    setValuesToRegister({ ...valuesToRegister, [name]: value });
  };

  const onSubmitForm = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    let response;
    try {
      response = await enrollService.post('users/auth', {
        ...valuesToRegister,
      })
      if (response.data && response.data.token) {
        const { token } = response.data;
        const userData = {
          userName: response.data.username,
          name: response.data.name,
          email: response.data.email,
          profile: response.data.profile,
        }
        localStorage.setItem('user', JSON.stringify(userData))
        localStorage.setItem('token', JSON.stringify(token))
        navigate('home')
      } else {
        toast.error('Ocorreu um erro no servidor')
      }
    } catch (err) {
      if (axios.isAxiosError(err)) {
        if (err.response?.status === 401) {
          toast.error("Credenciais incorretas")
        } else if (err.response?.status === 404) {
          toast.error("Usuário não encontrado")
        }
      } else {
        toast.error('Ocorreu um erro no servidor')
      }
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
        <LabelInput>Senha</LabelInput>
        <TextInput
          type="password"
          name="password"
          placeholder="Defina uma senha"
          onChange={updateFormValues}
          value={valuesToRegister.password}
        />
      </div>
      <RegisterButton type="submit"><FiLogIn size={28} /></RegisterButton>
      <RegisterButton onClick={() => backToHomeScreen()}>Voltar</RegisterButton>
    </Form>
  );
};
