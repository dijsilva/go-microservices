import React from "react";

import { useNavigate } from 'react-router-dom';

import { ContainerOptions, ContainerOption } from "./styles";
import { FiSend, FiList } from "react-icons/fi";


export const Options = () => {
  const navigate = useNavigate();

  const goTo = (page: string): void => {
    navigate(page)
  }

  return (
    <ContainerOptions>
      <ContainerOption onClick={() => goTo('/new')}>
        <FiSend size={80} />
        <p>Enviar novo espectro</p>
      </ContainerOption>
      <ContainerOption onClick={() => goTo('/list')}>
        <FiList size={80} />
        <p>Ver espectros enviados</p>
      </ContainerOption>
    </ContainerOptions>
  );
};
